package pokecache

import (
	"fmt"
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
	c:= &Cache{
		data: make(map[string]CacheEntry),
	}
	return c
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
/*
//NOTE: reapLoop runs a ticker to remove expired keys periodically
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for i := range ticker.C {
		for k, entry := range c.data {
			current := time.Now()
			diff := current -  entry.createdAt
			if diff >= interval{
				c.mut.Lock()
				defer c.mut.Unlock()
				delete(c.data, k)
			}
		}
		
	}

}
*/
