package getoffer

import (
	"ddd-implementation/pkg/offersearch/domain/entity/partial"
	"testing"
	"time"
)

// Interfaces enables us to easily mock dependencies
type daClientMock struct{}

func (m daClientMock) GetOffers(hotelId uint, from time.Time, to time.Time) ([]partial.Offer, error) {
	return []partial.Offer{
		{Id: 1, HotelId: 42, HotelName: "Test Hotel", City: "Test City", CountryCode: "TT", Price: 100},
	}, nil
}

func TestImpl_GetOffersFromSupplier(t *testing.T) {
	service := Impl{
		daClient: daClientMock{},
	}

	offers, err := service.GetOffersFromSupplier(da, 1, time.Now(), time.Now())
	if err != nil {
		t.Errorf("GetOffersFromSupplier failed: %v", err)

		return
	}

	if len(offers) != 1 {
		t.Errorf("wrong offer count: %d", len(offers))
	}

	if offers[0].Id != 1 {
		t.Errorf("GetOffersFromSupplier failed: wrong offer")
	}
}
