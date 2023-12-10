package patterns

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_Size_ReturnsSizeOfCache(t *testing.T) {
	cache := NewInMemoryCache[string, string](1*time.Second, 1*time.Second)

	cache.Put("foo-1", "bar")
	cache.Put("foo-2", "bar")

	require.Equal(t, cache.Size(), 2)
}

func Test_deleteExpiredItems_RemovesExpiredItems(t *testing.T) {
	cache := NewInMemoryCache[string, string](1*time.Second, 1*time.Second)

	cache.Put("foo-1", "bar")
	cache.Put("foo-2", "bar")
	time.Sleep(3 * time.Second)

	require.Zero(t, cache.Size())
}

func Test_Contains_ReturnsFalseWhenKeyIsNotPresent(t *testing.T) {
	cache := NewInMemoryCache[string, string](5*time.Minute, 15*time.Second)

	contains := cache.Contains("foo")

	require.False(t, contains)
}

func Test_Contains_ReturnsTrueWhenKeyIsPresent(t *testing.T) {
	cache := NewInMemoryCache[string, string](5*time.Minute, 15*time.Second)

	cache.Put("foo", "bar")
	contains := cache.Contains("foo")

	require.True(t, contains)
}
