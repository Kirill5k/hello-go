package patterns

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_IntToRoman_ConvertsIntegersToRomanNumerals(t *testing.T) {
	require.Equal(t, "I", IntToRoman(1))
	require.Equal(t, "II", IntToRoman(2))
	require.Equal(t, "III", IntToRoman(3))
	require.Equal(t, "IV", IntToRoman(4))
	require.Equal(t, "V", IntToRoman(5))
	require.Equal(t, "VI", IntToRoman(6))
	require.Equal(t, "VII", IntToRoman(7))
	require.Equal(t, "VIII", IntToRoman(8))
	require.Equal(t, "X", IntToRoman(10))
	require.Equal(t, "XXIV", IntToRoman(24))
	require.Equal(t, "XL", IntToRoman(40))
	require.Equal(t, "IX", IntToRoman(9))
	require.Equal(t, "XLIX", IntToRoman(49))
}
