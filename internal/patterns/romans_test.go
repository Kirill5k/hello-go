package patterns

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ToRoman_ConvertsIntToRomanNumber(t *testing.T) {
	require.Equal(t, "I", IntToRoman(1))
	require.Equal(t, "III", IntToRoman(3))
	require.Equal(t, "V", IntToRoman(5))
	require.Equal(t, "VI", IntToRoman(6))
	require.Equal(t, "VII", IntToRoman(7))
	require.Equal(t, "X", IntToRoman(10))
	require.Equal(t, "IV", IntToRoman(4))
}
