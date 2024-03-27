package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func init() {
	fmt.Println("init")
}

func main() {
	url := "postgres://dxrrvvah:2RyYR4ZuPvwebT8E8q6gS0qMZkJc0PKT@john.db.elephantsql.com/dxrrvvah"

	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("connect to database error", err)
	}
	defer db.Close()

	createTb := `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, age INT)`

	_, err = db.Exec(createTb)
	if err != nil {
		fmt.Println("Error creating table: ", err)
	}

	log.Println("connect to database")
}
