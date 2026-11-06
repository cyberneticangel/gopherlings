// The fmt package is Go's bread and butter for formatting text.
//
// `fmt.Sprintf` builds a string from a format: verbs like %s (string),
// %d (integer) and %v ("any value") are replaced by the arguments.
//
// Task: make `welcome` return "Welcome to gopherlings, <name>!" using
// fmt.Sprintf. The test checks a few different names.
package main

import "fmt"

// welcome builds a personalized welcome message.
func welcome(name string) string {
	return "" // TODO: build the message with fmt.Sprintf and %s
}

func main() {
	fmt.Println(welcome("Gopher"))
}

// I AM NOT DONE
