package config

import (
	"os"
	"path/filepath"

	"github.com/rockingrohit9639/cz/internal"
)

func getConfigDirPath() string {
	userConfigDir, err := os.UserConfigDir()
	internal.AbortOnError(err, "Could not find user's config directory.")

	czConfigDir := filepath.Join(userConfigDir, "cz")

	return czConfigDir
}

func getConfigFilePath() string {
	czConfigDir := getConfigDirPath()
	return filepath.Join(czConfigDir, "config.json")
}
