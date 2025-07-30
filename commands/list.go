package commands

import (
	"fmt"
	"sort"

	"task-runner/core"
	"task-runner/utils"
)

// List carrega todos os favoritos e exibe no console.
// Se não houver nenhum, mostra mensagem informativa.
func List() {
	favoritos, err := core.LoadFavoritos()
	if err != nil {
		utils.LogError(err.Error())
		return
	}

	if len(favoritos) == 0 {
		utils.LogInfo("Nenhum comando favorito cadastrado.")
		return
	}

	utils.LogInfo("\nComandos favoritos:")

	// Ordena nomes para saída previsível
	nomes := make([]string, 0, len(favoritos))
	for nome := range favoritos {
		nomes = append(nomes, nome)
	}
	sort.Strings(nomes)

	for _, nome := range nomes {
		fmt.Printf("  • %s → '%s'\n", nome, favoritos[nome])
	}

	fmt.Println()
}
