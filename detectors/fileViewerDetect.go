package detectors

import (
	"fmt"
	"os/exec"
	"rnnr/config"
)

func DetectEditor() (string, error) {
	configuration, err := config.GetConfig()
	defaultFileViewer := configuration.PreferedFileManager

	if err != nil || defaultFileViewer == "" {
		for _, editor := range configuration.EditorList {
			if path, err := exec.LookPath(editor); err == nil {
				return path, nil
			}
		}
	} else {
		for _, editor := range configuration.EditorList {
			if defaultFileViewer == editor {
				if path, err := exec.LookPath(editor); err == nil {
					return path, nil
				}
			}
		}
	}

	return "", fmt.Errorf("no supported text editor found, try running <rnnr initconfig> and if id doesnt work after try adding your text editor in config < rnnr config >")
}
