package helpers

import (
	"fmt"
)

func PrintDoc() {
	fmt.Println("a, add <alias> <command> to create command or add to existing alias")
	fmt.Println("r, run <alias> to run command or commands saved to alias")
	fmt.Println("la, listall to list all existing aliases")
	fmt.Println("s, show <alias> to show what commands are saved to this alias")
	fmt.Println("d, delete <alias> to delete already existing aliases -y to autoconfirm")
	fmt.Println("h, help to get help")
}
