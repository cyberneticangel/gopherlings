// Go interfaces are satisfied IMPLICITLY: a type implements an
// interface simply by having the required methods -- there is no
// `implements` keyword anywhere. If it has a Speak method, it is a
// Speaker.
//
// Someone gave Dog a Bark method instead, so Dog does not satisfy
// Speaker and the assignment below refuses to compile:
//
//	cannot use Dog{} as Speaker value: Dog does not implement
//	Speaker (missing method Speak)
//
// Task:
//
//  1. Rename Dog's Bark method to Speak.
//  2. Make announce return "It says: " + s.Speak() instead of "".
package main

import "fmt"

// Speaker is implemented by anything with a Speak method.
type Speaker interface {
	Speak() string
}

// Dog is a very loud animal.
type Dog struct{}

// Bark is Dog's sound ... but Speaker wants Speak!
func (d Dog) Bark() string {
	return "Woof" // TODO: rename this method to Speak
}

// Cat already satisfies Speaker, without announcing it anywhere.
type Cat struct{}

// Speak is Cat's sound.
func (c Cat) Speak() string {
	return "Meow"
}

// announce describes what a Speaker says.
func announce(s Speaker) string {
	return "" // TODO: return "It says: " + s.Speak()
}

func main() {
	var s Speaker = Dog{}
	fmt.Println(announce(s))
}

// I AM NOT DONE
