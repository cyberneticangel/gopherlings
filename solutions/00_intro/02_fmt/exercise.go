// The fmt package formats text. Use fmt.Sprintf with the %s verb to build
// the personalized welcome message.
package main

import "fmt"

// welcome builds a personalized welcome message.
func welcome(name string) string {
	return fmt.Sprintf("Welcome to gopherlings, %s!", name)
}

func main() {
	fmt.Println(welcome("Gopher"))
}
