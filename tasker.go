package main

import (
	"os"
	"strings"

	"tasker/commands"
	"tasker/core"
	"tasker/utils"
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
			os.Exit(1)
		}
		nome := rest[0]
		comandoStr := strings.Join(rest[1:], " ")
		commands.Add(nome, comandoStr)

	case "list":
		commands.List()

	case "remove":
		if len(rest) != 1 {
			utils.LogError("Uso incorreto do comando remove.")
			core.ShowHelp()
			os.Exit(1)
		}
		commands.Remove(rest[0])

	case "run":
		if len(rest) < 1 {
			utils.LogError("Uso incorreto do comando run.")
			core.ShowHelp()
			os.Exit(1)
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
