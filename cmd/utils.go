package cmd

import (
	"os"
	"path/filepath"

	"github.com/karstenpedersen/sheet/db"
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

	dbPath := filepath.Join(dataDir, "sheet.db")
	err = db.InitDatabase(dbPath)
	if err != nil {
		return err
	}

	return nil
}
