package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	handlers, err := handlersWithFileStorage()

	if err != nil {
		log.Fatal(fmt.Errorf("could not wire application: %w", err))
	}

	router := mux.NewRouter()

	router.HandleFunc("/cart/{id}", handlers.GetCart)
	router.HandleFunc("/cart/{id}/{sku}/{quantity}", handlers.AddItem).Methods("GET")

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
	}
	log.Print("Starting server")
	log.Fatal(srv.ListenAndServe())
}
