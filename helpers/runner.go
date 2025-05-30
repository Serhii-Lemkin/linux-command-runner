package helpers

import (
	"fmt"
	"os"
	"os/exec"
)

func RunByAlias() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: rnnr run <name>")
		return
	}

	name := os.Args[2]

	aliases, err := LoadAliases()
	if err != nil {
		LogError(err.Error())
		return
	}

	alias, exists := aliases[name]
	if !exists {
		fmt.Println("Alias not found:", name)
		return
	}

	terminal := "alacritty"
	var lastCmd *exec.Cmd
	for i, c := range alias.Commands {
		fmt.Printf("→ %s\n", c)
		//cmd := exec.Command(terminal, "bash", "-c", c)
		cmd := exec.Command(terminal, "-e", "bash", "-c", c+"; exec bash")

		if err := cmd.Run(); err != nil {
			fmt.Println("Command failed:", err)
			return
		}

		if i == len(alias.Commands)-1 {
			lastCmd = cmd // remember the last command
		}
	}

	if lastCmd != nil {
		err := lastCmd.Wait()
		if err != nil {
			fmt.Println(err)
		}
	}
}
