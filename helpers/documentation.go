package helpers

import (
	"fmt"
)

func PrintHelp() {
	fmt.Println("1. Basic useCases")
	PrintBasic()
	fmt.Println("rename                    => rename existing aliases")
	fmt.Println("2. Config and Files")
	fmt.Println("openaliases               => open aliases file in file editor")
	fmt.Println("config                    => open config file where you can choose some options")
	fmt.Println("initconfig                => init config file if none is found in the right location")
}

func PrintBasic() {
	fmt.Println("You can save the commands to different aliases. by default it runs in the new terminal window. There are parameters in < run > command to change that.")
	fmt.Println("Most Common UseCases, use < rnnr help > for more")
	fmt.Println("a, add <alias> <command>  => create command or add to existing alias")
	fmt.Println("r, run <alias>            => run command or commands saved to alias. Use < -keep > at the end to keep new terminal alive use < -detach > to open in new terminal window or < -here > to run in the same window. No parameter runs the default in your config <false>")
	fmt.Println("la, listall               => list all existing aliases")
	fmt.Println("s, show <alias>           => show what commands are saved to this alias")
	fmt.Println("d, delete <alias>         => delete already existing aliases -y to autoconfirm")
	fmt.Println("h, help                   => get help")
	fmt.Println("rename                    => rename existing aliases")
}
