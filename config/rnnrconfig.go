package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"rnnr/classes"
	"rnnr/helpers"
)

var configLocation string = ".rnnr/config.json"

func GetConfig() (*classes.Config, error) {
	home, err := os.UserHomeDir()
	path := filepath.Join(home, configLocation)
	var config classes.Config
	if err != nil {
		helpers.Log(err)
	}

	data, err := os.ReadFile(path)
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

func ShowConfig() {
	config, err := GetConfig()
	if err != nil {
		helpers.LogError(err)
	}

	fmt.Println(config)
}

func initConfig(path string) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	defaultCfg := classes.Config{}

	data, err := json.MarshalIndent(defaultCfg, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	helpers.Log("Config file was successfully created at", path)
	return nil
}
