package twofer

import "fmt"

// ShareWith - Be kind and always share.
func ShareWith(name string) string {
	subj := ""
	switch {
	case name != "":
		subj = name
	default:
		subj = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", subj)
}
