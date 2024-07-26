package repository

import (
	"awesomeProject/pkg/offersearch/domain/entity/full"
)

type Offer interface {
	GetAllOffers() ([]full.Offer, error)
	SaveOffers([]full.Offer) (int, error)
}
