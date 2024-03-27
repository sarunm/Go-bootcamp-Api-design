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

    stmt, err := db.Prepare("SELECT * FROM users where id=$1")
    if err != nil {
        fmt.Println("can't prepare query ", err)
    }

    rowId := 1
    row := stmt.QueryRow(rowId)
    var id int
    var name string
    var age int
    err = row.Scan(&id, &name, &age)
    if err != nil {
        fmt.Println("can't scan row into variables", err)
    }
    fmt.Println("id", id, "name", name, "age", age)
}
