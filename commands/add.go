package commands

import (
	"tasker/core"
	"tasker/utils"
)

// Add adiciona um favorito. Se já existir, loga erro.
func Add(name string, commandStr string) {
	favoritos, err := core.LoadFavoritos()
	if err != nil {
		utils.LogError(err.Error())
		return
	}

	if _, exists := favoritos[name]; exists {
		utils.LogError(`Já existe um favorito chamado "` + name + `"`)
		return
	}

	favoritos[name] = commandStr
	if err := core.SaveFavoritos(favoritos); err != nil {
		utils.LogError(err.Error())
		return
	}

	utils.LogSuccess(`Favorito "` + name + `" adicionado com sucesso.`)
}
