package accumulate

// Accumulate return string after the requested conversion
func Accumulate(given []string, converter func(string) string) (result []string) {

	for _, word := range given {

		result = append(result, converter(word))
	}
	return
}
