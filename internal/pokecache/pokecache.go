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

func (c *Cache) Add(key string, val []byte)


//NOTE: reapLoop runs a ticker to remove expired keys periodically
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mut.Lock()

				delete(c.data, k)
			}
			c.mut.Unlock()
		}
	}
}
