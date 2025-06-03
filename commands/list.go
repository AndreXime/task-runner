package commands

import (
	"fmt"
	"os"
	"sort"

	"tasker/core"
	"tasker/utils"
)

// List carrega todos os favoritos e exibe no console.
// Se não houver nenhum, mostra mensagem informativa.
func List() {
	favoritos, err := core.LoadFavoritos()
	if err != nil {
		utils.LogError(err.Error())
		os.Exit(1)
	}

	if len(favoritos) == 0 {
		utils.LogInfo("Nenhum comando favorito cadastrado.")
		return
	}

	fmt.Println("\n📜 Comandos favoritos:")

	// Ordena nomes para saída previsível
	nomes := make([]string, 0, len(favoritos))
	for nome := range favoritos {
		nomes = append(nomes, nome)
	}
	sort.Strings(nomes)

	for _, nome := range nomes {
		fmt.Printf("  • %s → %s\n", nome, favoritos[nome])
	}

	fmt.Println()
}
