// Assigning a slice does NOT copy its elements:
//
//	a := []int{1, 2, 3}
//	b := a        // b and a now SHARE the same backing array
//	b[0] = 99     // a[0] is 99 too -- surprise!
//
// Both names point at the same memory, because a slice value is
// just a header: pointer, length, capacity. To get an independent
// copy, copy the elements: slices.Clone(xs) allocates a fresh
// backing array and copies everything over.
//
// Task: make cloneOf return an independent copy of xs. The test
// writes through your clone and checks the original is untouched.
package main

import "fmt"

// cloneOf returns a copy of xs that shares no memory with it.
func cloneOf(xs []int) []int {
	return xs // TODO: return a real copy, not the same slice
}

func main() {
	orig := []int{1, 2, 3}
	twin := cloneOf(orig)
	twin[0] = 99
	fmt.Println("original:", orig)
	fmt.Println("clone:   ", twin)
}

// I AM NOT DONE
