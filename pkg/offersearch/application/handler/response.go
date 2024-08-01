package handler

import (
	"ddd-implementation/pkg/offersearch/domain/entity/full"
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonOfferResponse struct {
	Status int         `json:"status"`
	Count  int         `json:"count"`
	Data   []JsonOffer `json:"data"`
}

type JsonOffer struct {
	Id        uint   `json:"id"`
	Supplier  string `json:"supplier"`
	HotelId   uint   `json:"hotelId"`
	HotelName string `json:"hotelName"`
	Price     uint   `json:"price"`
}

func writeJsonOfferResponse(
	writer http.ResponseWriter,
	status int,
	offers []full.Offer,
	supplier string,
) {
	jsonOffers := make([]JsonOffer, 0, len(offers))
	for _, offer := range offers {
		jsonOffers = append(jsonOffers, JsonOffer{
			Id:        offer.Id,
			Supplier:  supplier,
			HotelId:   offer.HotelId,
			HotelName: offer.HotelName,
			Price:     offer.Price,
		})
	}

	jsonResponse := JsonOfferResponse{
		Status: status,
		Count:  len(jsonOffers),
		Data:   jsonOffers,
	}

	data, err := json.Marshal(jsonResponse)
	if err != nil {
		fmt.Println(err)

		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)

	_, err = writer.Write(data)
	if err != nil {
		fmt.Println(err)
	}

	return
}

type JsonErrorResponse struct {
	Errors []JsonError `json:"errors"`
}

type JsonError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func writeError(
	writer http.ResponseWriter,
	status int,
	title string,
	responseErr error,
) {
	newJsonError := JsonErrorResponse{
		Errors: []JsonError{
			{
				Status: status,
				Title:  title,
				Detail: responseErr.Error(),
			},
		},
	}

	data, err := json.Marshal(newJsonError)
	if err != nil {
		fmt.Println(err)

		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)

	_, err = writer.Write(data)
	if err != nil {
		fmt.Println(err)
	}

	return
}
