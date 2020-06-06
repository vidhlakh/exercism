// Package diffsquares has function Square sum, sumofsquares and difference functions
package diffsquares

//SquareOfSum finds the square of the sum of the first N natural numbers
func SquareOfSum(n int) int {
	//Alternate use formular sum of n numbers = n(n+1)/2 & then we take the square of result
	sumValue := 0
	//calculate sum of value from 1 to n
	for i := 1; i <= n; i++ {
		sumValue += i
	}
	//return square of sum of values
	return sumValue * sumValue
}

// SumOfSquares calculates sum of the square of integers from 1 to n
func SumOfSquares(n int) int {
	// sum of n squares = (2n + 1)n(n+1)/6
	sumSquares := 0
	// calculate sum of squares of 1 to n
	for i := 1; i <= n; i++ {

		sumSquares += i * i
	}
	return sumSquares
}

// Difference finds difference between square of the sum and the sum of the squares of the first N natural numbers
func Difference(n int) int {
	//call square of sum and sum of square and return the difference
	squareSum := SquareOfSum(n)
	sumSquares := SumOfSquares(n)

	difference := squareSum - sumSquares
	return difference
}
