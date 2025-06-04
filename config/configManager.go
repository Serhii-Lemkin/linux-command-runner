package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"rnnr/classes"
	"rnnr/logger"
	"rnnr/utils"
)

var configLocation string = ".rnnr/config.json"

func GetConfigLocation() (string, error) {
	home, err := os.UserHomeDir()
	path := filepath.Join(home, configLocation)
	return path, err
}

func GetConfig() (*classes.Config, error) {
	path, err := GetConfigLocation()
	if err != nil {
		logger.LogError(err)
	}

	var config classes.Config
	if err != nil {
		logger.Log(err)
	}

	data, err := utils.ReadFile(path)

	if err != nil {
		err = initConfig(path)
		return &config, err
	}

	err = json.Unmarshal(data, &config)

	if err != nil {
		initConfig(path)
	}

	return &config, err
}

func InitConfig() error {
	path, err := GetConfigLocation()
	if err != nil {
		logger.LogError(err)
	}

	return initConfig(path)
}

func initConfig(path string) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	defaultCfg := getDefaultConfig()
	data, err := json.MarshalIndent(defaultCfg, "", "  ")
	if err != nil {
		return err
	}

	err = utils.WriteFile(data, path)
	if err != nil {
		return err
	}

	logger.Log("Config file was successfully created at", path)
	return nil
}
