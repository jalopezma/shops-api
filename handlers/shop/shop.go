package shop

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jalopezma/shops-api/schemas"
)

var shops []schemas.Shop

/*
shops := []schemas.Shop{}
shops = append(shops, schemas.Shop{ID: 0, UniqueName: "car-workshop", Name: "Car Workshop", ExtendedName: "", Address: "1 Colcester Road", PostalCode: "ABC 123", Category: "Cars"})
shops = append(shops, schemas.Shop{ID: 1, UniqueName: "clarks", Name: "Clarks", ExtendedName: "", Address: "2 Oxford", PostalCode: "234 AAA", Category: "Shoes"})
shops = append(shops, schemas.Shop{ID: 2, UniqueName: "morrisons", Name: "Morrisons", ExtendedName: "", Address: "3 Regents", PostalCode: "142 WWW", Category: "Groceries"})
log.Printf("\nHandlers shop: %#v\n", shops)
*/

// GetShops -
func Shops(w http.ResponseWriter, r *http.Request) {
	log.Printf("\nShops: %#v\n", shops)
	json.NewEncoder(w).Encode(shops)
}

// GetShop -
func Shop(w http.ResponseWriter, r *http.Request) {}

// CreateShop -
func CreateShop(w http.ResponseWriter, r *http.Request) {}

// DeleteShop -
func DeleteShop(w http.ResponseWriter, r *http.Request) {}
