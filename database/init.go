package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Initialize initialises the database
func Initialize(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil || db == nil {
		panic("Error connecting to database")
	}

	