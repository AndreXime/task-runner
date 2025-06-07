package utils

import (
	"github.com/fatih/color"
)

func LogInfo(message string) {
	color.New(color.FgHiBlue).Printf("%s\n", message)
}

func LogError(message string) {
	color.New(color.FgHiRed).Printf("%s\n", message)
}

func LogSuccess(message string) {
	color.New(color.FgHiGreen).Printf("%s\n", message)
}

func LogWarning(message string) {
	color.New(color.FgHiYellow).Printf("%s\n", message)
}

var (
	BlueBold = color.New(color.FgBlue, color.Bold).SprintFunc()
	White    = color.New(color.FgWhite).SprintFunc()
	Bold     = color.New(color.Bold).SprintFunc()
	Cyan     = color.New(color.FgCyan).SprintFunc()
	Yellow   = color.New(color.FgYellow).SprintFunc()
	Green    = color.New(color.FgGreen).SprintFunc()
)
