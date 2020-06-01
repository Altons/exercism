package luhn

import (
	"strconv"
	"strings"
)

//Valid check whether or not it is valid per the Luhn formula
func Valid(s string) bool {
	newS := strings.ReplaceAll(s, " ", "")
	if len(newS) < 2 {
		return false
	}
	double := len(newS)%2 == 0
	var sum int
	for _, d := range newS {
		digit, err := strconv.Atoi(string(d))
		if err != nil {
			return false
		}
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
