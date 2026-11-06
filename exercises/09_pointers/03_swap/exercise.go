// Go is call-by-value: a function receives COPIES of its arguments,
// so this version of swap exchanges two copies and leaves the
// caller's variables untouched. To let a function change the
// caller's variables, pass pointers:
//
//	swap(&x, &y)
//
// Pointers plus Go's multiple assignment make the exchange a
// one-liner: a, b = b, a becomes *a, *b = *b, *a once a and b are
// pointers.
//
// Task: change swap's parameters from int to *int, and swap the
// values they point at with a single multiple assignment. The call
// sites below (and in the test) already pass addresses.
package main

import "fmt"

// swap exchanges the values pointed at by a and b.
func swap(a, b int) {
	a, b = b, a // TODO: take *int parameters and swap through the pointers
}

func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println("after swap:", x, y)
}

// I AM NOT DONE
