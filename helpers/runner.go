package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"rnnr/config"
	"rnnr/detectors"
	"rnnr/logger"
	"runtime"
	"slices"
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
		logger.LogError(err)
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
		logger.LogError(err)
		return
	}

	config, err := config.GetConfig()

	fullCommand := command
	keepOpen := slices.ContainsFunc(os.Args, func(s string) bool {
		return s == "-keep"
	})

	sameTerminal := config.DefaultSameTerminal == true

	sameTerminalParam := slices.ContainsFunc(os.Args, func(s string) bool {
		return s == "-here"
	})

	detachTerminalParam := slices.ContainsFunc(os.Args, func(s string) bool {
		return s == "-detach"
	})

	if sameTerminalParam && !detachTerminalParam {
		sameTerminal = true
	} else if !sameTerminalParam && detachTerminalParam {
		sameTerminal = false
	}

	if keepOpen {
		if runtime.GOOS == "windows" {
			fullCommand += "; pause"
		} else {
			fullCommand += "; exec bash"
		}
	}

	if sameTerminal {
		shell := os.Getenv("SHELL")
		var shellArgs []string

		switch runtime.GOOS {
		case "windows":
			shell = "powershell"
			shellArgs = []string{"-NoExit", "-Command", fmt.Sprintf("& { %s }", fullCommand)}
		default:
			shell = os.Getenv("SHELL")
			if shell == "" {
				shell = "bash"
			}

			shellArgs = []string{"-c", fullCommand}
		}

		cmd := exec.Command(shell, shellArgs...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		logger.Log("used", terminal, "to run => ", command)
		if err := cmd.Run(); err != nil {
			logger.LogError(err)
		}

		return
	} else {
		args = append(args, fullCommand)

		cmd := exec.Command(terminal, args...)
		cmd.Stdin = nil
		cmd.Stdout = nil
		cmd.Stderr = nil

		if runtime.GOOS != "windows" {
			cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
		}

		if err := cmd.Start(); err != nil {
			logger.LogError(err)
			return
		}

		logger.Log("Launched in new", terminal, "terminal with PID", cmd.Process.Pid)
	}
}
