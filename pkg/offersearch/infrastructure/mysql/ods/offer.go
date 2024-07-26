package ods

import "awesomeProject/pkg/offersearch/domain/entity/full"

type OfferRow struct {
	Id         uint `db:"id"`
	SupplierId uint `db:"supplier_id"`
	Price      uint `db:"price"`
}

type OfferRepository struct{}

func NewOfferRepository() OfferRepository {
	return OfferRepository{}
}

func (r OfferRepository) GetAllOffers() ([]full.Offer, error) {
	return nil, nil
}

func (r OfferRepository) SaveOffers(offers []full.Offer) (int, error) {
	offerRows := r.fullOfferToRow(offers)

	return r.saveOffers(offerRows)
}

func (r OfferRepository) fullOfferToRow(offers []full.Offer) (offerRows []OfferRow) {
	for _, offer := range offers {
		newOfferRow := OfferRow{}

		newOfferRow.Id = offer.Id
		newOfferRow.SupplierId = offer.SupplierId
		newOfferRow.Price = offer.Price

		offerRows = append(offerRows, newOfferRow)
	}

	return offerRows
}

func (r OfferRepository) saveOffers(offers []OfferRow) (int, error) {
	return len(offers), nil
}
