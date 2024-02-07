package patterns

import "strings"

type letterToNumber struct {
	letter string
	number int
}

var lettersToNumbers = []letterToNumber{
	{"L", 50},
	{"X", 10},
	{"V", 5},
	{"I", 1},
}

func IntToRoman(n int) string {
	remaining := n
	result := ""
	for i, ltn := range lettersToNumbers {
		if ltn.number <= remaining {
			count := remaining / ltn.number
			remaining = remaining % ltn.number
			if count <= 3 {
				result = result + strings.Repeat(ltn.letter, count)
			} else if i > 0 && count == 4 {
				prevLtn := lettersToNumbers[i-1]
				result = result + ltn.letter + prevLtn.letter
			}
		}
	}
	return result
}
