package detectors

import (
	"fmt"
	"os/exec"
	"rnnr/config"
)

func DetectTerminal() (string, []string, error) {
	configuration, err := config.GetConfig()
	defaultTerm := configuration.PreferedTerminal

	if err != nil || defaultTerm == "" {
		for _, term := range configuration.TerminalList {
			if path, err := exec.LookPath(term.Name); err == nil {
				return path, term.Args, nil
			}
		}

	}

	for _, term := range configuration.TerminalList {
		if term.Name == defaultTerm {
			if path, err := exec.LookPath(term.Name); err == nil {
				return path, term.Args, nil
			}
		}
	}

	return "", nil, fmt.Errorf("no supported terminal emulator found, try running <rnnr initconfig> and if id doesnt work after try adding your terminal in config < rnnr config >")
}
