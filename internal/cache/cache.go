package cache

import (
	"cmdTest/internal/cache/rc"
	"sync"
)

var RedisCache = &Cache{
	mu: sync.Mutex{},
}

type Cache struct {
	mu sync.Mutex
	rc *rc.Cache
}

func (c *Cache) Add(key string, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.rc == nil {
		c.rc = rc.New(nil)
	}
	c.rc.Add(key, value)
}

func (c *Cache) Get(key string) (value string, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.rc == nil {
		return
	}
	if v, ok := c.rc.Get(key); ok {
		return v, ok
	}
	return
}
