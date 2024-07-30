package repository

import "ddd-implementation/pkg/offersearch/domain/entity/full"

type Offer interface {
	GetAllOffers() ([]full.Offer, error)
	SaveOffers([]full.Offer) (int, error)
}
