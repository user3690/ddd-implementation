package repository

import "ddd-implementation/pkg/offersearch/domain/entity/full"

// Offer interface enables to separate external requirements like mysql database from business logic.
// Business logic should not have infrastructure dependencies.
type Offer interface {
	GetAllOffers() ([]full.Offer, error)
	SaveOffers([]full.Offer) ([]full.Offer, error)
}
