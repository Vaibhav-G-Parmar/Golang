package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// DB is the global database connection
var DB *sql.DB

// Initialize a database connection
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "events.db")

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Verify connection
	if err = DB.Ping(); err != nil {
		panic("Failed to ping database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

// Create necessary table
func createTables() {
	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		user_id INTEGER
	)`

	_, err := DB.Exec(createEventTable)

	if err != nil {
		panic("Failed to create events table: " + err.Error())
	}
}
