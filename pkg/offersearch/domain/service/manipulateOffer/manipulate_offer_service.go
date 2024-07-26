package manipulateOffer

import (
	"awesomeProject/pkg/offersearch/domain/entity/full"
	"awesomeProject/pkg/offersearch/domain/entity/partial"
	"awesomeProject/pkg/offersearch/domain/repository"
	"fmt"
	"time"
)

const (
	da = "da"
)

var supplierMap = map[string]uint{
	da: 1,
}

type Service interface {
	GetOffersFromSupplier(
		supplierId string,
		hotelId uint,
		from time.Time,
		to time.Time,
	) ([]full.Offer, error)
	SaveOffers(offers []full.Offer) (int, error)
}

type Impl struct {
	odsOfferRepo repository.Offer
	daClient     repository.SupplierOffer
}

func NewService(
	odsOfferRepo repository.Offer,
	daClient repository.SupplierOffer,
) Impl {
	return Impl{
		odsOfferRepo: odsOfferRepo,
		daClient:     daClient,
	}
}

func (i Impl) SaveOffers(offers []full.Offer) (int, error) {
	return 0, nil
}

func (i Impl) GetOffersFromSupplier(
	supplierId string,
	hotelId uint,
	from time.Time,
	to time.Time,
) ([]full.Offer, error) {
	switch supplierId {
	case da:
		return i.getOffersFromDirektAnbindung(hotelId, from, to)
	default:
		return nil, fmt.Errorf("supplier %s not implemented", supplierId)
	}
}

func (i Impl) getOffersFromDirektAnbindung(
	hotelId uint,
	from time.Time,
	to time.Time,
) ([]full.Offer, error) {
	partialOffers, err := i.daClient.GetOffers(hotelId, from, to)
	if err != nil {
		return nil, err
	}

	fullOffers := i.partialToFullOffer(partialOffers, supplierMap[da])

	return fullOffers, nil
}

func (i Impl) partialToFullOffer(partialOffers []partial.Offer, supplierId uint) []full.Offer {
	fullOffers := make([]full.Offer, 0, len(partialOffers))

	for _, partialOffer := range partialOffers {
		newFullOffer := full.Offer{}

		newFullOffer.Id = partialOffer.Id
		newFullOffer.HotelId = partialOffer.HotelId
		newFullOffer.HotelName = partialOffer.HotelName
		newFullOffer.Price = partialOffer.Price
		newFullOffer.City = partialOffer.City

		newFullOffer.SupplierId = supplierId

		fullOffers = append(fullOffers, newFullOffer)
	}

	return fullOffers
}
