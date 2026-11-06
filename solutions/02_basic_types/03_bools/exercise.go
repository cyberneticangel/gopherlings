// Combine comparisons with && (and), || (or) and ! (not). Inclusive
// ranges read naturally: x >= lo && x <= hi.
package main

import "fmt"

// inRange reports whether lo <= x <= hi.
func inRange(x, lo, hi int) bool {
	return x >= lo && x <= hi
}

func main() {
	fmt.Printf("inRange(5, 1, 10) = %t\n", inRange(5, 1, 10))
}
