package scrabble

import (
	"unicode"
)

//Score computes the Scrabble score for a given word.
func Score(s string) int {
	var score int
	for _, v := range s {
		v = unicode.ToUpper(v)
		switch string(v) {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			score++
		case "D", "G":
			score += 2
		case "B", "C", "M", "P":
			score += 3
		case "F", "H", "V", "W", "Y":
			score += 4
		case "J", "X":
			score += 8
		case "Q", "Z":
			score += 10
		default: // K
			score += 5
		}
	}

	return score

}
