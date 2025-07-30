package commands

import (
	"os"
	"os/exec"
	"strings"

	"task-runner/core"
	"task-runner/utils"
)

// Run executa o favorito: substitui $1 se existir placeholder ou avisa se houver argumento extra.
func Run(name string, argumento string) {
	favoritos, err := core.LoadFavoritos()
	if err != nil {
		utils.LogError(err.Error())
		return
	}

	favorito, exists := favoritos[name]
	if !exists {
		utils.LogError(`Não existe favorito com o nome "` + name + `"`)
		return
	}

	comandoFinal := favorito

	if strings.Contains(favorito, "$1") {
		if argumento == "" {
			utils.LogError(`O favorito "` + name + `" requer um argumento para substituir $1.`)
			return
		}
		comandoFinal = strings.ReplaceAll(favorito, "$1", argumento)
	} else if argumento != "" {
		utils.LogWarning(`Esse comando não possui argumento, e assim será ignorado.`)
	}

	utils.LogInfo("Executando: " + comandoFinal + "\n")

	cmd := exec.Command("sh", "-c", comandoFinal)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}
