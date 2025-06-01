package helpers

import (
	"fmt"
	"os/exec"
)

var terminalCommands = map[string][]string{
	"alacritty":      {"-e", "bash", "-c"},
	"kitty":          {"bash", "-c"},
	"wezterm":        {"start", "--", "bash", "-c"},
	"gnome-terminal": {"--", "bash", "-c"},
	"xterm":          {"-e", "bash", "-c"},
}

func DetectTerminal() (string, []string, error) {
	for term, args := range terminalCommands {
		if path, err := exec.LookPath(term); err == nil {
			return path, args, nil
		}
	}

	return "", nil, fmt.Errorf("no supported terminal emulator found")
}
