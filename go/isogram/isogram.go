package isogram

import (
	"strings"
)

//IsIsogram determines whether a word or phrase is an isogram
func IsIsogram(s string) bool {
	repeated := make(map[rune]bool)
	for _, c := range strings.ToLower(s) {
		if c == '-' || c == ' ' {
			continue
		}
		if repeated[c] {

			return false
		}
		repeated[c] = true
	}
	return true
}
