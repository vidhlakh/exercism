package main

import (
	"fmt"
	"math"
)

//IsNumber tells whether it is armstrong number
func IsNumber(input int) bool {
	actual := input
	count := float64(0)
	var digit []float64
	for input >= 1 {
		fmt.Println("rem ", input%10)
		digit = append(digit, float64(input%10))
		input = input / 10
		count++
	}
	fmt.Println("digit:", digit, count)
	value := 0
	var d float64
	for _, d = range digit {
		fmt.Println("d:", d, "math", math.Pow(d, count))
		value += int(math.Pow(d, count))
		fmt.Println("Value:", value)
	}
	if value == actual || count == 1 {
		return true
	}

	return false
}
func main() {
	fmt.Println("Is Armstrong number:", IsNumber(153))
}
