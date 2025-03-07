package cache

import (
	"cz/internal"
	"os"
	"path/filepath"
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
