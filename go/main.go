package main

import (
	"errors"
	"fmt"
	"strings"
)

//Hamming
func Distance(a, b string) (int, error) {
	count := 0 //Int Variable to Count the Difference

	if len(a) == len(b) {

		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				count++ //Increment when a string is different from b string
			}

		}
		return count, nil
	}
	return 0, errors.New("strings of unequal length")

}

// Isogram
func IsIsogram(word string) bool {
	uniqueLetters := map[rune]bool{}
	/*Writing in multiple lines
	wordLow := strings.ToLower(word)                    // word in lower case
	removehyphen := strings.Replace(wordLow, "-", " ", -1) //replace hyphen and space as duplicaions are allowed for hyphen and space
	finalWord := strings.Replace(removehyphen," ","",-1)  // remove space
	*/
	//remove hyphen with space , remove space with "" , lower the case to get the finalWord
	finalWord := strings.ToLower(strings.TrimSpace(strings.Replace(strings.Replace(word, "-", " ", -1), " ", "", -1)))

	for _, letter := range finalWord {

		if uniqueLetters[letter] == true {
			return false //return false as duplicates are found
		} else {
			uniqueLetters[letter] = true

		}
	}
	return true // if all letters are unique
}
func main() {
	st, err := Distance("AAGT", "AAGDB")
	if err != nil {
		fmt.Println("error ", err)
	}
	fmt.Println("returned content ", st)

	//call IsIsogram function
	word := "six-year-old"
	result := IsIsogram(word)

	fmt.Printf("Word %v is %v", word, result)
}
