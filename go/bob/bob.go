//Package bob  provide responses depending o nthe punctuations
package bob

import (
	"fmt"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	fmt.Println(remark)
	// when normal text is found with no shouting(capitals), no questions
	if remark == "" {
		return "Fine. Be that way!"
		//check whether it is fully upper case (shouting) and it is not purely numbers
	} else if strings.ToUpper(remark) == remark && strings.ToUpper(remark) != strings.ToLower(remark) {
		//HAving questions
		if strings.HasSuffix(remark, "?") {
			return "Calm down, I know what I'm doing!"
		}
		//If only capital shouting with no question means yelling
		return "Whoa, chill out!"
		//no need to be capital but normal questions
	} else if strings.HasSuffix(remark, "?") {
		return "Sure."

	} else {
		return "Whatever."
	}
}
