package isogram

import (
	"strings"
)

//IsIsogram determines whether a word or phrase is an isogram
func IsIsogram(s string) bool {
	replacer := strings.NewReplacer(" ", "", "-", "")
	s = strings.ToLower(replacer.Replace(s))
	lst := strings.Split(s, "")
	repeated := make(map[string]bool)
	for _, c := range lst {
		if _, ok := repeated[c]; ok {
			return false
		}
		repeated[c] = true
	}
	return true
}
