package main

import (
	_ "github.com/barshat7/go-store/data"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/categories", getCategories).Methods("GET")

	err := http.ListenAndServe(
		":4000",
		r,
	)
	log.Fatal(err)
}