package helpers

import (
	"fmt"
	"os"
	"rnnr/classes"
	"strings"
)

func DeleteAliases() {
	aliases, err := LoadAliases()
	if err != nil {
		LogError(err.Error())
		return
	}

	aliasToDelete, exists := aliases[os.Args[2]]
	if !exists {
		LogError("No such alias found")
		return
	}

	if GetUserConfirmation() {
		delete(aliases, os.Args[2])
		SaveAliases(aliases)
		LogAlias("The next command was deleted \n", aliasToDelete)
	} else {
		Log("Delete aborted")
	}
}

func ShowSpecificAlias() {
	aliases, err := LoadAliases()
	if err != nil {
		fmt.Println("Error loading aliases:", err)
	}

	alias, exists := aliases[os.Args[2]]
	commands := alias.Commands

	if !exists {
		LogError("No such alias found")
		return
	}

	for i, command := range commands {
		fmt.Println(i, command)
	}
}

func ListAll() {
	aliases, err := LoadAliases()
	if err != nil {
		fmt.Println("Error loading aliases:", err)
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
		fmt.Println("Error loading aliases:", err)
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
		fmt.Println("Error saving aliases:", err)
		return
	}
}

