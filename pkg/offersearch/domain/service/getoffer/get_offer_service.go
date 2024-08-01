package getoffer

import (
	"ddd-implementation/pkg/offersearch/domain/entity/full"
	"ddd-implementation/pkg/offersearch/domain/entity/partial"
	"ddd-implementation/pkg/offersearch/domain/repository"
	"ddd-implementation/pkg/offersearch/domain/value"
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
}

type Impl struct {
	daClient repository.SupplierOffer
}

func NewService(
	daClient repository.SupplierOffer,
) Impl {
	return Impl{
		daClient: daClient,
	}
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
		newFullOffer.City = value.NewCity(partialOffer.City, partialOffer.CountryCode)

		newFullOffer.SupplierId = supplierId

		fullOffers = append(fullOffers, newFullOffer)
	}

	return fullOffers
}
