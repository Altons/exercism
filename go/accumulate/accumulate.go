package accumulate

//Accumulate will apply a function
func Accumulate(s []string, f func(string) string) []string {
	for i, value := range s {
		s[i] = f(value)
	}
	return s
}
