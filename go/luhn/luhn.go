package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

//Valid check whether or not it is valid per the Luhn formula
func Valid(s string) bool {
	newS := strings.ReplaceAll(s, " ", "")

	if len(newS) < 2 {
		return false
	}
	nDigits := len(newS)
	// parity := nDigits % 2
	sum := 0
	double := false
	for i := nDigits - 1; i >= 0; i-- {
		char := newS[i]
		if unicode.IsDigit(rune(char)) == false {
			return false
		}
		digit, _ := strconv.Atoi(string(char))
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}

		}
		sum += digit
		double = !double

	}
	return sum%10 == 0
}
