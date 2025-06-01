package detectors

import (
	"fmt"
	"os/exec"
)

type Editor struct {
	Name string
}

var editorList = []Editor{
	{"nvim"},
	{"gedit"},
	{"nano"},
	{"vi"},
}

func DetectEditor() (string, error) {
	for _, editor := range editorList {
		if path, err := exec.LookPath(editor.Name); err == nil {
			return path, nil
		}
	}

	return "", fmt.Errorf("no supported text editor found")
}
