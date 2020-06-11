package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// cryptosquare
func crypto() {
	result := strings.Builder{}
	st := "Hello425"
	for _, letter := range st {
		result.WriteRune(letter)
	}
	fmt.Println("Result:", result.String())
}

//normalize
func normalize() string {
	// Make a Regex to say we only want letters and numbers
	pt := "s#$%^&plunk"
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(pt, "")
	processedString = strings.ToLower(processedString)
	fmt.Println("Processed string:", processedString)
	return processedString
}

// CalculateColumnRow calculate the column and the row of normalized input where c >=r and c - r <= 1
func calculateColumnRow(lenInput int) (col, row int) {
	for i := 1; i <= lenInput; i++ {
		for j := 1; j <= lenInput; j++ {
			if i*j >= lenInput && i >= j && i-j <= 1 {
				col = i
				row = j
				return
			}
		}
	}

	return
}

// EncodeText encode the input by reading down the columns going left to right
func encodeText(input string, col, row int) string {
	if len(input) < 2 {
		return input
	}

	result := strings.Builder{}
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			position := (j * col) + i
			if position >= len(input) {
				result.WriteRune(' ')
				continue
			}
			result.WriteByte(input[position])
		}
		if i != (col - 1) {
			result.WriteRune(' ')
		}
	}
	return result.String()
}
func main() {
	crypto()
	data := normalize()
	col, row := calculateColumnRow(len(data))
	fmt.Println("col:", col, "row:", row)
	result := encodeText(data, col, row)
	fmt.Println("Encoded text : ", result)
}
