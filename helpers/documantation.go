package helpers

import (
	"fmt"
)

func PrintDoc() {
	fmt.Println("a, add <alias> <command>  => create command or add to existing alias")
	fmt.Println("r, run <alias>            => run command or commands saved to alias. use < -keep > at the end to keep new terminal alive")
	fmt.Println("la, listall               => list all existing aliases")
	fmt.Println("s, show <alias>           => show what commands are saved to this alias")
	fmt.Println("d, delete <alias>         => delete already existing aliases -y to autoconfirm")
	fmt.Println("h, help                   => get help")
	fmt.Println("config                    => open config file where you can choose some options")
	fmt.Println("initconfig                => init config file if none is found in the right location")
}
