package patterns

import (
	"sync"
	"time"
)

type Cache[K any, V any] interface {
	Get(key K) *V
	Put(key K, value V)
	Contains(key K) bool
}

type cacheEntry[V any] struct {
	value V
	time  time.Time
}

type inMemoryCache[K any, V any] struct {
	mu     sync.Mutex
	values map[string]cacheEntry[V]
}
