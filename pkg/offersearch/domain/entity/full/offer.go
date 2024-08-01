package full

import "ddd-implementation/pkg/offersearch/domain/value"

// Offer is a mutable entity.
// That is why the struct fields are exported.
type Offer struct {
	Id         uint
	SupplierId uint
	HotelId    uint
	HotelName  string
	City       value.City
	Price      uint
}
