package cache

import "time"

type Cacher[T comparable, V any] map[T]*Item[V]

func NewCacher[T comparable, V any]() Cacher[T, V] {
	return make(Cacher[T, V])
}

func (c Cacher[T, V]) Set(key T, value V, ttl time.Duration) {
	c[key] = &Item[V]{
		ttl:   time.Now().Add(ttl),
		value: value,
	}
}

func (c Cacher[T, V]) Get(key T) *Item[V] {
	if item, found := c[key]; found {
		if item.isExpired() {
			delete(c, key)

			return nil
		}

		return item
	}

	return nil
}

type Item[T any] struct {
	ttl   time.Time
	value T
}

func (i Item[T]) isExpired() bool {
	return time.Now().After(i.ttl)
}

func (i Item[T]) Value() T {
	return i.Value()
}
