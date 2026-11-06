// Go does not have exceptions for everyday failures -- errors are
// ordinary VALUES, returned next to the result:
//
//	result, err := divide(6, 3)
//	if err != nil { ... }
//
// A nil error means success; a non-nil one explains the problem,
// usually built with errors.New("some message").
//
// Task: make `divide` refuse to divide by zero. When b is 0, return
// 0 and an error saying so; otherwise return a / b with a nil error.
// You will need the errors package -- import it alongside fmt.
package main

import "fmt"

// divide returns a / b, or an error when b is zero.
func divide(a, b float64) (float64, error) {
	return 0, nil // TODO: check b first -- dividing by zero needs an error
}

func main() {
	result, err := divide(6, 3)
	if err != nil {
		fmt.Println("oops:", err)
		return
	}
	fmt.Println("6 / 3 =", result)
}

// I AM NOT DONE
