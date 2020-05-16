package diffsquares

import "math"

//SquareOfSum returns the square of the sum of the first ten natural numbers
func SquareOfSum(n int) int {
	var res int
	for i := 1; i <= n; i++ {
		res += i
	}
	return int(math.Pow(float64(res), 2.0))
}

//SumOfSquares return the sum of the squares of the first ten natural numbers
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

//Difference returns the difference of SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
