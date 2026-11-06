// The zero value of every pointer is nil: a pointer to nothing.
// Dereferencing nil does not fail at compile time -- it panics at
// RUNTIME:
//
//	var p *string
//	fmt.Println(*p) // panic: invalid memory address or nil dereference
//
// So any function that accepts a pointer must decide what to do with
// a nil one before following it.
//
// Task: make `safeDeref` return "" when p is nil and *p otherwise.
// The test really does call it with nil, so the broken version
// panics -- check the pointer before you follow it.
package main

import "fmt"

// safeDeref returns the string p points at, or "" when p is nil.
func safeDeref(p *string) string {
	return *p // TODO: following a nil pointer panics -- check first
}

func main() {
	name := "Gopher"
	fmt.Println(safeDeref(&name))
	fmt.Println(safeDeref(nil))
}

// I AM NOT DONE
