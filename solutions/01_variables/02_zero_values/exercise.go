// Every type has a zero value: 0 for int, "" for string, false for
// bool. Variables declared with `var` and no initializer start there,
// with no assignment needed.
package main

import "fmt"

// zeroValues returns the zero values of int, string and bool.
func zeroValues() (int, string, bool) {
	var n int
	var s string
	var b bool
	return n, s, b
}

func main() {
	n, s, b := zeroValues()
	fmt.Printf("zero int: %d, zero string: %q, zero bool: %t\n", n, s, b)
}
