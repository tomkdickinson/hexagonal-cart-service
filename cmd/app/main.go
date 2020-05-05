package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tomkdickinson/hexagonal-cart-service/internal/api"
	"log"
	"net/http"
	"os"
)

func main() {
	var handlers *api.Handlers
	var err error

	switch os.Getenv("CART_STORAGE") {
	case "file":
		log.Println("Starting server with file repository")
		handlers, err = handlersWithFileStorage()
	default:
		log.Println("Starting server with mongo repository")
		handlers, err = handlersWithMongoStorage()
	}
	if err != nil {
		log.Fatal(fmt.Errorf("could not wire application: %w", err))
	}

	startServer(handlers)
}

func startServer(handlers *api.Handlers) {
	router := mux.NewRouter()
	router.HandleFunc("/cart/{id}", handlers.GetCart)
	router.HandleFunc("/cart/{id}/{sku}/{quantity}", handlers.AddItem).Methods("GET")

	srv := &http.Server{
		Handler: router,
		Addr:    ":8000",
	}
	log.Print("Starting server")
	log.Fatal(srv.ListenAndServe())
}
