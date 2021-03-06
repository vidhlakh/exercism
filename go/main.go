package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

//Frequency is int array with size 26 for holding english letters
type Frequency [26]int

// Distance provide difference b/w 2 DNA strands
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

// IsIsogram finds whether word is an isogram with no repetitions
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
		}
	}
	return true // if all letters are unique
}

//Reverse Function
func Reverse(word string) string {
	result := ""
	for _, letter := range word {
		result = string(letter) + result
	}
	return result
}

// Convert print Plint plong plang depending on th e number
func Convert(num int) string {
	result := ""
	if num%3 == 0 {
		result = "Pling"
	}
	if num%5 == 0 {
		result = result + "Plang"
	}
	if num%7 == 0 {
		result = result + "Plong"
	}
	if num%3 != 0 && num%5 != 0 && num%7 != 0 {
		result = strconv.Itoa(num)
	}

	return result
}

// Hey respons s according to questions
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	fmt.Println(remark)
	if remark == "" {
		return "Fine. Be that way!"

	} else if strings.ToUpper(remark) == remark && strings.ToUpper(remark) != strings.ToLower(remark) {
		if strings.HasSuffix(remark, "?") {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"

	} else if strings.HasSuffix(remark, "?") {
		return "Sure."

	} else {
		return "Whatever."
	}
}

// Abbreviate creates acronym for the words
func Abbreviate(s string) string {
	words := strings.Fields(strings.Replace(strings.Replace(s, "-", " ", -1), "_", "", -1))
	fmt.Println("words ", words)
	acronym := ""
	for _, word := range words {
		acronym += string(word[0])
	}

	return acronym
}

//Count the unicode letter. We index for 26 letter. Increment the index for letter found
func Count(word string) (f Frequency) {
	for _, char := range word {
		if unicode.IsLetter(char) {
			f[char-'a']++
		}
	}

	return
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

	//call reverse function
	ww := "SweetGirl"
	reversed := Reverse(ww)
	fmt.Printf("Reversed string for %s is %s", ww, reversed)

	//call convert for raindrops function
	raindropsResult := Convert(34)
	fmt.Println("Raindrops : ", raindropsResult)

	// call bob's hey
	bobResponse := Hey("Tom-ay-to, tom-aaaah-to.")
	fmt.Println("Bob response", bobResponse)

	// Acronym for Halley's Comet
	acronym := Abbreviate("Halley's - Comet")
	fmt.Println("Acronym:", acronym)
	// Alternative fo my Acronym function
	// Good use of Regex
	pattern := regexp.MustCompile("[\\s_-]+")
	tokens := pattern.Split("vidhya lakshmdf -Gsfd - Dvf _Done", -1)
	fmt.Println(tokens)
	abb := ""
	for _, token := range tokens {
		abb += string(token[0])
	}
	fmt.Println(abb)

	// []string
	given := []string{"Cat", "Dofg", "ffsdf"}
	for _, word := range given {
		fmt.Println(word)
	}
	//map ispangram
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alpha := make(map[string]int)
	for _, a := range alphabet {
		alpha[string(a)]++
	}
	fmt.Println("map alpha:", alpha)
	inputString := strings.ToLower("the quickk brown fox jumps over the lazy dog")
	for _, letter := range inputString {
		if alpha[string(letter)] > 1 {
			alpha[string(letter)]--
		} else {
			delete(alpha, string(letter))
		}
	}
	if len(alpha) == 0 {
		fmt.Println("map alpha:", alpha)
	}
	//Anagram
	subject := "master"
	candidates := []string{"stream", "pigeon", "maters"}
	subjectLower := strings.ToLower(subject)
	var anagram []string
	countSubject := Count(subjectLower)

	for _, word := range candidates {
		wordLower := strings.ToLower(word)
		if len(subject) != len(wordLower) {
			continue
		}
		countCandidate := Count(wordLower)

		if countSubject == countCandidate {
			anagram = append(anagram, word)
		}

	}
	fmt.Println("Anagram ", anagram)
}
