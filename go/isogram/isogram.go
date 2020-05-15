package isogram

import (
	"strings"
)

//IsIsogram determines whether a word or phrase is an isogram
func IsIsogram(s string) bool {
	lst := strings.Split(strings.ToLower(s), "")
	repeated := make(map[string]bool)
	for _, c := range lst {
		if c == "-" || c == " " {
			continue
		}
		if repeated[c] {

			return false
		}
		repeated[c] = true
	}
	return true
}
