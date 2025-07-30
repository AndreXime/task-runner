package core

import (
	"encoding/json"
	"errors"
	"os"
	"os/user"
	"path/filepath"
)

const fileName = ".fav-taskrunner.json"

type FavoritosMap map[string]string

func getFilePath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, fileName), nil
}

func LoadFavoritos() (FavoritosMap, error) {
	filePath, err := getFilePath()
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		// Se arquivo não existir, retorna mapa vazio
		if errors.Is(err, os.ErrNotExist) {
			return FavoritosMap{}, nil
		}
		return nil, err
	}

	var favs FavoritosMap
	if err := json.Unmarshal(content, &favs); err != nil {
		return nil, err
	}

	return favs, nil
}

func SaveFavoritos(favs FavoritosMap) error {
	filePath, err := getFilePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(favs, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}
