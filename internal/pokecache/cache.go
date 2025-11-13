package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

type Cache struct {
	mu			*sync.RWMutex
	entries		map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		mu:			&sync.RWMutex{},
		entries:	make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.entries[key] = cacheEntry{
		createdAt:	time.Now(),
		val:		val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	timePassed := time.Now().Add(-interval)
	for key, entry := range c.entries {
		if entry.createdAt.Before(timePassed) {
			delete(c.entries, key)
		}
	}
}