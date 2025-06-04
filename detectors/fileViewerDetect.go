package detectors

import (
	"fmt"
	"os/exec"
	"rnnr/classes"
	"rnnr/config"
)

var editorList = []classes.Editor{
	{Name: "nvim"},
	{Name: "gedit"},
	{Name: "nano"},
	{Name: "vi"},
}

func DetectEditor() (string, error) {
	configuration, err := config.GetConfig()
	defaultFileViewer := configuration.PreferedFileManager
	if err != nil || defaultFileViewer == "" {
		for _, editor := range editorList {
			if path, err := exec.LookPath(editor.Name); err == nil {
				return path, nil
			}
		}
	} else {
		for _, editor := range editorList {
			if defaultFileViewer == editor.Name {
				if path, err := exec.LookPath(editor.Name); err == nil {
					return path, nil

				}
			}
		}
	}
	return "", fmt.Errorf("no supported text editor found")
}
