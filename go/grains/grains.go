//Package grains calculates the number of grains of wheat on a chessboard given that the number on each square doubles.
package grains

import (
	"fmt"
)

//Square calculates the number of wheat grains on a given square.
func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return uint64(1), fmt.Errorf("Choose a positive number")
	}
	return 1 << (number - 1), nil
}

//Total calculates the total number of grains on the chessboard.
func Total() uint64 {
	return 1<<64 - 1
}
