package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var shops []Shop

// our main function
func main() {

	shops = []Shop{}
	shops = append(shops, Shop{ID: 0, UniqueName: "car-workshop", Name: "Car Workshop", ExtendedName: "", Address: "1 Colcester Road", PostalCode: "ABC 123", Category: "Cars"})
	shops = append(shops, Shop{ID: 1, UniqueName: "clarks", Name: "Clarks", ExtendedName: "", Address: "2 Oxford", PostalCode: "234 AAA", Category: "Shoes"})
	shops = append(shops, Shop{ID: 2, UniqueName: "morrisons", Name: "Morrisons", ExtendedName: "", Address: "3 Regents", PostalCode: "142 WWW", Category: "Groceries"})
	log.Printf("\nMain: %#v\n", shops)

	router := mux.NewRouter()
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/shop", GetShops).Methods("GET")
	v1.HandleFunc("/shop/{id:[a-zA-Z-]}", GetShop).Methods("GET")
	v1.HandleFunc("/shop/{id:[a-zA-Z-]}", CreateShop).Methods("POST")
	v1.HandleFunc("/shop/{id:[a-zA-Z-]}", DeleteShop).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// GetShops -
func GetShops(w http.ResponseWriter, r *http.Request) {
	log.Printf("\nShops: %#v\n", shops)
	json.NewEncoder(w).Encode(shops)
}

// GetShop -
func GetShop(w http.ResponseWriter, r *http.Request) {}

// CreateShop -
func CreateShop(w http.ResponseWriter, r *http.Request) {}

// DeleteShop -
func DeleteShop(w http.ResponseWriter, r *http.Request) {}

// Shop -
type Shop struct {
	ID           int    `json:"id"`
	UniqueName   string `json:"uniqueName"`
	Name         string `json:"name"`
	ExtendedName string `json:"extendedName"`
	Address      string `json:"address"`
	PostalCode   string `json:"postalCode"`
	Category     string `json:"category"`
}
