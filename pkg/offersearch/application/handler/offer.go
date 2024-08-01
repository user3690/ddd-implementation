package handler

import (
	"ddd-implementation/pkg/offersearch/domain/service/getoffer"
	"ddd-implementation/pkg/offersearch/domain/service/saveoffer"
	"net/http"
	"strconv"
	"time"
)

type Offer struct {
	getOfferService  getoffer.Service
	saveOfferService saveoffer.Service
}

func NewOfferHandler(
	getOfferService getoffer.Service,
	saveOfferService saveoffer.Service,
) Offer {
	return Offer{
		getOfferService:  getOfferService,
		saveOfferService: saveOfferService,
	}
}

// Handle validates und coordinates requests, program flow and responds accordingly
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

	fullOffers, err := h.getOfferService.GetOffersFromSupplier(supplierParam, uint(hotelId), from, to)
	if err != nil {
		writeError(
			response,
			http.StatusInternalServerError,
			"error getting offers",
			err,
		)

		return
	}

	offers, err := h.saveOfferService.SaveOffers(fullOffers)
	if err != nil {
		writeError(
			response,
			http.StatusInternalServerError,
			"error saving offers",
			err,
		)

		return
	}

	writeJsonOfferResponse(response, http.StatusOK, offers, supplierParam)
}
