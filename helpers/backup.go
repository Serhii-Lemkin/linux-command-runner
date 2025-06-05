package helpers

import (
	"io"
	"os"
	"path/filepath"
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
	//To Do: instead of copy paste copy the aliases to the dict.
	if len(os.Args) != 3 {
		logger.Log("the usage is <rnnr importaliases your-location-here>")
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

	err = copyFile(fullPath, path)
	if err != nil {
		logger.Log("the aliases file was not imported")
		logger.LogError(err)
	} else {
		logger.Log("the aliases file was imported from", fullPath)
	}
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
