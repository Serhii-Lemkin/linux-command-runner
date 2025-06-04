package utils

import (
	"os"
	"os/exec"
	"rnnr/detectors"
	"rnnr/logger"
)

func ShowConfig(path string) {
	fileViewer, err := detectors.DetectEditor()
	if err != nil {
		logger.LogError(err)
	}

	fullCommand := fileViewer + " " + path
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "bash"
	}

	cmd := exec.Command(shell, "-c", fullCommand)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		logger.LogError(err)
	}
}
