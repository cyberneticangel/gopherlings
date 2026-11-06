// Errors are ordinary values returned next to the result. The caller
// checks `if err != nil` before trusting the result; errors.New
// builds a simple error carrying a message.
package main

import (
	"errors"
	"fmt"
)

// divide returns a / b, or an error when b is zero.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(6, 3)
	if err != nil {
		fmt.Println("oops:", err)
		return
	}
	fmt.Println("6 / 3 =", result)
}
