//Package luhn implements luhn algorithm
package luhn

import (
	"strings"
	"unicode"
)

// Valid finds the validity of the credit card using luhn algorithm
func Valid(input string) bool {
	cardNumber := strings.ReplaceAll(input, " ", "")
	var result int
	if len(cardNumber) <= 1 {
		return false

	}
	even := len(cardNumber)%2 == 0
	for _, num := range cardNumber {
		//eevery second digit that has to be doubled if length is even
		if !unicode.IsNumber(num) {
			return false
		}
		numInt := int(num - '0')
		if even {
			numInt *= 2
			if numInt > 9 {
				numInt = numInt - 9
			}
		}
		result += numInt
		even = !even
	}
	// card number is valid only if sum of the digits is divisible by 10
	return result%10 == 0
}
