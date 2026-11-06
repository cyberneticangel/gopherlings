// One interface, many types: that is polymorphism in Go. A []Speaker
// can hold Dogs, Cats and Robots -- anything with a Speak method --
// and the compiler checked they all qualify when the slice was
// built.
//
// Code written against the interface, like speakAll below, never
// needs to know which concrete type it is handling. Add a new
// Speaker tomorrow and speakAll keeps working unchanged.
//
// Task: make `speakAll` range over the speakers and collect each
// one's Speak() into a []string, preserving order.
package main

import "fmt"

// Speaker is implemented by anything with a Speak method.
type Speaker interface {
	Speak() string
}

// Dog speaks one way ...
type Dog struct{}

// Speak is Dog's sound.
func (d Dog) Speak() string { return "Woof" }

// Cat speaks another.
type Cat struct{}

// Speak is Cat's sound.
func (c Cat) Speak() string { return "Meow" }

// Robot does not breathe, but it still speaks.
type Robot struct{}

// Speak is Robot's sound.
func (r Robot) Speak() string { return "Beep" }

// speakAll collects what every speaker says, in order.
func speakAll(speakers []Speaker) []string {
	return nil // TODO: range over speakers and collect each Speak() into a slice
}

func main() {
	zoo := []Speaker{Dog{}, Cat{}, Robot{}}
	fmt.Println(speakAll(zoo))
}

// I AM NOT DONE
