package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram determines whether a word or phrase is an isogram
func IsIsogram(s string) bool {
	dict := make(map[rune]int)
	replacer := strings.NewReplacer(" ", "", "-", "")
	s = replacer.Replace(s)
	for _, c := range s {
		c = unicode.ToLower(c)
		if _, ok := dict[c]; ok {
			dict[c]++
		} else {
			dict[c] = 1
		}
		if dict[c] > 1 {
			return false
		}
	}
	return true
}
