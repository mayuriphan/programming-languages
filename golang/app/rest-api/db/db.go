package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB //global variable to hold the database connection, can be used by outside modules

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to the database !")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createEventsTable := ` CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date DATETIME NOT NULL,
		user_id INTEGER NOT NULL
	);`
	_, err := DB.Exec(createEventsTable) // to execute statements
	if err != nil {
		panic("Could not create Event tables !")
	}
}
