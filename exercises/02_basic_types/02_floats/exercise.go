// Floats are approximations, not exact decimals. Many decimal numbers
// cannot be represented exactly in binary, so tiny errors creep in:
//
//	0.1 + 0.2 == 0.3 // false! the left side is 0.30000000000000004...
//
// Because of this, comparing floats with == is unreliable. The fix is a
// tolerance (an "epsilon"): the values count as equal when their
// difference is smaller than some tiny threshold. math.Abs gives the
// absolute value, so the check works in both directions.
//
// Task: make nearlyEqual use an epsilon of 1e-9 instead of ==.
package main

import (
	"fmt"
	"math"
)

// nearlyEqual reports whether a and b are equal within 1e-9.
func nearlyEqual(a, b float64) bool {
	return a == b // TODO: too strict -- compare the difference against an epsilon
}

func main() {
	diff := math.Abs((0.1 + 0.2) - 0.3)
	fmt.Printf("0.1 + 0.2 - 0.3 = %g, but == says equal? %t\n",
		diff, nearlyEqual(0.1+0.2, 0.3))
}

// I AM NOT DONE
