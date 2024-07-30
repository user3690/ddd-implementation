package full

import "ddd-implementation/pkg/offersearch/domain/value"

type Offer struct {
	Id         uint
	SupplierId uint
	HotelId    uint
	HotelName  string
	City       value.City
	Price      uint
}
