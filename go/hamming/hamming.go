//Package hamming implements Hamming's distance
package hamming

import "errors"

//Distance calculate Hamming distance of 2 DNA strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Err")
	}
	var sum int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			sum++
		}
	}

	return sum, nil
}
