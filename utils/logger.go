package utils

import (
	"github.com/fatih/color"
)

func LogInfo(message string) {
	color.New(color.FgBlue).Printf("%s\n", message)
}

func LogError(message string) {
	color.New(color.FgRed).Printf("%s\n", message)
}

func LogSuccess(message string) {
	color.New(color.FgGreen).Printf("%s\n", message)
}

func LogWarning(message string) {
	color.New(color.FgYellow).Printf("%s\n", message)
}
