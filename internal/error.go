package internal

import (
	"fmt"
	"os"
)

// ANSI color codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func Success(message string) {
	fmt.Println(Green + "Success: " + message + Reset)
}

func Error(message string) {
	fmt.Println(Red + "Error: " + message + Reset)
}

func Warn(message string) {
	fmt.Println(Yellow + "Warning: " + message + Reset)
}

func Info(message string) {
	fmt.Println(Blue + "Info: " + message + Reset)
}

func AbortOnError(err error, message string) {
	if err != nil {
		Error(message)
		os.Exit(1)
	}
}
