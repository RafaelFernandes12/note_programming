package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "file:/home/rafael/source/notesDb/notes.db?_foreign_keys=on&_journal_mode=WAL")
	if err != nil {
		return nil, err
	}
	return db, nil
}
