package cache

import (
	"cz/internal"
	"os"
)

var cache = Cache{}

// This function initializes the cache directory for the 'cz'.
// It creates the directory if it doesn't exist. If any errors occur,
// the function aborts and displays an appropriate error message.
func Init() {
	czCacheDir := getCacheDirPath()

	// Create cache directory if it does not exists yet
	err := os.MkdirAll(czCacheDir, 0755)
	internal.AbortOnError(err, "could not create cz cache directory")

	// Create cz.json cache file with empty {} json if it does not exists yet
	// Create profiles.json if it does not exists
	czCacheFile := getCacheFilePath()
	if _, err := os.Stat(czCacheFile); os.IsNotExist(err) {
		createEmptyCacheFile()
		return
	}

	loadCache()
}

func SetPrevCommit(message string) {
	cache.PrevCommit = message
	saveCache()
}

func GetPrevCommit() string {
	return cache.PrevCommit
}
