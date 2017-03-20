package main

import (
	"log"
	"net/http"
)

func main() {
	addr := "127.0.0.1:8081"
	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
