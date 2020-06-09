//Package pangram finds whether given string is pangram or not
package pangram

import "strings"

//IsPangram returns true if the input has all letters int the alphabet. Case insensitive
func IsPangram(input string) (result bool) {
	result = false
	inputString := strings.ToLower(input)
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alpha := make(map[string]int)
	for _, a := range alphabet {
		alpha[string(a)]++
	}
	for _, letter := range inputString {
		if alpha[string(letter)] > 1 {
			alpha[string(letter)]--
		} else {
			delete(alpha, string(letter))
		}
	}
	if len(alpha) == 0 {
		result = true
	}
	return
}

//Another solution//IsPangram returns true if input is recognized as pangram
// func IsPangram(input string) bool {
// 	if len(input) == 0 {
// 		return false
// 	}
// 	str := strings.ToLower(input)
// 	var chrs = make([]int, 26)
// 	for _, r := range str {
// 		if r >= 'a' && r <= 'z' {
// 			chrs[r-'a']++
// 		}
// 	}
// 	for _, r := range chrs {
// 		if r == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
