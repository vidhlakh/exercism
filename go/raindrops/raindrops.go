package raindrops

import "strconv"

//Convert gets a number and return a string depending on the number
func Convert(num int) string {
	result := ""
	// if number is factor of 3 append Pling
	if num%3 == 0 {
		result = "Pling"
	}
	//if number is factor of 5 append Plang
	if num%5 == 0 {
		result = result + "Plang"
	}
	//if number is factor of 7 append Plong
	if num%7 == 0 {
		result = result + "Plong"
	}
	//if number if not a factor of 3,5,7 then return number
	if num%3 != 0 && num%5 != 0 && num%7 != 0 {
		result = strconv.Itoa(num)
	}

	return result
}
