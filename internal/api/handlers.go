package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tomkdickinson/hexagonal-cart-service/internal/domain/cart"
	"log"
	"net/http"
	"strconv"
)

type Handlers struct {
	service *cart.Service
}

func ProvideHandlers(service *cart.Service) *Handlers {
	return &Handlers{service: service}
}

func (a Handlers) GetCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if c, err := a.service.GetCart(params["id"]); err != nil {
		log.Print(err)
		sendErrorResponse(w)
	} else {
		response := newCartResponse(c)
		sendResponse(w, response)
	}
}

func (a Handlers) AddItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	sku := params["sku"]
	quantity, err := strconv.Atoi(params["quantity"])
	if err != nil {
		sendErrorResponse(w)
	}
	if c, err := a.service.AddItemToCart(id, sku, quantity); err != nil {
		log.Println(err)
		sendErrorResponse(w)
	} else {
		response := newCartResponse(c)
		sendResponse(w, response)
	}
}

func sendResponse(w http.ResponseWriter, response *Cart) {
	w.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&response)
	if err != nil {
		log.Print(fmt.Errorf("could not encode json for response: %w", err))
	}
}

func sendErrorResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadGateway)
}
