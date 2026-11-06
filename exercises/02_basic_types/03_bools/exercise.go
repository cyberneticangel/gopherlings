// Comparisons produce bool values: ==, !=, <, <=, >, >=. Bools combine
// with three logic operators:
//
//	&& -- and: true only when BOTH sides are true
//	|| -- or:  true when EITHER side is true
//	!  -- not: flips true and false
//
// Task: make inRange report whether x lies between lo and hi,
// INCLUSIVE: x must be >= lo and <= hi. Fix the comparison and check
// both bounds with &&.
package main

import "fmt"

// inRange reports whether lo <= x <= hi.
func inRange(x, lo, hi int) bool {
	return x > lo // TODO: wrong comparison, and hi is never checked
}

func main() {
	fmt.Printf("inRange(5, 1, 10) = %t\n", inRange(5, 1, 10))
}

// I AM NOT DONE
