package full

import "awesomeProject/pkg/offersearch/domain/value"

type Offer struct {
	Id         uint
	SupplierId uint
	HotelId    uint
	HotelName  string
	City       value.City
	Price      uint
}
