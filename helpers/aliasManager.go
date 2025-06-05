package helpers

import (
	"fmt"
	"os"
	"rnnr/classes"
	"rnnr/logger"
	"strings"
)

func DeleteAliases() {
	aliases, err := LoadAliases()
	if err != nil {
		logger.LogError(err)
		return
	}

	aliasToDelete, exists := aliases[os.Args[2]]
	if !exists {
		logger.Log("No such alias found")
		return
	}

	verify := true
	if len(os.Args) >= 4 && os.Args[3] == "-y" {
		verify = false
	}

	if verify {
		if logger.GetUserConfirmation("Are you sure?") {
			deleteInner(aliases, aliasToDelete)
		} else {
			logger.Log("Delete aborted")
		}

	} else {
		deleteInner(aliases, aliasToDelete)
	}
}

func deleteInner(aliases map[string]classes.Alias, aliasToDelete classes.Alias) {
	delete(aliases, os.Args[2])
	SaveAliases(aliases)
	logger.LogAlias("The next command was deleted \n", aliasToDelete, os.Args[2])
}

func ShowSpecificAlias() {
	aliases, err := LoadAliases()
	if err != nil {
		logger.LogError(err)
	}

	alias, exists := aliases[os.Args[2]]
	commands := alias.Commands

	if !exists {
		logger.Log("No such alias found")
		return
	}

	for i, command := range commands {
		fmt.Println(i, command)
	}
}

func ListAll() {
	aliases, err := LoadAliases()
	if err != nil {
		logger.LogError(err)
		return
	}

	if len(aliases) == 0 {
		fmt.Println("No aliases saved.")
		return
	}

	fmt.Println("Saved aliases:")
	for name, a := range aliases {
		fmt.Printf("  %s:\n", name)
		for i, c := range a.Commands {
			fmt.Printf("    %d. %s\n", i+1, c)
		}
	}
}

func CreateAlias() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: rnnr create <name> <command>")
		return
	}

	name := os.Args[2]
	command := strings.Join(os.Args[3:], " ")

	aliases, err := LoadAliases()
	if err != nil {
		logger.LogError(err)
		aliases = map[string]classes.Alias{}
	}

	if _, exists := aliases[name]; !exists {
		aliases[name] = classes.Alias{Commands: []string{command}}
		fmt.Printf("Created alias '%s' with 1 command\n", name)
	} else {
		slice := append(aliases[name].Commands, command)
		aliases[name] = classes.Alias{Commands: slice}
		fmt.Printf("Appended to alias '%s': %q\n", name, command)
	}

	if err := SaveAliases(aliases); err != nil {
		logger.LogError(err)
		return
	}
}

func RenameAlias() {
	if len(os.Args) < 4 {
		logger.Log("The usage is rnnr rename <old name> <new name>")
	}

	aliases, err := LoadAliases()
	if err != nil {
		logger.LogError(err)
	}
	oldName := os.Args[2]
	newName := os.Args[3]
	oldAlias, exists := aliases[oldName]
	if !exists {
		logger.Log("No such alias found")
		return
	}

	_, exists = aliases[newName]

	if exists {
		logger.Log("The name is already taken")
		return
	}

	commands := oldAlias.Commands
	deleteInner(aliases, oldAlias)
	newAlias := classes.Alias{Commands: commands}
	aliases[newName] = newAlias
	logger.LogAlias("New alias created instead: \n", newAlias, newName)
	SaveAliases(aliases)
}
