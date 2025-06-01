package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"rnnr/detectors"
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

	for _, command := range alias.Commands {
		Run(command)
	}
}

func Run(command string) {
	terminal, args, err := detectors.DetectTerminal()
	if err != nil {
		LogError(err)
		return
	}

	fullCommand := command
	if os.Args[len(os.Args)-1] == "-keep" {
		fullCommand += "; exec bash"
	}

	args = append(args, fullCommand)

	cmd := exec.Command(terminal, args...)
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	if err := cmd.Start(); err != nil {
		LogError(err)
		return
	}

	Log("Launched in new", terminal, "terminal with PID", cmd.Process.Pid)
}
