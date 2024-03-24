package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
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
    // multiplexer
    mux := http.NewServeMux()

    mux.HandleFunc("/users", logMiddleware(handler))
    mux.HandleFunc("/health", logMiddleware(healthHandler))

    logMux := Logger{Handler: mux}

    srv := http.Server{
        Addr:    ":8080",
        Handler: logMux,
    }

    log.Println("server started at :8080")
    log.Fatal(srv.ListenAndServe())
    log.Println("server stopped")
}

func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        b, err := json.Marshal(User)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(err.Error()))
            return
        }

        w.Write(b)
        return
    }

    if r.Method == http.MethodPost {
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            log.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        //w.WriteHeader(http.StatusCreated)
        var u Users
        err = json.Unmarshal(body, &u)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte(err.Error()))
            return
        }

        User = append(User, u)
        fmt.Fprintf(w, "Hello, %s", u.Name)
        //w.Write(b)
        return
    }

    //w.Write([]byte(`{"word":"Hello, World!!!"}`))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"status":"ok"}`))
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        //log.Println(r.Method, r.URL.Path)
        start := time.Now()
        log.Printf("Server http middleware %s %s %s %s", r.RemoteAddr, r.Method, r.URL.Path, time.Since(start))
        next.ServeHTTP(w, r)
    }
}
