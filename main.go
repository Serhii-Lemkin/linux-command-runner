package main

import (
	"fmt"
	"os"
	"rnnr/config"
	"rnnr/helpers"
	"rnnr/logger"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		helpers.PrintBasic()
		return
	}

	cmd := strings.ToLower(os.Args[1])

	switch cmd {
	case "add", "a":
		helpers.CreateAlias()
	case "run", "r":
		helpers.RunByAlias()
	case "listall", "la":
		helpers.ListAll()
	case "show", "s":
		helpers.ShowSpecificAlias()
	case "help", "h":
		helpers.PrintHelp()
	case "delete", "d":
		helpers.DeleteAliases()
	case "config":
		path, err := config.GetConfigLocation()
		if err != nil {
			logger.LogError(err)
		} else {
			helpers.OpenConfigFile(path)
		}
	case "openaliases":
		helpers.OpenAliases()
	case "initconfig":
		config.InitConfig()
	case "rename":
		helpers.RenameAlias()
	case "fuzzy", "f":
		helpers.FuzzyFindAliasByName()
	case "exportconfig":
		helpers.ExportConfig()
	case "exportaliases":
		helpers.ExportAliases()
	case "importconfig":
		helpers.ImportConfig()
	case "importaliases":
		helpers.ImportAliases()
	default:
		fmt.Println("Unknown command:", cmd)
	}
}
