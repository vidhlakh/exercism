package reverse

//Reverse accepts a string and returns a reversed string
func Reverse(word string) string {
	result := "" //variable to hold the reversed string
	for _, letter := range word {
		result = string(letter) + result // Parse letter(which is rune ) to string and store the letter next to the previous letter
	}
	return result
}
