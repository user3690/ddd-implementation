package direktanbindung

import (
	"awesomeProject/pkg/offersearch/domain/entity/partial"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	baseUrl string
}

func NewClient(baseUrl string) *Client {
	return &Client{baseUrl: baseUrl}
}

func (c Client) GetOffers(hotelId uint, from time.Time, to time.Time) ([]partial.Offer, error) {
	request, err := c.createOfferRequest(hotelId, from, to)
	if err != nil {
		return nil, err
	}

	response, err := c.send(request)
	if err != nil {
		return nil, err
	}

	offers := c.parseResponse(response)

	return offers, nil
}

func (c Client) send(request *http.Request) (http.Response, error) {
	return http.Response{}, nil
}

func (c Client) parseResponse(response http.Response) []partial.Offer {
	return []partial.Offer{
		{
			Id:        1,
			HotelId:   42,
			HotelName: "Hotel 42",
			Price:     1234,
			City:      "Amsterdam",
		},
		{
			Id:        2,
			HotelId:   42,
			HotelName: "Hotel 42",
			Price:     2345,
			City:      "Amsterdam",
		},
	}
}

func (c Client) createOfferRequest(hotelId uint, from time.Time, to time.Time) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, c.baseUrl+"/offers", nil)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("hotelid", strconv.FormatUint(uint64(hotelId), 10))
	params.Add("from", from.Format("2006-01-02"))
	params.Add("to", to.Format("2006-01-02"))

	request.URL.RawQuery = params.Encode()

	return request, nil
}
