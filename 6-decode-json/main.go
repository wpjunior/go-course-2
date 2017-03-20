package main

import (
	"encoding/json"
	"fmt"
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
	person := &Person{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Recebemos o json!")
	fmt.Fprintf(w, "first name: %s\n", person.FirstName)
	fmt.Fprintf(w, "last name: %s\n", person.LastName)
	fmt.Fprintf(w, "age: %d\n", person.Age)

	// Execute o comando abaixo após rodar
	// curl -i http://127.0.0.1:8081 -d'{"firstName": "Kaio", "lastName": "Vinicius", "age": 24}'
}

func main() {
	addr := "127.0.0.1:8081"
	handler := &MyHandler{}
	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
