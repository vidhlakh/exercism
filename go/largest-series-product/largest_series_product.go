//Package lsproduct provides functionality for finding largest product of substring in a given string
package lsproduct

import (
	"errors"
	"strconv"
)

// LargestSeriesProduct return the largest series product from digits of length span
func LargestSeriesProduct(digits string, span int) (int, error) {
	if span < 0 {
		return 0, errors.New("span must be positive")
	}
	if len(digits) < span {
		return 0, errors.New("span must be at least the length of digits")
	}

	max := 0
	digitsAsSlice := []rune(digits)

	for i := 0; i <= len(digitsAsSlice)-span; i++ {
		n := 1
		for j := i; j < i+span; j++ {
			m, err := strconv.Atoi(string(digitsAsSlice[j]))
			if err != nil {
				return 0, err
			}
			n *= m
		}
		if n > max {
			max = n
		}
	}

	return max, nil
}
