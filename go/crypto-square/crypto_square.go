package cryptosquare

import (
	"log"
	"regexp"
	"strings"
)

//Encode encodes the given string into rectangle
func Encode(pt string) string {
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(pt, "")
	processedString = strings.ToLower(processedString)
	lenInput := len(processedString)
	//calculate col, row
	col, row := calculateColumnRow(lenInput)

	//Encoding the text
	if lenInput < 2 {
		return processedString
	}

	result := strings.Builder{}
	for i := 0; i < col; i++ {
		for j := 0; j < row; j++ {
			position := (j * col) + i
			if position >= len(processedString) {
				result.WriteRune(' ')
				continue
			}
			result.WriteByte(processedString[position])
		}
		if i != (col - 1) {
			result.WriteRune(' ')
		}
	}
	return result.String()

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
