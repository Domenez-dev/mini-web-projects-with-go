package main

import (
	"database/sql"
	"log"
)

type Todo struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

var DB *sql.DB

func InitDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./awesome.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`
                CREATE TABLE IF NOT EXISTS todos (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                title TEXT,
                status TEXT
                );`)
	if err != nil {
		log.Fatal("Database not created, ", err)
	}
}
