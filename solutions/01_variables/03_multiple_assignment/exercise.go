// Multiple assignment evaluates the right-hand side first, so
// `a, b = b, a` swaps two values in one statement, with no temporary
// variable.
package main

import "fmt"

// swap returns a and b in opposite order.
func swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func main() {
	x, y := swap(1, 2)
	fmt.Printf("swap(1, 2) = (%d, %d)\n", x, y)
}
