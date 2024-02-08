package romans

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_IntToRoman_ConvertsIntegersToRomanNumerals(t *testing.T) {
	require.Equal(t, "I", FromInt(1))
	require.Equal(t, "II", FromInt(2))
	require.Equal(t, "III", FromInt(3))
	require.Equal(t, "IV", FromInt(4))
	require.Equal(t, "V", FromInt(5))
	require.Equal(t, "VI", FromInt(6))
	require.Equal(t, "VII", FromInt(7))
	require.Equal(t, "VIII", FromInt(8))
	require.Equal(t, "X", FromInt(10))
	require.Equal(t, "XXIV", FromInt(24))
	require.Equal(t, "XL", FromInt(40))
	require.Equal(t, "IX", FromInt(9))
	require.Equal(t, "XLIX", FromInt(49))
	require.Equal(t, "XIX", FromInt(19))
}
