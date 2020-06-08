//Package wordcount has function to count the words in the given input
package wordcount

import (
	"regexp"
	"strings"
)

//Frequency map to hold the word and the count
type Frequency map[string]int

//WordCount counts the number of words. Word can have numbers, letters, apostrophe (no other punctuations )
//Remove '' also
func WordCount(input string) Frequency {
	count := make(map[string]int)
	inputString := strings.ToLower(input)
	// Make a Regex to say we only want letters and numbers
	reg := regexp.MustCompile("[a-zA-Z0-9']+")

	words := reg.FindAllString(inputString, -1)

	for _, word := range words {
		word := strings.Trim(word, "'")
		count[word]++
	}

	return count

}
