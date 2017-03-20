package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dimfeld/httptreemux"
)

type UpsertCarHandler struct{}

func (h *UpsertCarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
	fmt.Fprintf(w, "Eu deveria criar um carro chamado: %s!", params["id"])
	fmt.Fprintln(w, "Não crio por que sou mal!")
}

type GetCarHandler struct{}

func (h *GetCarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
	fmt.Fprintf(w, "Eu deveria busca um carro chamado: %s!", params["id"])
	fmt.Fprintln(w, "Não busco por que estou com preguiça!")
}

func main() {
	addr := "127.0.0.1:8081"
	router := httptreemux.NewContextMux()
	router.Handler(http.MethodGet, "/cars/:id", &GetCarHandler{})
	router.Handler(http.MethodPut, "/cars/:id", &UpsertCarHandler{})

	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))

	// execute
	// curl http://localhost:8081/cars/gol
	// curl -XPUT http://localhost:8081/cars/fusca -d'{"name": 1}'
}
