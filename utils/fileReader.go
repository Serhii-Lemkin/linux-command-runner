package utils

import "os"

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func WriteFile(data []byte, path string) error {
	mode := os.FileMode(0644)
	return os.WriteFile(path, data, mode)
}
