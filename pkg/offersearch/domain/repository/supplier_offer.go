package repository

import (
	"awesomeProject/pkg/offersearch/domain/entity/partial"
	"time"
)

type SupplierOffer interface {
	GetOffers(hotelId uint, from time.Time, to time.Time) ([]partial.Offer, error)
}
