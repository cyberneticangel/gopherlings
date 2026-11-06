// & takes the address of a value and * follows a pointer back to
// it. Go has no pointer arithmetic -- pointers are only ever taken
// and followed, never shifted around like in C.
package main

import "fmt"

// double multiplies the int that p points at by 2.
func double(p *int) {
	*p *= 2
}

func main() {
	x := 21
	double(&x)
	fmt.Println("x is now", x)
}
