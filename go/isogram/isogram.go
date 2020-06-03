package isogram

import "strings"

//IsIsogram checks whether word is isogram or not
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
