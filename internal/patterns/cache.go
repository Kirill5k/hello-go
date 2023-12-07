package patterns

type Cache[K any, V any] interface {
	Get(key K) *V
	Put(key K, value V)
	Contains(key K) bool
}
