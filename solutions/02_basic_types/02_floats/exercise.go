// Floats carry tiny rounding errors (0.1 + 0.2 is not exactly 0.3), so
// compare them with a tolerance: equal when math.Abs(a-b) is smaller
// than an epsilon such as 1e-9.
package main

import (
	"fmt"
	"math"
)

// nearlyEqual reports whether a and b are equal within 1e-9.
func nearlyEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func main() {
	diff := math.Abs((0.1 + 0.2) - 0.3)
	fmt.Printf("0.1 + 0.2 - 0.3 = %g, but == says equal? %t\n",
		diff, nearlyEqual(0.1+0.2, 0.3))
}
