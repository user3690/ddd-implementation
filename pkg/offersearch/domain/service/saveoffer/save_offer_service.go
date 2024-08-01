package saveoffer

import (
	"ddd-implementation/pkg/offersearch/domain/entity/full"
	"ddd-implementation/pkg/offersearch/infrastructure/mysql/ods"
)

type Service interface {
	SaveOffers(offers []full.Offer) ([]full.Offer, error)
}

type Impl struct {
	odsOfferRepo ods.OfferRepository
}

func NewService(odsOfferRepo ods.OfferRepository) Impl {
	return Impl{
		odsOfferRepo: odsOfferRepo,
	}
}

func (i Impl) SaveOffers(offers []full.Offer) ([]full.Offer, error) {
	return i.odsOfferRepo.SaveOffers(offers)
}
