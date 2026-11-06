// Code written against an interface works for every type that
// satisfies it. speakAll never asks which concrete types it is
// holding -- it just calls Speak, and dynamic dispatch does the
// rest. That is polymorphism, Go style.
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
	said := make([]string, 0, len(speakers))
	for _, s := range speakers {
		said = append(said, s.Speak())
	}
	return said
}

func main() {
	zoo := []Speaker{Dog{}, Cat{}, Robot{}}
	fmt.Println(speakAll(zoo))
}
