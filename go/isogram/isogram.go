package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram determines whether a word or phrase is an isogram
func IsIsogram(s string) bool {
	replacer := strings.NewReplacer(" ", "", "-", "")
	s = replacer.Replace(s)
	for _, c := range s {
		c = unicode.ToLower(c)
		if strings.Count(strings.ToLower(s), string(c)) > 1 {
			return false
		}
	}
	return true
}
