package providers

import (
	"log"

	"github.com/jalopezma/shops-api/schemas"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ShopsService - Shops service interface
type ShopsService interface {
	Create(shop schemas.Shop)
	Get(name string) *schemas.Shop
}

// ShopsProvider is a ShopService implementation
type ShopsProvider struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

// Close method to close the session
func (s *ShopsProvider) Close() {
	s.Session.Close()
}

// Create a shop
func (s *ShopsProvider) Create(shop schemas.Shop) {

}

// Get a shop by name
func (s *ShopsProvider) Get(name string) *schemas.Shop {
	log.Printf("[ShopsProvider] Get %q", name)
	query := bson.M{}
	var res *schemas.Shop
	err := s.Collection.Find(query).One(&res)
	log.Printf("Result:\n%+v\n%+v", res, err)
	return res
}
