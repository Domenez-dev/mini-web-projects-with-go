package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var DB *sql.DB

func InitDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./awesome.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	_, err = DB.Exec(`
                CREATE TABLE IF NOT EXISTS todos (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                title TEXT,
                status TEXT
                );`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}

func CreateTodo(title, status string) (int64, error) {
	result, err := DB.Exec("INSERT INTO todos (title, status) VALUES (?, ?)", title, status)
	if err != nil {
		fmt.Println("Error inserting todo:", err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error getting last insert ID:", err)
		return 0, err
	}
	return id, nil
}

func DeleteTodo(id int64) error {
	_, err := DB.Exec("DELETE FROM todos WHERE id=?", id)
	if err != nil {
		fmt.Println("Error deleting todo:", err)
	}
	return err
}

func GetAllTodos() ([]Todo, error) {
	rows, err := DB.Query("SELECT id, title, status FROM todos")
	if err != nil {
		fmt.Println("Error querying todos:", err)
		return []Todo{}, err
	}
	defer rows.Close()

	todos := make([]Todo, 0)
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Status)
		if err != nil {
			fmt.Println("Error scanning todo:", err)
			return []Todo{}, err
		}
		todos = append(todos, todo)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("Error with rows:", err)
		return []Todo{}, err
	}
	return todos, nil
}
