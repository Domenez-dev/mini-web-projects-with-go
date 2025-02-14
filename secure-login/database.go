package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./awesome.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	_, err = DB.Exec(`
                CREATE TABLE IF NOT EXISTS users (
                username TEXT PRIMARY KEY,
                hashedPassword TEXT,
                sessionToken TEXT,
                CSRFToken TEXT
                );`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}
