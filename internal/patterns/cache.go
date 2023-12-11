package patterns

import (
	"sync"
	"time"
)

type Cache[K comparable, V any] interface {
	Get(key K) *V
	Put(key K, value V)
	Contains(key K) bool
	Size() int
}

type cacheEntry[V any] struct {
	value V
	time  time.Time
}

type InMemoryCache[K comparable, V any] struct {
	mu     sync.Mutex
	values map[K]cacheEntry[V]
}

func NewInMemoryCache[K comparable, V any](expiresIn time.Duration, checkEvery time.Duration) *InMemoryCache[K, V] {
	cache := &InMemoryCache[K, V]{values: make(map[K]cacheEntry[V])}
	go deleteExpiredItems(cache, expiresIn, checkEvery)
	return cache
}

func deleteExpiredItems[K comparable, V any](c *InMemoryCache[K, V], expiresIn time.Duration, checkEvery time.Duration) {
	c.deleteExpiredItems(time.Now(), expiresIn)
	time.Sleep(checkEvery)
	deleteExpiredItems(c, expiresIn, checkEvery)
}

func (c *InMemoryCache[K, V]) Get(key K) *V {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.values[key]
	if !ok {
		return nil
	}
	return &v.value
}

func (c *InMemoryCache[K, V]) Contains(key K) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.values[key]
	return ok
}

func (c *InMemoryCache[K, V]) Put(key K, value V) {
	c.mu.Lock()
	c.values[key] = cacheEntry[V]{value: value, time: time.Now()}
	c.mu.Unlock()
}

func (c *InMemoryCache[K, V]) Size() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.values)
}

func (c *InMemoryCache[K, V]) deleteExpiredItems(currentTime time.Time, expiresIn time.Duration) {
	c.mu.Lock()
	for k, v := range c.values {
		if v.time.Add(expiresIn).After(currentTime) {
			delete(c.values, k)
		}
	}
	c.mu.Unlock()
}
