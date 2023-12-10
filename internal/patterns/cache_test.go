package patterns

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

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
