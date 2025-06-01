package main

import (
	"fmt"
	"os"
	"rnnr/config"
	"rnnr/helpers"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		helpers.PrintDoc()
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
		helpers.PrintDoc()
	case "delete", "d":
		helpers.DeleteAliases()
	case "config":
		config.ShowConfig()
	case "openaliases":
		helpers.OpenAliases()
	case "initconfig":
		config.InitConfig()
	default:
		fmt.Println("Unknown command:", cmd)
	}
}
