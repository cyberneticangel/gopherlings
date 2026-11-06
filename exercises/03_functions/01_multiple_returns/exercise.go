// Functions in Go can return more than one value. List the types after
// the parameters and return the values as a comma-separated list.
//
// Callers receive all the values at once:
//
//	q, r := divMod(17, 5) // q is 3, r is 2
//
// Task: make `divMod` return the quotient (a / b) and the remainder
// (a % b) as two separate values.
package main

import "fmt"

// divMod returns the quotient and the remainder of a divided by b.
func divMod(a, b int) (int, int) {
	return 0, 0 // TODO: return the quotient and the remainder
}

func main() {
	q, r := divMod(17, 5)
	fmt.Printf("17 / 5 = %d, remainder %d\n", q, r)
}

// I AM NOT DONE
