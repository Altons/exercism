package reverse

import (
	"strings"
)

//Reverse take a string and return a reversed version of the string
func Reverse(s string) (r string) {
	lst := strings.Split(s, "")
	var reversed strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		reversed.WriteString(lst[i])
	}
	return reversed.String()
}
