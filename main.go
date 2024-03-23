package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", handle)

    log.Println("server started at :8080")
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
    log.Println("server stopped")
}

func handle(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        w.WriteHeader(http.StatusCreated)
        w.Write([]byte(`"test POST"`))
        return
    }

    w.Write([]byte(`{"word":"Hello, World!!!"}`))
}
