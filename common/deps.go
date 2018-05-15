package common

import "github.com/jalopezma/shops-api/providers"

// Dependencies - Struct with all dependencies used
type Dependencies struct {
	Shops providers.ShopsService
}

// Close - Closes all dependencies
func (d *Dependencies) Close() {

	// d.Shops.C
}

// Deps
var Deps Dependencies
