package cache

import (
	"os"
	"path/filepath"

	"github.com/rockingrohit9639/cz/internal"
)

func getCacheDirPath() string {
	userCacheDir, err := os.UserCacheDir()
	internal.AbortOnError(err, "Could not find user's cache directory.")

	czCacheDir := filepath.Join(userCacheDir, "cz")

	return czCacheDir
}

func getCacheFilePath() string {
	czCacheDir := getCacheDirPath()
	return filepath.Join(czCacheDir, "cz.json")
}
