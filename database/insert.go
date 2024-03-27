package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	//urls := os.Getenv("DATABASE_URL") // DATABASE_URL=postgres://dxrrvvah:2RyYR4ZuPvwebT8E8q6gS0qMZkJc0PKT@john.db.elephantsql.com/dxrrvvah go run insert.go
	url := "postgres://dxrrvvah:2RyYR4ZuPvwebT8E8q6gS0qMZkJc0PKT@john.db.elephantsql.com/dxrrvvah"

	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println("connect to database error", err)
	}
	defer db.Close()

	row := db.QueryRow("INSERT INTO users (name, age) values ($1, $2) RETURNING id", "nick", 11)

	var id int
	err = row.Scan(&id)
	if err != nil {
		fmt.Println("can't insert row:", err)
	}

	fmt.Println("inserted row id:", id)

}
