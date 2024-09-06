package database

import (
	"database/sql"

	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Open() *sql.DB {
	db, err := sql.Open("sqlite3", "./lottery.db")
	if err != nil {
		log.Fatal("Error opening database", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database", err)
	}
	createTable(db)

	return db
}

func createTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS tickets (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  ticket INTEGER NOT NULL,
  name text NOT NULL,
  createdAt DATETIME NOT NULL DEFAULT current_timestamp
  );`
	statment, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal("Error making database", err)
	}
	statment.Exec()
}
