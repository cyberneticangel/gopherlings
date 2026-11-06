// A variable declared without a value is not empty in Go: every type
// has a "zero value" that new variables start out with.
//
//	int    -> 0
//	string -> "" (the empty string)
//	bool   -> false
//
// This is why Go has no "used before initialized" surprises: the zero
// value is always there, without you assigning anything.
//
// Task: make zeroValues declare one variable of each type with `var`
// and NO initializer, then return all three. No assignments needed --
// the zero values come for free.
package main

import "fmt"

// zeroValues returns the zero values of int, string and bool.
func zeroValues() (int, string, bool) {
	// TODO: declare n (int), s (string) and b (bool) with `var` and
	// no initializer, then return them instead of the values below.
	return 1, "x", true
}

func main() {
	n, s, b := zeroValues()
	fmt.Printf("zero int: %d, zero string: %q, zero bool: %t\n", n, s, b)
}

// I AM NOT DONE
