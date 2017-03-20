package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello warriors")
}

func main() {
	addr := "127.0.0.1:8081"
	handler := &MyHandler{}
	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
