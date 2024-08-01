package direktanbindung

type DaOfferResponse struct {
	Data []DaOffer `json:"data"`
}

type DaOffer struct {
	Id        uint   `json:"id"`
	HotelId   uint   `json:"hotelId"`
	HotelName string `json:"hotelName"`
	Price     uint   `json:"price"`
	Country   string `json:"country"`
	City      string `json:"city"`
}
