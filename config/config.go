package config

import (
	"cz/internal"
	"fmt"
	"os"
	"text/template"
)

var config = Config{}

// This function initializes the config directory for the 'cz'.
// It creates the directory if it doesn't exist. If any errors occur,
// the function aborts and displays an appropriate error message.
func Init() {
	czConfigDir := getConfigDirPath()
	fmt.Println(czConfigDir)

	// Create config directory if it does not exists yet
	err := os.MkdirAll(czConfigDir, 0755)
	internal.AbortOnError(err, "could not create cz config directory")

	// Create config.json file with empty {} json if it does not exists yet
	czConfigFile := getConfigFilePath()
	if _, err := os.Stat(czConfigFile); os.IsNotExist(err) {
		createEmptyConfigFile()
		return
	}

	// Load the config file
	loadConfig()
}

func AddFormat(name string, pattern string) {
	format := Format{
		Name:    name,
		Pattern: pattern,
	}

	_, err := template.New(name).Parse(pattern)
	internal.AbortOnError(err, "failed to parse the pattern")

	config.Formats = append(config.Formats, format)
	saveConfig()
}
