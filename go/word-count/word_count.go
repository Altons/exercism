package wordcount

import (
	"regexp"
	"strings"
)

//Frequency is a structure to hold num of occur of given words
type Frequency map[string]int

//WordCount returns the count the occurrences of each _word_ in that phrase
func WordCount(s string) Frequency {
	s = strings.ToLower(s)
	regex, _ := regexp.Compile(`\n|'\B|[^a-zA-Z0-9'? ]|\s'`)
	lst := strings.Split(regex.ReplaceAllString(s, " "), " ")
	res := Frequency{}

	for _, c := range lst {
		if c == "" {
			continue
		}
		newS := string(c)
		if _, ok := res[newS]; ok {
			res[newS]++
		} else {
			res[newS] = 1
		}
	}
	return res
}
