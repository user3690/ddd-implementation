package direktanbindung

import (
	"ddd-implementation/pkg/offersearch/domain/entity/partial"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Client struct {
	baseUrl string
}

func NewClient(baseUrl string) *Client {
	return &Client{baseUrl: baseUrl}
}

// GetOffers send request to api, parse response, map to domain entity
func (c Client) GetOffers(hotelId uint, from time.Time, to time.Time) ([]partial.Offer, error) {
	request, err := c.createOfferRequest(hotelId, from, to)
	if err != nil {
		return nil, err
	}

	response, err := c.send(request)
	if err != nil {
		return nil, err
	}

	parsedResponse, err := c.parseResponse(response)
	if err != nil {
		return nil, err
	}

	return c.mapResponse(parsedResponse), nil
}

func (c Client) send(request *http.Request) ([]byte, error) {
	path := filepath.Join("pkg", "offersearch", "infrastructure", "web", "direktanbindung")
	path = path + string(os.PathSeparator) + "response.json"
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

func (c Client) parseResponse(responseContent []byte) (DaOfferResponse, error) {
	var response DaOfferResponse

	err := json.Unmarshal(responseContent, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (c Client) mapResponse(daResponse DaOfferResponse) []partial.Offer {
	var partialOffers []partial.Offer

	for _, daOffer := range daResponse.Data {
		partialOffer := partial.Offer{
			Id:          daOffer.Id,
			HotelId:     daOffer.HotelId,
			HotelName:   daOffer.HotelName,
			City:        daOffer.City,
			CountryCode: daOffer.Country,
			Price:       daOffer.Price,
		}

		partialOffers = append(partialOffers, partialOffer)
	}

	return partialOffers
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
