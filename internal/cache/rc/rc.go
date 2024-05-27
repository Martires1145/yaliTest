package rc

import (
	"cmdTest/internal/database"
	"context"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	cache     *redis.Client
	OnEvicted func(key string, value string)
}

func New(onEvicted func(key string, value string)) *Cache {
	return &Cache{
		cache:     database.NewRedis(),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value string, ok bool) {
	value, err := c.cache.Get(context.Background(), key).Result()
	if err == nil {
		return value, true
	}
	return
}

func (c *Cache) Add(key string, value string) {
	v, err := c.cache.Get(context.Background(), key).Result()
	if err != nil || v == "" {
		c.cache.Set(context.Background(), key, value, -1)
	}
}
