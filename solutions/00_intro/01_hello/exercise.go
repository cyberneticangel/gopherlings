// Welcome to gopherlings! 🐹
//
// Fix the compile error by returning the `message` variable — Go is
// case-sensitive.
package main

import "fmt"

// greeting builds the message gopherlings greets you with.
func greeting() string {
	message := "Hello, Gopher!"
	return message
}

func main() {
	fmt.Println(greeting())
}
