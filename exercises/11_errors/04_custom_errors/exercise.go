// Sometimes a message is not enough -- callers want the DATA inside
// an error: which field failed, what code, what limit. Give the
// error its own type; any type with an Error() string method IS an
// error, satisfied implicitly like every other interface:
//
//	type ValidationError struct {
//		Field string
//	}
//
//	func (e *ValidationError) Error() string {
//		return fmt.Sprintf("invalid value for field %q", e.Field)
//	}
//
// errors.As then digs a matching type out of an error chain:
//
//	var ve *ValidationError
//	if errors.As(err, &ve) {
//		fmt.Println("bad field:", ve.Field)
//	}
//
// Task: make `validateAge` return a *ValidationError for the "age"
// field when age is negative, and nil otherwise. The plain
// fmt.Errorf it returns now is an error too -- but no caller can
// reach a Field inside it.
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
		return fmt.Errorf("invalid age: %d", age) // TODO: return a *ValidationError for the "age" field instead
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

// I AM NOT DONE
