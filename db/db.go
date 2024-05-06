package db

import (
	"log"

	"github.com/karstenpedersen/sheet/models"
	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase(dataSourceName string) error {
	db, err := openDatabase()
	if err != nil {
		return nil
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return err
	}
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS shortcuts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			shortcut TEXT NOT NULL,
			scope TEXT
		);
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	return nil
}

func InsertShortcut(shortcut models.Shortcut) error {
	db, err := openDatabase()
	if err != nil {
		return nil
	}
	defer db.Close()

	stmt, err := db.Prepare(`
		INSERT INTO shortcuts(title, description, shortcut, scope)
		VALUES(?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(shortcut.Title, shortcut.Description, shortcut.Shortcut, shortcut.Scope)
	if err != nil {
		return err
	}

	return nil
}

func SelectShortcuts() ([]models.Shortcut, error) {
	db, err := openDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT id, title, description, shortcut, scope
		FROM shortcuts;
	`)
	defer rows.Close()

	shortcuts := []models.Shortcut{}
	for rows.Next() {
		var sc models.Shortcut
		err := rows.Scan(&sc.Id, &sc.Title, &sc.Description, &sc.Scope, &sc.Shortcut)
		if err != nil {
			log.Fatal(err)
		}

		shortcuts = append(shortcuts, sc)
	}

	return shortcuts, nil
}
