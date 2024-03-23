package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var User = []Users{
	{ID: 1, Name: "John Doe", Age: 25},
}

func main() {
	http.HandleFunc("/", handle)

	log.Println("server started at :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
	log.Println("server stopped")
}

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusCreated)

		b, err := json.Marshal(User)
		if err != nil {
			log.Println(err)
			return
		}

		w.Write(b)
		return
	}

	w.Write([]byte(`{"word":"Hello, World!!!"}`))
}
