package Goutils

import (
	"os"
	"path/filepath"
)

// SelfDir returns the directory path of the current executable file.
//
// It uses `os.Executable()` to get the path of the current executable file,
// and then uses `filepath.Dir()` to extract the directory path.
//
// It returns the directory path as a string.
func SelfDir() (string, error) {
	dir, err := os.Executable()
	if err != nil {
		return "", err
	}
	exPath := filepath.Dir(dir)
	return exPath, nil
}
