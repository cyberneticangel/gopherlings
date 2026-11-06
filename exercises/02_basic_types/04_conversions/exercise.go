// Strings and numbers are different types, and Go never converts
// between them silently: `"2" * 2` is a compile error, not "22" or 4.
// Conversions go through the strconv package:
//
//	n, err := strconv.Atoi("42") // string -> int
//	s := strconv.Itoa(42)        // int -> string
//
// Atoi returns two results: the number and an error. When the string is
// not a valid number the error is set and the number is 0. For now we
// ignore the error with _ -- handling errors properly gets its own
// section later.
//
// Task: fix doubled so it parses s with strconv.Atoi, doubles the
// number, and returns it formatted with strconv.Itoa. Input that cannot
// be parsed counts as 0, so doubled("") is "0".
package main

import "fmt"

// doubled parses s as an integer and returns double the value, as a
// string. Unparseable input counts as 0.
func doubled(s string) string {
	n := s * 2 // TODO: s is a string, not a number -- parse it first
	return s   // TODO: then double n and format it back to a string
}

func main() {
	fmt.Println("doubled(\"21\") =", doubled("21"))
	fmt.Println("doubled(\"abc\") =", doubled("abc"))
}

// I AM NOT DONE
