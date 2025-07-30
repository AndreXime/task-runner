package core

import (
	"fmt"
	"task-runner/utils"
)

type helpCommands = []struct {
	cmd  string
	args string
	desc string
}

func ShowHelp() {
	cmdArgsWidth := 32
	fmt.Println(utils.BlueBold("\nTaskRunner v0.7.0 🚀"))
	fmt.Println(utils.White("Gerencie seus comandos favoritos.\n"))

	fmt.Println(utils.Bold("Comandos:"))

	lines := helpCommands{
		{utils.Green("trun add"), utils.Yellow("<nome> '<comando>'"), fmt.Sprintf("Adiciona um favorito, use %s para marcar como argumento.", utils.Cyan("$1"))},
		{utils.Green("trun add"), utils.Yellow("<nome> --stdin"), "Lê o comando da entrada padrão (útil com aspas e caracteres especiais)."},
		{utils.Green("trun list"), "", "Lista todos os favoritos salvos."},
		{utils.Green("trun remove"), utils.Yellow("<nome>"), "Remove um favorito pelo nome."},
		{utils.Green("trun run"), utils.Yellow("<nome> [arg]"), fmt.Sprintf("Executa o favorito, substituindo %s pelo argumento fornecido.", utils.Cyan("$1"))},
		{utils.Green("trun setup"), "", "Adiciona o CLI à variável PATH."},
	}

	for _, line := range lines {
		cmdAndArgs := line.cmd
		if line.args != "" {
			cmdAndArgs += "  " + line.args
		}

		fmt.Printf("  %s %s\n",
			utils.PadRightVisual(cmdAndArgs, cmdArgsWidth),
			line.desc,
		)
	}

	fmt.Println(utils.Bold("\nExemplos de uso:"))
	fmt.Printf("  trun add startDB %s\n", utils.Cyan("'docker compose start database'"))
	fmt.Printf("  trun run startDB\n\n")

}
