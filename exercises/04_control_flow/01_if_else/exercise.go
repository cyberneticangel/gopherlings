// The if statement runs a block only when its condition is true. Chain
// conditions with `else if` and catch everything left with `else`:
//
//	if x > 0 {
//		// ...
//	} else if x < 0 {
//		// ...
//	} else {
//		// ...
//	}
//
// Note that if takes no parentheses around its condition, but the
// braces are required.
//
// Task: make `sign` return "positive" for x > 0, "negative" for x < 0
// and "zero" for x == 0.
package main

import "fmt"

// sign classifies x as "positive", "negative" or "zero".
func sign(x int) string {
	return "positive" // TODO: only true for x > 0 — handle all three cases
}

func main() {
	fmt.Println("sign(-2) =", sign(-2))
}

// I AM NOT DONE
