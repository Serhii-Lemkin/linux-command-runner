package helpers

import (
	"fmt"
)

func PrintHelp() {
	fmt.Println("You can save the commands to different aliases. By default it runs in the same terminal window. Runnung multiple commands at ones might need to be opened in separate windows. There are parameters in <run> command to change that.")
	fmt.Println("1. Basic useCases")
	PrintDetailed()
	fmt.Println("2. Config and Files")
	PrintConfigAndFilesDoc()
}

func PrintConfigAndFilesDoc() {
	fmt.Println("openaliases                  => open aliases file in file editor")
	fmt.Println("config                       => open config file where you can choose some options")
	fmt.Println("initconfig                   => init config file if none is found in the right location or restoreDefaults")
}

func PrintDetailed() {
	fmt.Println("a, add <alias> <command>     => create command or add to existing alias")
	fmt.Println("r, run <alias>               => run command or commands saved to alias. Use <-keep> at the end to keep new terminal alive use <-detach> to open in new terminal window or <-here> to run in the same window. No parameter runs the default in your config <false>")
	fmt.Println("la, listall                  => list all existing aliases")
	fmt.Println("s, show <alias>              => show what commands are saved to this alias")
	fmt.Println("d, delete <alias>            => delete already existing aliases <-y> to autoconfirm")
	fmt.Println("h, help                      => get help")
	fmt.Println("fuzzy, f                     => fuzzy find commands by Alias, uses contains <-alias> or <-a> to search only in names <-command> <-c> to search in commands, the default searches in both")
	fmt.Println("rename <oldalias> <newAlias> => rename existing aliases")
}

func PrintBasic() {
	fmt.Println("Most Common UseCases, use < rnnr help > for more")
	fmt.Println("a, add <alias> <command>     => create command or add to existing alias")
	fmt.Println("r, run <alias>               => run command or commands saved to alias.")
	fmt.Println("la, listall                  => list all existing aliases")
	fmt.Println("s, show <alias>              => show what commands are saved to this alias")
	fmt.Println("d, delete <alias>            => delete already existing aliases")
	fmt.Println("h, help                      => get help")
	fmt.Println("rename                       => rename existing aliases")
	fmt.Println("fuzzy, f                     => fuzzy find commands by Alias")
	fmt.Println("rename <oldalias> <newAlias> => rename existing aliases")
}
