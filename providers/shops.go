package providers

import (
	"github.com/jalopezma/shops-api/schemas"
	mgo "gopkg.in/mgo.v2"
)

// ShopsService - Shops service interface
type ShopsService interface {
	Create(shop schemas.Shop)
	Get(name string) *schemas.Shop
}

// ShopsProvider is a ShopService implementation
type ShopsProvider struct {
	collection *mgo.Collection
}

// Create a shop
func (s *ShopsProvider) Create(shop schemas.Shop) {

}

// Get a shop by name
func (s *ShopsProvider) Get(name string) *schemas.Shop {
	return nil
}
