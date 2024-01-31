package patterns

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_IsLeapYear_ReturnFalseForNonLeapYear(t *testing.T) {
	require.False(t, IsLeapYear(1))
	require.False(t, IsLeapYear(1800))
	require.False(t, IsLeapYear(2001))
	require.False(t, IsLeapYear(2100))
	require.False(t, IsLeapYear(2500))
}

func Test_IsLeapYear_ReturnTrueNonLeapYear(t *testing.T) {
	require.True(t, IsLeapYear(2000))
	require.True(t, IsLeapYear(2004))
	require.True(t, IsLeapYear(2016))
	require.True(t, IsLeapYear(2400))
}
