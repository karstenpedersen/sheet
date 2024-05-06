package db

import (
	"database/sql"
	"os"
	"path/filepath"
)

func getDatabasePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dbPath := filepath.Join(homeDir, "/.sheet/sheet.db")
	return dbPath, nil
}

func openDatabase() (*sql.DB, error) {
	dbPath, err := getDatabasePath()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}
