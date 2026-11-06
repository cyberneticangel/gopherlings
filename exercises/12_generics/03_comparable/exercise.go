// The comparable constraint marks every type that supports == and !=:
// numbers, strings, booleans, pointers, channels, and arrays or
// structs built from those. Plain `any` does NOT -- you cannot run
// `x == v` when x might be a slice or a func, so the compiler refuses
// the comparison. comparable is the standard constraint for "I only
// need to compare values".
//
// Task: implement indexOf. Loop over xs and return the index of the
// FIRST element equal to v; if nothing matches, return -1.
package main

import "fmt"

// indexOf reports the position of the first element of xs that equals
// v, or -1 if xs does not contain v.
func indexOf[T comparable](xs []T, v T) int {
	return -1 // TODO: find the first element equal to v and return its index
}

func main() {
	fmt.Println("index of b:", indexOf([]string{"a", "b", "c"}, "b"))
	fmt.Println("index of 9:", indexOf([]int{1, 2, 3}, 9))
}

// I AM NOT DONE
