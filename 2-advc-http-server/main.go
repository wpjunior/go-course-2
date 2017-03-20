package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:         "127.0.0.1:8081",
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Second,
	}

	log.Printf("Running web server on: http://%s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
