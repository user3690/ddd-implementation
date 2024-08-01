package ods

import "ddd-implementation/pkg/offersearch/domain/entity/full"

type OfferRow struct {
	Id         uint   `db:"id"`
	SupplierId uint   `db:"supplier_id"`
	Price      uint   `db:"price"`
	CityName   string `db:"city_name"`
}

// OfferRepository named after the table, package named after database
type OfferRepository struct{}

func NewOfferRepository() OfferRepository {
	return OfferRepository{}
}

func (r OfferRepository) GetAllOffers() ([]full.Offer, error) {
	return nil, nil
}

func (r OfferRepository) SaveOffers(offers []full.Offer) ([]full.Offer, error) {
	offerRows := r.fullOfferToRow(offers)

	err := r.saveOffers(offerRows)
	if err != nil {
		return nil, err
	}

	return offers, err
}

// fullOfferToRow maps the domain entity data to the mysql compatible data
func (r OfferRepository) fullOfferToRow(offers []full.Offer) (offerRows []OfferRow) {
	for _, offer := range offers {
		newOfferRow := OfferRow{}

		newOfferRow.Id = offer.Id
		newOfferRow.SupplierId = offer.SupplierId
		newOfferRow.Price = offer.Price
		newOfferRow.CityName = offer.City.Name()

		offerRows = append(offerRows, newOfferRow)
	}

	return offerRows
}

func (r OfferRepository) saveOffers(offers []OfferRow) error {
	return nil
}
