package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	g.GET("/users", getUserHandler)
	g.POST("/users", createUserHandler)
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

func healthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
