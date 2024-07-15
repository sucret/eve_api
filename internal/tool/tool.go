package tool

import (
	"os"
	"path/filepath"
)

func GetRootDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	rootDir, err := filepath.Abs(wd)
	if err != nil {
		return "", err
	}
	return rootDir, nil
}
