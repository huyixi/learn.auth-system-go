package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB connection successful")

	createTable()
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table created successfully")
}

func CloseDB() {
	err := DB.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB connection closed")
}
