//Package anagram tells whether word given has anagram words in the slice of candidates given
package anagram

import (
	"strings"
	"unicode"
)

//Frequency is int array with size 26 for holding english letters
type Frequency [26]int

//Count the unicode letter. We index for 26 letter. Increment the index for letter found
func Count(word string) (f Frequency) {
	for _, char := range word {
		if unicode.IsLetter(char) {
			f[char-'a']++
		}
	}

	return
}

//Detect finds the new word with same letters present in subject
func Detect(subject string, candidates []string) []string {
	subjectLower := strings.ToLower(subject)
	var anagram []string
	countSubject := Count(subjectLower)

	for _, word := range candidates {
		wordLower := strings.ToLower(word)
		if len(subject) != len(wordLower) || subjectLower == wordLower {
			continue
		}
		countCandidate := Count(wordLower)

		if countSubject == countCandidate {
			anagram = append(anagram, word)
		}

	}
	return anagram
}
