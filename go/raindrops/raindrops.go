//Package raindrops provides a handy method for raindrops
package raindrops

import "strconv"

//Convert create a specific string when some criteria is met
func Convert(x int) string {
	//The rules of raindrops are that if a given number:
	//has 3 as a factor, add 'Pling' to the result.
	//has 5 as a factor, add 'Plang' to the result.
	//has 7 as a factor, add 'Plong' to the result.
	//does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number.

	var r string

	if x%3 == 0 {
		r += "Pling"
	}
	if x%5 == 0 {
		r += "Plang"
	}
	if x%7 == 0 {
		r += "Plong"
	}
	if r == "" {
		r = strconv.Itoa(x)
	}
	return r
}
