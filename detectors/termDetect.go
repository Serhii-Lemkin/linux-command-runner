package detectors

import (
	"fmt"
	"os/exec"
	"rnnr/classes"
	"rnnr/config"
)

var terminalCommands = []classes.TerminalCommand{
	{Name: "ptyxis", Args: []string{"--", "bash", "-c"}},
	{Name: "alacritty", Args: []string{"-e", "bash", "-c"}},
	{Name: "kitty", Args: []string{"bash", "-c"}},
	{Name: "wezterm", Args: []string{"start", "--", "bash", "-c"}},
	{Name: "gnome-terminal", Args: []string{"--", "bash", "-c"}},
	{Name: "xterm", Args: []string{"-e", "bash", "-c"}},
}

func DetectTerminal() (string, []string, error) {
	configuration, err := config.GetConfig()
	defaultTerm := configuration.PreferedTerminal

	if err != nil || defaultTerm == "" {
		for _, term := range terminalCommands {
			if path, err := exec.LookPath(term.Name); err == nil {
				return path, term.Args, nil
			}
		}

	}

	for _, term := range terminalCommands {
		if term.Name == defaultTerm {
			if path, err := exec.LookPath(term.Name); err == nil {
				return path, term.Args, nil
			}
		}
	}

	return "", nil, fmt.Errorf("no supported terminal emulator found")
}
