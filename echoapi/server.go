package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Logger struct {
	Handler http.Handler
}

type Err struct {
	Message string `json:"message"`
}

var User = []Users{
	{ID: 1, Name: "John Doe", Age: 25},
	//{ID: 2, Name: "Jane Doe", Age: 26},
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.Handler.ServeHTTP(w, r)
	log.Printf("Server http middleware %s %s %s %s", r.RemoteAddr, r.Method, r.URL.Path, time.Since(start))

}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	g := e.Group("/api")
	g.Use(middleware.BasicAuth(func(u, p string, context echo.Context) (bool, error) {
		if u == "apidesign" && p == "45678" {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/users", getUsersFromDB)
	g.POST("/users", createUserToDbHandler)
	g.GET("/health", healthHandler)
	e.Start(":8080")

	log.Println("server started at :8080")
	//log.Fatal(srv.ListenAndServe())
	log.Println("server stopped")
}

func getUserHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, User)
}

func createUserHandler(c echo.Context) error {
	u := Users{}
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	User = append(User, u)
	return c.JSON(http.StatusCreated, u)
}

func createUserToDbHandler(c echo.Context) error {

	u := Users{}
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	url := "postgres://dxrrvvah:2RyYR4ZuPvwebT8E8q6gS0qMZkJc0PKT@john.db.elephantsql.com/dxrrvvah"
	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println("connect to database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println("can't prepare query", err)
	}

	row := stmt.QueryRow(u.Name, u.Age)

	err = row.Scan(&u.ID)
	if err != nil {
		fmt.Println("can't scan id", err)

	}
	return c.JSON(http.StatusOK, u)
}

func getUsersFromDB(c echo.Context) error {
	url := "postgres://dxrrvvah:2RyYR4ZuPvwebT8E8q6gS0qMZkJc0PKT@john.db.elephantsql.com/dxrrvvah"
	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println("connect to database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM users")
	if err != nil {
		fmt.Println("can't prepare query", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("can't query all rows ", err)
	}

	var users = []Users{}

	if err != nil {
		fmt.Println("can't scan row into variables", err)
	}
	for rows.Next() {

		var u Users
		err = rows.Scan(&u.ID, &u.Name, &u.Age)

		//err = rows.Scan(&user)
		if err != nil {
			fmt.Println("can't scan row into variables", err)
		}
		fmt.Println("id", u.ID, "name", u.Name, "age", u.Age)

		users = append(users, u)
	}

	return c.JSON(http.StatusOK, users)
}

func healthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
