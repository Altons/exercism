package grains

import (
	"fmt"
	"math"
)

//Square ...
func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return uint64(1), fmt.Errorf("Choose a positive number")
	}
	return uint64(math.Pow(2, float64(number)-1)), nil
}

//Total ....
func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		total += uint64(math.Pow(2, float64(i)-1))
	}
	return total
}
