// Package triangle finds the type of triangle
package triangle

import (
	"fmt"
	"math"
)

// Kind represents the kind fo triangle
type Kind int

//Const to hold various traingles
const (
	// Pick values for the following identifiers used by the test program.
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides finds the type of triangle with the input sides
func KindFromSides(a, b, c float64) Kind {
	// if sum of two sides is greater than or equal to third side, it iis atrianlge
	// if 3 sides equal - equilatrial , 2 sides - isoceles . all diff sides - scalene triangle
	var k Kind
	fmt.Println(a, b, c)
	if a <= 0 || b <= 0 || c <= 0 || a+b < c || a+c < b || c+b < a || math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		k = NaT
	} else if a == b && a == c {
		k = Equ
	} else if a == b && a != c || b == c && a != b || a == c && a != b {
		k = Iso
	} else if a != b && a != c && b != c {
		k = Sca
	} else {
		k = NaT
	}

	return k
}
