// Pointers hold the memory address of a value. Two operators do all
// the work:
//
//	&x  takes the address of x   (type *int when x is an int)
//	*p  follows the pointer      (the int that p points at)
//
// Unlike C, Go has NO pointer arithmetic: you can never write p+1.
// Pointers are only ever taken and followed.
//
// Task: make `double` double the value THROUGH the pointer, so the
// caller's variable really changes. One dereference is all it takes.
package main

import "fmt"

// double multiplies the int that p points at by 2.
func double(p *int) {
	p = p * 2 // TODO: p is an address, not a number -- dereference it first
}

func main() {
	x := 21
	double(&x)
	fmt.Println("x is now", x)
}

// I AM NOT DONE
