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

	stmt, err := db.Prepare("SELECT * FROM users")
	if err != nil {
		fmt.Println("can't prepare query ", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("can't query all rows ", err)
	}

	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println("can't scan row into variables", err)
		}
		fmt.Println("id", id, "name", name, "age", age)
	}

	fmt.Println("query all rows success")

}
