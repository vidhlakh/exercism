package hamming

import "errors"

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
	//When len of string are different or when empty string are passed
	return 0, errors.New("strings of unequal length or empty")
}
