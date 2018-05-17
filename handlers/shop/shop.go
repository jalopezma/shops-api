package shop

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jalopezma/shops-api/common"
	"github.com/jalopezma/shops-api/schemas"
)

var shops []schemas.Shop

// Shops -
func Shops(w http.ResponseWriter, r *http.Request) {
	log.Printf("\nShops: %#v\n", shops)
	json.NewEncoder(w).Encode(shops)
}

// GetShop -
func Shop(w http.ResponseWriter, r *http.Request) {
	shop := common.Deps.Shops.Get("something")
	log.Printf("GetShop: %+v", shop)
}

// CreateShop -
func CreateShop(w http.ResponseWriter, r *http.Request) {}

// DeleteShop -
func DeleteShop(w http.ResponseWriter, r *http.Request) {}
