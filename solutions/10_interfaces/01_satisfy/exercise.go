// A type satisfies an interface implicitly, just by having the
// required methods. No `implements` keyword -- which also means your
// types can satisfy interfaces declared in packages you do not
// control.
package main

import "fmt"

// Speaker is implemented by anything with a Speak method.
type Speaker interface {
	Speak() string
}

// Dog is a very loud animal.
type Dog struct{}

// Speak is Dog's sound.
func (d Dog) Speak() string {
	return "Woof"
}

// Cat satisfies Speaker without announcing it anywhere.
type Cat struct{}

// Speak is Cat's sound.
func (c Cat) Speak() string {
	return "Meow"
}

// announce describes what a Speaker says.
func announce(s Speaker) string {
	return "It says: " + s.Speak()
}

func main() {
	var s Speaker = Dog{}
	fmt.Println(announce(s))
}
