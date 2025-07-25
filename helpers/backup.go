package helpers

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"rnnr/classes"
	"rnnr/config"
	"rnnr/logger"
	"strings"
)

func ExportConfig() {
	if len(os.Args) != 3 {
		logger.Log("the usage is <rnnr exportconfig your-location-here>")
	}
	path, err := config.GetConfigLocation()

	if err != nil {
		logger.LogError(err)
	}

	filename := "config.json"
	fullPath := os.Args[2]
	if !strings.HasSuffix(fullPath, filename) {
		fullPath = filepath.Join(fullPath, filename)
	}

	err = copyFile(path, fullPath)
	if err != nil {
		logger.Log("the config was not exported")
		logger.LogError(err)
	} else {
		logger.Log("the config file was exported to", fullPath)
	}
}

func ExportAliases() {
	if len(os.Args) != 3 {
		logger.Log("the usage is <rnnr exportaliases your-location-here>")
	}
	path, err := getAliasLocation()

	if err != nil {
		logger.LogError(err)
	}

	filename := "aliases.json"
	fullPath := os.Args[2]
	if !strings.HasSuffix(fullPath, filename) {
		fullPath = filepath.Join(fullPath, filename)
	}

	err = copyFile(path, fullPath)
	if err != nil {
		logger.Log("the aliases file was not exported")
		logger.LogError(err)
	} else {
		logger.Log("the aliases file was exported to", fullPath)
	}
}

func ImportAliases() {
	if len(os.Args) != 3 {
		logger.Log("the usage is <rnnr importaliases your-location-here>")
	}

	filename := "aliases.json"
	fullPath := os.Args[2]
	if !strings.HasSuffix(fullPath, filename) {
		fullPath = filepath.Join(fullPath, filename)
	}

	data, err := os.ReadFile(fullPath)
	if err != nil {
		logger.LogError(err)
	}

	newAliases := make(map[string]classes.Alias)
	err = json.Unmarshal(data, &newAliases)
	if err != nil {
		logger.LogError(err)
		return
	}

	oldAliases, err := LoadAliases()
	if err != nil {
		logger.Log(err)
		return
	}

	countImported := 0
	for name, newAlias := range newAliases {
		_, exists := oldAliases[name]
		if exists {
			logger.Log("alias <", name, "> exists in imported file and local.")
			replace := logger.GetUserConfirmation("Do you want to replace existing alias with new?")
			if replace {
				oldAliases[name] = newAlias
				countImported++
			}
		} else {
			oldAliases[name] = newAlias
			countImported++
		}
	}

	SaveAliases(oldAliases)
	logger.Log("imported", countImported, "aliases")
}

func ImportConfig() {
	if len(os.Args) != 3 {
		logger.Log("the usage is <rnnr importaliases your-location-here>")
	}
	path, err := config.GetConfigLocation()

	if err != nil {
		logger.LogError(err)
	}

	filename := "config.json"
	fullPath := os.Args[2]
	if !strings.HasSuffix(fullPath, filename) {
		fullPath = filepath.Join(fullPath, filename)
	}

	err = copyFile(fullPath, path)
	if err != nil {
		logger.Log("the config file was not imported")
		logger.LogError(err)
	} else {
		logger.Log("the config file was imported from", fullPath)
	}
}

func copyFile(soursePath string, destinationPath string) error {
	srcFile, err := os.Open(soursePath)
	if err != nil {
		return err
	}

	defer srcFile.Close()

	dstFile, err := os.Create(destinationPath)

	if err != nil {
		return err
	}

	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
