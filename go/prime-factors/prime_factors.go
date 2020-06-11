//Package prime has factors functions
package prime

// Factors returns the prime factors fo the number
func Factors(input int64) (factors []int64) {
	factors = []int64{}
	for i := int64(2); input > 1; i++ {
		for ; input%i == 0; input = input / i {
			factors = append(factors, i)

		}
	}
	return
}

// Prime gives the list of prime numbers
// func Prime(n int) (prime []int) {
// 	for i := 2; i < n+1; i++ {
// 		for j := 2; j < n; j++ {
// 			if i%j != 0 {
// 				prime = append(prime, i)
// 			}
// 		}
// 	}
// 	return
// }
