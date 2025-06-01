package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func RunByAlias() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: rnnr run <name>")
		return
	}

	name := os.Args[2]

	aliases, err := LoadAliases()
	if err != nil {
		LogError(err)
		return
	}

	alias, exists := aliases[name]
	if !exists {
		fmt.Println("Alias not found:", name)
		return
	}

	terminal, args, err := DetectTerminal()

	if err != nil {
		Log("No supported terminals found")
	}

	for _, command := range alias.Commands {
		args = append(args, command, "; exec bash")
		cmd := exec.Command(terminal, args...)
		cmd.Stdout = nil
		cmd.Stderr = nil
		cmd.Stdin = nil
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setpgid: true,
		}

		if err := cmd.Start(); err != nil {
			LogError(err)
			return
		}

		Log("Launched in new terminal with PID", cmd.Process.Pid)
	}
}
