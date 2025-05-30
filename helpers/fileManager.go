package helpers

import (
	"encoding/json"
	"os"
	"rnnr/classes"
)

var aliasLocation string = "/home/serhii/.rnnr/aliases.json"

func LoadAliases() (map[string]classes.Alias, error) {
	aliases := make(map[string]classes.Alias)
	data, err := os.ReadFile(aliasLocation)
	if err == nil {
		err = json.Unmarshal(data, &aliases)
	}

	return aliases, err
}

func SaveAliases(aliases map[string]classes.Alias) error {
	data, err := json.MarshalIndent(aliases, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(aliasLocation, data, 0644)
}
