package commands

import (
	"os"

	"tasker/core"
	"tasker/utils"
)

// Add adiciona um favorito. Se já existir, loga erro e faz Exit(1).
func Add(name string, commandStr string) {
	favoritos, err := core.LoadFavoritos()
	if err != nil {
		utils.LogError(err.Error())
		os.Exit(1)
	}

	if _, exists := favoritos[name]; exists {
		utils.LogError(`Já existe um favorito chamado "` + name + `"`)
		os.Exit(1)
	}

	favoritos[name] = commandStr
	if err := core.SaveFavoritos(favoritos); err != nil {
		utils.LogError(err.Error())
		os.Exit(1)
	}

	utils.LogSuccess(`Favorito "` + name + `" adicionado com sucesso.`)
}
