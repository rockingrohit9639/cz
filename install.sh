#!/bin/sh
# Installer for cz - https://github.com/rockingrohit9639/cz
#
#   curl -sSfL https://raw.githubusercontent.com/rockingrohit9639/cz/main/install.sh | sh
#
# Environment variables:
#   CZ_VERSION    version to install (e.g. v2.1.0). Default: latest release.
#   CZ_INSTALL_DIR  where to install. Default: /usr/local/bin, falling back to
#                   ~/.local/bin when /usr/local/bin is not writable.

set -eu

REPO="rockingrohit9639/cz"
BINARY="cz"

info() { printf '\033[0;34m==>\033[0m %s\n' "$1"; }
warn() { printf '\033[0;33mwarn:\033[0m %s\n' "$1" >&2; }
fail() { printf '\033[0;31merror:\033[0m %s\n' "$1" >&2; exit 1; }

need() {
  command -v "$1" >/dev/null 2>&1 || fail "'$1' is required but was not found in PATH."
}

detect_platform() {
  os="$(uname -s)"
  case "$os" in
    Linux) os="linux" ;;
    Darwin) os="darwin" ;;
    *) fail "Unsupported operating system: $os. On Windows, use install.ps1 instead." ;;
  esac

  arch="$(uname -m)"
  case "$arch" in
    x86_64 | amd64) arch="amd64" ;;
    arm64 | aarch64) arch="arm64" ;;
    *) fail "Unsupported architecture: $arch. Prebuilt binaries exist for amd64 and arm64 only." ;;
  esac
}

# Downloads $1 to stdout.
fetch() {
  if command -v curl >/dev/null 2>&1; then
    curl -sSfL "$1"
  else
    wget -qO- "$1"
  fi
}

# Downloads $1 to file $2.
fetch_to() {
  if command -v curl >/dev/null 2>&1; then
    curl -sSfL -o "$2" "$1"
  else
    wget -qO "$2" "$1"
  fi
}

resolve_version() {
  if [ -n "${CZ_VERSION:-}" ]; then
    version="$CZ_VERSION"
    return
  fi

  info "Resolving latest release..."
  version="$(
    fetch "https://api.github.com/repos/${REPO}/releases/latest" |
      sed -n 's/.*"tag_name"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/p' |
      head -n 1
  )"

  [ -n "$version" ] || fail "Could not determine the latest release of ${REPO}. Set CZ_VERSION to install a specific version."
}

resolve_install_dir() {
  if [ -n "${CZ_INSTALL_DIR:-}" ]; then
    install_dir="$CZ_INSTALL_DIR"
  elif [ -w /usr/local/bin ] 2>/dev/null; then
    install_dir="/usr/local/bin"
  else
    install_dir="$HOME/.local/bin"
  fi
  mkdir -p "$install_dir" || fail "Could not create install directory: $install_dir"
  [ -w "$install_dir" ] || fail "Install directory is not writable: $install_dir. Set CZ_INSTALL_DIR to somewhere you can write."
}

verify_checksum() {
  # The checksum file is best-effort: if we can't fetch it or have no sha256
  # tool, warn rather than block the install.
  checksums="${tmp}/checksums.txt"
  if ! fetch_to "https://github.com/${REPO}/releases/download/${version}/checksums.txt" "$checksums" 2>/dev/null; then
    warn "Could not download checksums.txt; skipping integrity check."
    return
  fi

  if command -v sha256sum >/dev/null 2>&1; then
    actual="$(sha256sum "$1" | awk '{print $1}')"
  elif command -v shasum >/dev/null 2>&1; then
    actual="$(shasum -a 256 "$1" | awk '{print $1}')"
  else
    warn "No sha256sum or shasum found; skipping integrity check."
    return
  fi

  expected="$(awk -v f="$archive_name" '$2 == f || $2 == "*" f {print $1}' "$checksums" | head -n 1)"
  if [ -z "$expected" ]; then
    warn "No checksum listed for ${archive_name}; skipping integrity check."
    return
  fi

  [ "$actual" = "$expected" ] || fail "Checksum mismatch for ${archive_name}. Expected ${expected}, got ${actual}."
  info "Checksum verified."
}

main() {
  need uname
  need tar
  if ! command -v curl >/dev/null 2>&1 && ! command -v wget >/dev/null 2>&1; then
    fail "Either 'curl' or 'wget' is required."
  fi

  detect_platform
  resolve_version
  resolve_install_dir

  # Release archives are named without the leading "v" of the tag.
  bare_version="${version#v}"
  archive_name="${BINARY}_${bare_version}_${os}_${arch}.tar.gz"
  url="https://github.com/${REPO}/releases/download/${version}/${archive_name}"

  tmp="$(mktemp -d)"
  trap 'rm -rf "$tmp"' EXIT INT TERM

  info "Downloading ${BINARY} ${version} (${os}/${arch})..."
  fetch_to "$url" "${tmp}/${archive_name}" || fail "Download failed: $url"

  verify_checksum "${tmp}/${archive_name}"

  tar -xzf "${tmp}/${archive_name}" -C "$tmp" || fail "Could not extract ${archive_name}."
  [ -f "${tmp}/${BINARY}" ] || fail "Archive did not contain a '${BINARY}' binary."

  chmod +x "${tmp}/${BINARY}"
  mv "${tmp}/${BINARY}" "${install_dir}/${BINARY}" ||
    fail "Could not install to ${install_dir}. Try: CZ_INSTALL_DIR=\$HOME/.local/bin sh install.sh"

  info "Installed ${BINARY} ${version} to ${install_dir}/${BINARY}"

  case ":${PATH}:" in
    *":${install_dir}:"*)
      info "Run '${BINARY} version' to get started."
      ;;
    *)
      warn "${install_dir} is not on your PATH. Add this to your shell profile:"
      printf '\n    export PATH="%s:$PATH"\n\n' "$install_dir"
      ;;
  esac
}

main "$@"
