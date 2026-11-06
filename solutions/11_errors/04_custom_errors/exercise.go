// A custom error type carries data, not just a message: any type
// with an Error() string method satisfies the error interface
// implicitly. errors.As searches the chain for that concrete type
// and stores it in the target, so callers can read fields like
// Field.
package main

import (
	"errors"
	"fmt"
)

// ValidationError reports the field whose value did not check out.
type ValidationError struct {
	Field string
}

// Error makes *ValidationError an error.
func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid value for field %q", e.Field)
}

// validateAge checks that age is not negative.
func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age"}
	}
	return nil
}

func main() {
	err := validateAge(-3)
	if err == nil {
		return
	}
	fmt.Println(err)

	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println("bad field:", ve.Field)
	}
}
