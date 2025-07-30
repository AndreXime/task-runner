package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"task-runner/commands"
	"task-runner/core"
	"task-runner/utils"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		core.ShowHelp()
		return
	}

	subcommand := args[0]
	rest := args[1:]

	switch subcommand {
	case "add":
		if len(rest) < 2 {
			utils.LogError("Uso incorreto do comando add.")
			core.ShowHelp()
			return
		}
		nome := rest[0]
		var comandoStr string

		if len(rest) == 2 && rest[1] == "--stdin" {
			fmt.Println("Cole o comando completo abaixo e pressione Ctrl+D para finalizar:")
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				utils.LogError("Erro ao ler do stdin: " + err.Error())
				return
			}
			comandoStr = strings.TrimSpace(string(data))
			if comandoStr == "" {
				utils.LogError("Não foi passado nenhum comando")
				return
			}
		} else {
			comandoStr = strings.Join(rest[1:], " ")
		}
		commands.Add(nome, comandoStr)

	case "list":
		commands.List()

	case "remove":
		if len(rest) != 1 {
			utils.LogError("Uso incorreto do comando remove.")
			core.ShowHelp()
			return
		}
		commands.Remove(rest[0])

	case "run":
		if len(rest) < 1 {
			utils.LogError("Uso incorreto do comando run.")
			core.ShowHelp()
			return
		}
		nome := rest[0]
		arg := ""
		if len(rest) > 1 {
			arg = rest[1]
		}
		commands.Run(nome, arg)

	case "setup":
		commands.SetupPath()

	case "help":
		fallthrough

	default:
		core.ShowHelp()
	}
}
