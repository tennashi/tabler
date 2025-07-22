package metadata

import "sync"

type Cache struct {
	mu    sync.RWMutex
	items map[string]*ExtractedMetadata
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]*ExtractedMetadata),
	}
}

func (c *Cache) Get(key string) (*ExtractedMetadata, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	return item, found
}

func (c *Cache) Set(key string, value *ExtractedMetadata) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = value
}
