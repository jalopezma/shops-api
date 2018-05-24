package providers

import (
	"github.com/jalopezma/shops-api/schemas"
)

// ShopsMockProvider is a ShopService implementation
type ShopsMockProvider struct {
}

// Create a shop
func (s *ShopsMockProvider) Create(shop schemas.Shop) {
}

// Get a shop by name
func (s *ShopsMockProvider) Get(name string) *schemas.Shop {
	shop := schemas.Shop{ID: 0, UniqueName: "car-workshop", Name: "Car Workshop", ExtendedName: "", Address: "1 Colcester Road", PostalCode: "ABC 123", Category: "Cars"}
	return &shop
}
