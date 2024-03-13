package internal

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	ttl   time.Duration
	mu    *sync.RWMutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{val: val}
}

func (c *Cache) Get(key string) (val []byte, ok bool) {
	c.mu.RLock()
	c.mu.RUnlock()
	entry, ok := c.cache[key]
	if !ok {
		return []byte{}, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.ttl)
	for range ticker.C {
		fmt.Println("TICK")
	}
}

func NewCache(ttl time.Duration) Cache {
	c := Cache{
		ttl:   ttl,
		cache: map[string]cacheEntry{},
		mu:    &sync.RWMutex{},
	}

	go c.reapLoop()

	return c
}
