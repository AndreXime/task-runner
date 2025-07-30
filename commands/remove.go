package commands

import (
	"task-runner/core"
	"task-runner/utils"
)

// Remove exclui o favorito baseado no nome. Se não existir, loga erro e sai.
func Remove(name string) {
	favoritos, err := core.LoadFavoritos()
	if err != nil {
		utils.LogError(err.Error())
		return
	}

	if _, exists := favoritos[name]; !exists {
		utils.LogError(`Não existe favorito com o nome "` + name + `"`)
		return
	}

	delete(favoritos, name)
	if err := core.SaveFavoritos(favoritos); err != nil {
		utils.LogError(err.Error())
		return
	}

	utils.LogSuccess(`Favorito "` + name + `" removido.`)
}
