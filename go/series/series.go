//Package series provide consecutive subsequence of n digits
package series

// All function returns consecutive substrings of lenght n from s
func All(n int, s string) []string {
	var result []string
	if n > len(s) {
		return nil
	}

	for i := 0; i < len(s); i++ {
		result = append(result, string(s[i:i+n]))
	}
	return result
}

//UnsafeFirst returns the first consecutive subsequence of lenght n
func UnsafeFirst(n int, s string) string {

	if n > len(s) {
		return ""
	}

	return string(s[:n])
}
