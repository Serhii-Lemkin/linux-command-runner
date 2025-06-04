package helpers

import (
	"encoding/json"
	"os"
	"path/filepath"
	"rnnr/classes"
	"rnnr/detectors"
	"rnnr/logger"
)

var aliasLocation string = ".rnnr/aliases.json"

func getAliasLocation() (string, error) {
	home, err := os.UserHomeDir()
	path := filepath.Join(home, aliasLocation)
	return path, err
}

func OpenAliases() {
	fileViewer, err := detectors.DetectEditor()
	if err != nil {
		logger.LogError(err)
	}

	path, err := getAliasLocation()
	if err != nil {
		logger.LogError(err)
	}

	Run(fileViewer + " " + path)
}

func LoadAliases() (map[string]classes.Alias, error) {
	aliases := make(map[string]classes.Alias)
	path, err := getAliasLocation()
	if err != nil {
		logger.LogError(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		err := initAliasesFile(path)
		if err != nil {
			logger.Log(err)
		} else {
			return aliases, nil
		}
	}

	err = json.Unmarshal(data, &aliases)
	if err != nil {
		logger.LogError(err)
	}

	return aliases, err
}

func SaveAliases(aliases map[string]classes.Alias) error {
	data, err := json.MarshalIndent(aliases, "", "  ")
	if err != nil {
		return err
	}

	path, err := getAliasLocation()
	return os.WriteFile(path, data, 0644)
}

func initAliasesFile(path string) error {
	dir := filepath.Dir(path)
	defaultAliases := make(map[string]classes.Alias)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(defaultAliases, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	logger.Log("Aliases file was successfully created at", path)
	return nil
}
