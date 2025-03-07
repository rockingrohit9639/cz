package cache

import (
	"cz/internal"
	"encoding/json"
	"os"
)

func createEmptyCacheFile() {
	czCacheFile := getCacheFilePath()

	emptyJSON := []byte("{}")

	err := os.WriteFile(czCacheFile, emptyJSON, 0644)
	internal.AbortOnError(err, "could not create cz cache file")
}

func loadCache() {
	czCacheFile := getCacheFilePath()

	// read the content of cz.json file
	data, err := os.ReadFile(czCacheFile)
	internal.AbortOnError(err, "could not read cz cache file")

	// Parse and load cz.json in &Cache
	err = json.Unmarshal(data, &cache)
	internal.AbortOnError(err, "could not load cz cache file")
}

func saveCache() {
	czCacheFile := getCacheFilePath()

	data, err := json.MarshalIndent(cache, "", "    ")
	internal.AbortOnError(err, "could not parse cache data")

	os.WriteFile(czCacheFile, data, 0644)
}
