package schemas

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
