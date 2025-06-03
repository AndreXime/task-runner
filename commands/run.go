package commands

import (
	"os"
	"os/exec"
	"strings"

	"tasker/core"
	"tasker/utils"
)

// Run executa o favorito: substitui $1 se existir placeholder ou avisa se houver argumento extra.
func Run(name string, argumento string) {
	favoritos, err := core.LoadFavoritos()
	if err != nil {
		utils.LogError(err.Error())
		os.Exit(1)
	}

	favorito, exists := favoritos[name]
	if !exists {
		utils.LogError(`Não existe favorito com o nome "` + name + `"`)
		os.Exit(1)
	}

	comandoFinal := favorito

	if strings.Contains(favorito, "$1") {
		if argumento == "" {
			utils.LogError(`O favorito "` + name + `" requer um argumento para substituir $1.`)
			os.Exit(1)
		}
		comandoFinal = strings.ReplaceAll(favorito, "$1", argumento)
	} else if argumento != "" {
		utils.LogWarning(`O comando "` + name + `" não possui placeholder $1. O argumento "` + argumento + `" será ignorado.`)
	}

	utils.LogInfo("Executando: " + comandoFinal + "\n")

	cmd := exec.Command("sh", "-c", comandoFinal)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
}
