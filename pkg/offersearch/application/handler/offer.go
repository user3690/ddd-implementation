package handler

import (
	"awesomeProject/pkg/offersearch/domain/service/manipulateOffer"
	"net/http"
	"strconv"
	"time"
)

type Offer struct {
	manipulateOfferService manipulateOffer.Service
}

func NewOfferHandler(
	manipulateOfferService manipulateOffer.Service,
) Offer {
	return Offer{
		manipulateOfferService: manipulateOfferService,
	}
}

// Handle validates und coordinates requests and responds accordingly
func (h Offer) Handle(response http.ResponseWriter, request *http.Request) {
	supplierParam := request.URL.Query().Get("supplier")
	hotelIdParam := request.URL.Query().Get("hotelid")
	fromParam := request.URL.Query().Get("from")
	toParam := request.URL.Query().Get("to")

	hotelId, err := strconv.ParseUint(hotelIdParam, 10, 64)
	if err != nil {
		writeError(
			response,
			http.StatusBadRequest,
			"invalid hotelId param",
			err,
		)

		return
	}

	from, err := time.Parse("2006-01-02", fromParam)
	if err != nil {
		writeError(
			response,
			http.StatusBadRequest,
			"invalid from param",
			err,
		)

		return
	}

	to, err := time.Parse("2006-01-02", toParam)
	if err != nil {
		writeError(
			response,
			http.StatusBadRequest,
			"invalid to param",
			err,
		)

		return
	}

	fullOffers, err := h.manipulateOfferService.GetOffersFromSupplier(supplierParam, uint(hotelId), from, to)
	if err != nil {
		writeError(
			response,
			http.StatusInternalServerError,
			"error getting offers",
			err,
		)

		return
	}

	h.manipulateOfferService.SaveOffers(fullOffers)
}
