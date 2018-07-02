package system

import (
	"database/sql"
	"log"
)

var db *sql.DB

// GetDB returns pointer to database connection
func GetDB() *sql.DB {
	if db != nil {
		return db
	}
	var err error
	db, err = sql.Open("sqlite3", "./data/sqlite.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
