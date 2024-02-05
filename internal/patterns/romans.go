package patterns

import "strings"

type letterToNumber struct {
	letter string
	number int
}

var lettersToNumbers = []letterToNumber{
	{"X", 10},
	{"V", 5},
	{"I", 1},
}

func IntToRoman(n int) string {
	remaining := n
	result := ""
	for _, ltn := range lettersToNumbers {
		if ltn.number <= remaining {
			count := remaining / ltn.number
			remaining = remaining % ltn.number
			result = result + strings.Repeat(ltn.letter, count)
		}
	}
	return result
}
