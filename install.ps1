# Installer for cz - https://github.com/rockingrohit9639/cz
#
#   irm https://raw.githubusercontent.com/rockingrohit9639/cz/main/install.ps1 | iex
#
# Parameters can also be supplied as environment variables CZ_VERSION and
# CZ_INSTALL_DIR when piping the script into iex.

[CmdletBinding()]
param(
    [string]$Version = $env:CZ_VERSION,
    [string]$InstallDir = $env:CZ_INSTALL_DIR
)

$ErrorActionPreference = 'Stop'
$ProgressPreference = 'SilentlyContinue'  # makes Invoke-WebRequest much faster

$Repo = 'rockingrohit9639/cz'
$Binary = 'cz'

# Windows PowerShell 5.1 negotiates TLS 1.0/1.1 by default, which github.com
# rejects. Without this the download fails with "Could not create SSL/TLS
# secure channel". Harmless on PowerShell 7+, which already defaults to 1.2+.
try {
    [Net.ServicePointManager]::SecurityProtocol =
        [Net.ServicePointManager]::SecurityProtocol -bor [Net.SecurityProtocolType]::Tls12
} catch {
    # Not available on some platforms; the default is already TLS 1.2+ there.
}

function Write-Info { param($Message) Write-Host "==> $Message" -ForegroundColor Blue }
function Write-Warn { param($Message) Write-Host "warn: $Message" -ForegroundColor Yellow }

function Get-Architecture {
    # PROCESSOR_ARCHITECTURE reports the process arch, which is wrong under
    # WOW64; OSArchitecture reports the machine. Fall back to the env var when
    # CIM is unavailable (e.g. PowerShell 7 on a stripped-down image).
    $arch = $null
    try { $arch = (Get-CimInstance Win32_OperatingSystem -ErrorAction Stop).OSArchitecture } catch { }
    if (-not $arch) { $arch = $env:PROCESSOR_ARCHITEW6432; }
    if (-not $arch) { $arch = $env:PROCESSOR_ARCHITECTURE }

    if ($arch -match 'ARM|aarch64') { return 'arm64' }
    if ($arch -match '64') { return 'amd64' }
    if ([Environment]::Is64BitOperatingSystem) { return 'amd64' }
    throw "Unsupported architecture: $arch. Prebuilt binaries exist for amd64 and arm64 only."
}

function Get-LatestVersion {
    Write-Info 'Resolving latest release...'
    try {
        $release = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/latest" -UseBasicParsing
    } catch {
        throw "Could not determine the latest release of $Repo. Set CZ_VERSION to install a specific version. ($($_.Exception.Message))"
    }
    if (-not $release.tag_name) {
        throw "Could not determine the latest release of $Repo. Set CZ_VERSION to install a specific version."
    }
    return $release.tag_name
}

function Test-Checksum {
    param($ArchivePath, $ArchiveName, $Version, $TempDir)

    $checksumPath = Join-Path $TempDir 'checksums.txt'
    try {
        Invoke-WebRequest -Uri "https://github.com/$Repo/releases/download/$Version/checksums.txt" `
            -OutFile $checksumPath -UseBasicParsing
    } catch {
        Write-Warn 'Could not download checksums.txt; skipping integrity check.'
        return
    }

    $expected = $null
    foreach ($line in Get-Content $checksumPath) {
        $parts = $line -split '\s+', 2
        if ($parts.Count -eq 2 -and $parts[1].TrimStart('*').Trim() -eq $ArchiveName) {
            $expected = $parts[0].Trim()
            break
        }
    }
    if (-not $expected) {
        Write-Warn "No checksum listed for $ArchiveName; skipping integrity check."
        return
    }

    $actual = (Get-FileHash -Path $ArchivePath -Algorithm SHA256).Hash
    if ($actual -ne $expected.ToUpperInvariant() -and $actual.ToLowerInvariant() -ne $expected.ToLowerInvariant()) {
        throw "Checksum mismatch for $ArchiveName. Expected $expected, got $actual."
    }
    Write-Info 'Checksum verified.'
}

function Add-ToUserPath {
    param($Directory)

    $userPath = [Environment]::GetEnvironmentVariable('Path', 'User')
    $entries = @()
    if ($userPath) { $entries = $userPath -split ';' | Where-Object { $_ } }

    if ($entries -contains $Directory) {
        return $false
    }

    $newPath = if ($userPath) { "$userPath;$Directory" } else { $Directory }
    [Environment]::SetEnvironmentVariable('Path', $newPath, 'User')
    # Also update the current session so cz works without reopening the shell.
    $env:Path = "$env:Path;$Directory"
    return $true
}

# --- main ---------------------------------------------------------------

function Invoke-Install {
    $arch = Get-Architecture

    if (-not $Version) { $Version = Get-LatestVersion }

    if (-not $InstallDir) {
        $InstallDir = Join-Path $env:LOCALAPPDATA "$Binary\bin"
    }
    New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null

    # Release archives are named without the leading "v" of the tag.
    $bareVersion = $Version -replace '^v', ''
    $archiveName = "${Binary}_${bareVersion}_windows_${arch}.zip"
    $url = "https://github.com/$Repo/releases/download/$Version/$archiveName"

    $tempDir = Join-Path ([System.IO.Path]::GetTempPath()) "cz-install-$([guid]::NewGuid())"
    New-Item -ItemType Directory -Force -Path $tempDir | Out-Null

    try {
        $archivePath = Join-Path $tempDir $archiveName

        Write-Info "Downloading $Binary $Version (windows/$arch)..."
        try {
            Invoke-WebRequest -Uri $url -OutFile $archivePath -UseBasicParsing
        } catch {
            throw "Download failed: $url ($($_.Exception.Message))"
        }

        Test-Checksum -ArchivePath $archivePath -ArchiveName $archiveName -Version $Version -TempDir $tempDir

        Expand-Archive -Path $archivePath -DestinationPath $tempDir -Force

        $exeSource = Join-Path $tempDir "$Binary.exe"
        if (-not (Test-Path $exeSource)) {
            throw "Archive did not contain a $Binary.exe binary."
        }

        $exeTarget = Join-Path $InstallDir "$Binary.exe"
        # Move-Item fails if cz.exe is running; surface that clearly.
        try {
            Move-Item -Path $exeSource -Destination $exeTarget -Force
        } catch {
            throw "Could not install to $exeTarget. Close any running '$Binary' process and try again. ($($_.Exception.Message))"
        }

        Write-Info "Installed $Binary $Version to $exeTarget"

        if (Add-ToUserPath -Directory $InstallDir) {
            Write-Info "Added $InstallDir to your user PATH."
            Write-Warn 'Open a new terminal for the PATH change to apply everywhere.'
        }
        Write-Info "Run '$Binary version' to get started."
    } finally {
        Remove-Item -Recurse -Force $tempDir -ErrorAction SilentlyContinue
    }
}

try {
    Invoke-Install
} catch {
    # Print a single readable line rather than a PowerShell stack trace.
    Write-Host "error: $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}
