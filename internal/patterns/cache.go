package patterns

import (
	"sync"
	"time"
)

type Cache[K comparable, V any] interface {
	Get(key K) *V
	Put(key K, value V)
	Contains(key K) bool
}

type cacheEntry[V any] struct {
	value V
	time  time.Time
}

type inMemoryCache[K comparable, V any] struct {
	mu     sync.Mutex
	values map[K]cacheEntry[V]
}

func (c *inMemoryCache[K, V]) Get(key K) *V {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.values[key]
	if !ok {
		return nil
	}
	return &v.value
}

func (c *inMemoryCache[K, V]) Contains(key K) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.values[key]
	return ok
}

func (c *inMemoryCache[K, V]) Put(key K, value V) {
	c.mu.Lock()
	c.values[key] = cacheEntry[V]{value: value, time: time.Now()}
	c.mu.Unlock()
}
