// Slices share their backing array: copying the slice header does
// not copy the elements. slices.Clone allocates a fresh array and
// copies everything over, so the clone can change without the
// original noticing (and vice versa).
package main

import (
	"fmt"
	"slices"
)

// cloneOf returns a copy of xs that shares no memory with it.
func cloneOf(xs []int) []int {
	return slices.Clone(xs)
}

func main() {
	orig := []int{1, 2, 3}
	twin := cloneOf(orig)
	twin[0] = 99
	fmt.Println("original:", orig)
	fmt.Println("clone:   ", twin)
}
