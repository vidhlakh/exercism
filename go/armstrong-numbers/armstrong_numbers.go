// Package armstrong has IsNumber function
package armstrong

import (
	"math"
)

//IsNumber tells whether it is armstrong number
func IsNumber(input int) bool {
	actual := input
	count := float64(0)
	var digit []float64
	for input >= 1 {
		digit = append(digit, float64(input%10))
		input = input / 10
		count++
	}

	value := 0
	var d float64
	for _, d = range digit {
		value += int(math.Pow(d, count))

	}
	if value == actual || count == 1 {
		return true
	}

	return false
}
