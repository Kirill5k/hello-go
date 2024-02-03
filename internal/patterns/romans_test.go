package patterns

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ToRoman_ConvertsIntToRomanNumber(t *testing.T) {
	require.Equal(t, IntToRoman(1), "I")
	require.Equal(t, IntToRoman(3), "III")
	require.Equal(t, IntToRoman(4), "IV")
	require.Equal(t, IntToRoman(5), "V")
	require.Equal(t, IntToRoman(6), "VI")
	require.Equal(t, IntToRoman(7), "VI7")
	require.Equal(t, IntToRoman(10), "X")
}
