package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("start server")
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`Hello World`))
	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	fmt.Println("server started at :8080")

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	fmt.Println("shutting down server")
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Println("Shuting down error", err)
	}
	fmt.Println("server stopped")

}
