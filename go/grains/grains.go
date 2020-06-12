package grains

import (
	"fmt"
)

//Square ...
func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return uint64(1), fmt.Errorf("Choose a positive number")
	}
	return Pow(2, uint64(number)-1), nil
}

//Total ....
func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		total += Pow(2, uint64(i)-1)
	}
	return total
}

//Pow calculates compute x^n using binary powering algorithm
func Pow(a, b uint64) uint64 {
	var p uint64
	p = 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
