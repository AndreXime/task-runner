package utils

import (
	"regexp"
	"strings"

	"github.com/mattn/go-runewidth"
)

var ansiEscape = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripAnsi(str string) string {
	return ansiEscape.ReplaceAllString(str, "")
}

func PadRightVisual(s string, width int) string {
	// Remove sequências ANSI para contar só os caracteres visíveis
	plain := stripAnsi(s)
	visibleLen := runewidth.StringWidth(plain)
	if visibleLen >= width {
		return s
	}
	return s + strings.Repeat(" ", width-visibleLen)
}
