package core

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/go-runewidth"
)

func ShowHelp() {
	fmt.Println(color.New(color.FgBlue, color.Bold).Sprint("Tasker v0.6.0 🚀"))
	fmt.Println(color.New(color.FgWhite).Sprint("Gerencie seus comandos favoritos."))

	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	fmt.Println()

	fmt.Println(bold("Uso:"))

	cmdArgsWidth := 32

	lines := []struct {
		cmd  string
		args string
		desc string
	}{
		{green("tasker add"), yellow("<nome> '<comando>'"),  fmt.Sprintf("→ Adiciona um favorito, use %s para marcar como argumento.", cyan("$1"))},
		{green("tasker list"), "", "→ Lista os favoritos."},
		{green("tasker remove"), yellow("<nome>"), "→ Remove um favorito."},
		{green("tasker run"), yellow("<nome> [arg]"), fmt.Sprintf("→ Executa o favorito, substituindo %s pelo argumento.", cyan("$1"))},
		{green("tasker setup"), "", "→ Aplica o CLI na sua PATH."},
	}

	for _, line := range lines {
		cmdAndArgs := line.cmd
		if line.args != "" {
			cmdAndArgs += "  " + line.args
		}

		fmt.Printf("  %s %s\n",
			padRightVisual(cmdAndArgs, cmdArgsWidth),
			line.desc,
		)
	}

	fmt.Println(bold("Exemplos de uso:"))
	fmt.Printf("  tasker add startDB %s\n", cyan("'docker compose start database'"))
	fmt.Printf("  tasker run startDB\n\n")

}

var ansiEscape = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripAnsi(str string) string {
	return ansiEscape.ReplaceAllString(str, "")
}

func padRightVisual(s string, width int) string {
	// Remove sequências ANSI para contar só os caracteres visíveis
	plain := stripAnsi(s)
	visibleLen := runewidth.StringWidth(plain)
	if visibleLen >= width {
		return s
	}
	return s + strings.Repeat(" ", width-visibleLen)
}
