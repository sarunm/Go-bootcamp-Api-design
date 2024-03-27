package main

import (
	"database/sql"
	"fmt"
)

func main() {

	url := "postgres://dxrrvvah:2RyYR4ZuPvwebT8E8q6gS0qMZkJc0PKT@john.db.elephantsql.com/dxrrvvah"

	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println("connect to database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM users WHERE id=$1")
	if err != nil {
		fmt.Println("can't prepare query", err)
	}

	if _err := stmt.QueryRow(1); _err != nil {
		fmt.Println("can't delete row", _err)
	}

	fmt.Println("delete row success")
}
