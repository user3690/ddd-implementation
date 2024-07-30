package direktanbindung

import (
	"ddd-implementation/pkg/offersearch/domain/entity/partial"
	"ddd-implementation/pkg/offersearch/infrastructure/cache"
	"time"
)

type ClientCache struct {
	client Client
	cacher cache.Cacher[cacheKey, []partial.Offer]
}

type cacheKey struct {
	hotelId uint
	from    time.Time
	to      time.Time
}

func NewClientCache(client Client) ClientCache {
	return ClientCache{
		client: client,
		cacher: make(cache.Cacher[cacheKey, []partial.Offer]),
	}
}

func (c ClientCache) GetOffers(hotelId uint, from time.Time, to time.Time) ([]partial.Offer, error) {
	newKey := cacheKey{hotelId, from, to}

	item := c.cacher.Get(newKey)
	if item != nil {
		return item.Value(), nil
	}

	result, err := c.client.GetOffers(hotelId, from, to)
	if err != nil {
		return nil, err
	}

	c.cacher.Set(newKey, result, time.Minute)

	return result, nil
}
