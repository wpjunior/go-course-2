package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	person := &Person{
		FirstName: "Wilson",
		LastName:  "Júnior",
		Age:       24,
	}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	addr := "127.0.0.1:8081"
	handler := &MyHandler{}
	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
