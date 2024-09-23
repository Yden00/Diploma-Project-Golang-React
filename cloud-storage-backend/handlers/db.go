package handlers

import (
	"database/sql"
	"log"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func GetDB() *sql.DB {
	if db == nil {
		log.Fatal("Database connection is not initialized")
	}
	return db
}
