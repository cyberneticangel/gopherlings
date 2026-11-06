// Functions in Go can return more than one value. List the types after
// the parameters and return the values as a comma-separated list.
package main

import "fmt"

// divMod returns the quotient and the remainder of a divided by b.
func divMod(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	q, r := divMod(17, 5)
	fmt.Printf("17 / 5 = %d, remainder %d\n", q, r)
}
