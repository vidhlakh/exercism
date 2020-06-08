package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency calculate the frequency in concurrent manner using Goroutines
func ConcurrentFrequency(allLang []string) FreqMap {

	msg := make(chan FreqMap)

	for _, input := range allLang {

		go func(input string) {
			msg <- Frequency(input)
		}(input)
	}
	result := make(FreqMap)
	for range allLang {
		for key, value := range <-msg {
			result[key] += value
		}
	}
	return result

}
