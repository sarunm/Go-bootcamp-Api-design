package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	url := "postgres://dxrrvvah:2RyYR4ZuPvwebT8E8q6gS0qMZkJc0PKT@john.db.elephantsql.com/dxrrvvah"

	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println("connect to database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE users SET name=$2, age=$3 WHERE id=$1")
	if err != nil {
		fmt.Println("can't prepare query", err)
	}

	if _, err = stmt.Exec(1, "nick", 12); err != nil {
		fmt.Println("can't update row", err)
	}
	fmt.Println("update row success")

}
