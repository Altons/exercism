package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

//Valid check whether or not it is valid per the Luhn formula
func Valid(s string) bool {
	newS := strings.Replace(s, " ", "", -1)

	if len(newS) < 2 {
		return false
	}
	nDigits := len(newS)
	parity := nDigits % 2
	sum := 0
	for i := nDigits - 1; i >= 0; i-- {
		char := newS[i]
		if unicode.IsDigit(rune(char)) == false {
			return false
		}
		digit, _ := strconv.Atoi(string(char))
		if i%2 == parity {
			digit = digit * 2
		}
		if digit > 9 {
			digit = digit - 9
		}
		sum += digit
	}
	return sum%10 == 0
}
