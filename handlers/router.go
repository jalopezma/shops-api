package handlers

import (
	"github.com/gorilla/mux"
	"github.com/jalopezma/shops-api/handlers/shop"
)

// NewRouter - Initializes the router
func NewRouter() *mux.Router {

	router := mux.NewRouter()
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/shop", shop.Shops).Methods("GET")
	v1.HandleFunc("/shop/{id:[a-zA-Z-]{5,10}}", shop.Shop).Methods("GET")
	v1.HandleFunc("/shop/{id:[a-zA-Z-]}", shop.CreateShop).Methods("POST")
	v1.HandleFunc("/shop/{id:[a-zA-Z-]}", shop.DeleteShop).Methods("DELETE")

	return router
}
