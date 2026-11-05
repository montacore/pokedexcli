package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	data		map[string]CacheEntry
	mut			sync.Mutex
}

type CacheEntry struct {
	createdAt		time.Time
	val				[]byte
}

func NewCache(interval time.Duration) *Cache {
	cache:= &Cache{
		data: make(map[string]CacheEntry),
	}
	go cache.reapLoop(interval)
	
	return cache
}

func (c *Cache) Add(key string, val []byte){

	c.mut.Lock()
	defer c.mut.Unlock()
	cachedEntry := CacheEntry{}
	cachedEntry.createdAt = time.Now()
	cachedEntry.val = val
	c.data[key] = cachedEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.Lock()
	defer c.mut.Unlock()
	entry, ok := c.data[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

//NOTE: reapLoop runs a ticker to remove expired keys periodically
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mut.Lock()
		for k, entry := range c.data {
			age := time.Now().Sub(entry.createdAt)
			if age >= interval{
				delete(c.data, k)
			}
		}
	c.mut.Unlock()	
	}

}

