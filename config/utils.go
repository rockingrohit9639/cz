package config

import (
	"cz/internal"
	"encoding/json"
	"os"
)

func createEmptyConfigFile() {
	czConfigFile := getConfigFilePath()

	emptyJSON := []byte("{}")

	err := os.WriteFile(czConfigFile, emptyJSON, 0644)
	internal.AbortOnError(err, "could not create cz config file")
}

func loadConfig() {
	czConfigFile := getConfigFilePath()

	data, err := os.ReadFile(czConfigFile)
	internal.AbortOnError(err, "could not read cz config file")

	err = json.Unmarshal(data, &config)
	internal.AbortOnError(err, "could not load cz config file")
}

func saveConfig() {
	czConfigFile := getConfigFilePath()

	data, err := json.MarshalIndent(config, "", "    ")
	internal.AbortOnError(err, "could not parse config data")

	os.WriteFile(czConfigFile, data, 0644)
}
