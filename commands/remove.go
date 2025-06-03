package commands

import (
	"os"

	"tasker/core"
	"tasker/utils"
)

// Remove exclui o favorito baseado no nome. Se não existir, loga erro e sai.
func Remove(name string) {
	favoritos, err := core.LoadFavoritos()
	if err != nil {
		utils.LogError(err.Error())
		os.Exit(1)
	}

	if _, exists := favoritos[name]; !exists {
		utils.LogError(`Não existe favorito com o nome "` + name + `"`)
		os.Exit(1)
	}

	delete(favoritos, name)
	if err := core.SaveFavoritos(favoritos); err != nil {
		utils.LogError(err.Error())
		os.Exit(1)
	}

	utils.LogSuccess(`Favorito "` + name + `" removido.`)
}
