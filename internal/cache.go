package internal

import "time"

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	ttl   time.Duration
}

func (c *Cache) Add(key string, val []byte) {

}

func (c *Cache) Get(key string) (val []byte, ok bool) {
	return []byte{}, false
}

func NewCache(ttl time.Duration) Cache {
	return Cache{ttl: ttl}
}
