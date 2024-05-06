package cmd

import (
	"os"
	"path/filepath"
)

func getDataDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dataDir := filepath.Join(homeDir, "/.sheet/")
	return dataDir, nil
}

func createDataDir() error {
	dataDir, err := getDataDir()
	if err != nil {
		return err
	}
	err = os.MkdirAll(dataDir, 0755)
	if err != nil {
		return err
	}
	return nil
}
