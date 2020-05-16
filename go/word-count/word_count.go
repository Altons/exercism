package wordcount

import (
	"strings"
)

//Frequency is a structure to hold num of occur of given words
type Frequency map[string]int

//WordCount returns the count the occurrences of each _word_ in that phrase
func WordCount(s string) Frequency {
	replacer := strings.NewReplacer(",", " ", "\n", "", "&@$%^&", "", "!", "", ":", "", ".", "", "'", "")
	lst := strings.Split(replacer.Replace(s), " ")
	res := Frequency{}

	for _, c := range lst {
		if c == "" {
			continue
		}
		newS := strings.ToLower(string(c))
		if _, ok := res[newS]; ok {
			res[newS]++
		} else {
			res[newS] = 1
		}
	}
	return res
}
