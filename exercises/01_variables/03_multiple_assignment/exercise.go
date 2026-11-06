// Go can assign to more than one variable in a single statement:
//
//	x, y := 1, 2 // declare both at once (the short form)
//	x, y = y, x  // swap them -- no temporary variable needed
//
// The right-hand side is fully evaluated before anything is assigned,
// so the swap above just works. This is the idiomatic Go swap.
//
// Task: make swap exchange its two arguments using ONE multiple
// assignment (no temporary variable, no arithmetic tricks), then return
// them in their new order.
package main

import "fmt"

// swap returns a and b in opposite order.
func swap(a, b int) (int, int) {
	// TODO: exchange a and b with a single multiple assignment
	return a, b
}

func main() {
	x, y := swap(1, 2)
	fmt.Printf("swap(1, 2) = (%d, %d)\n", x, y)
}

// I AM NOT DONE
