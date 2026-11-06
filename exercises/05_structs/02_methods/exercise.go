// Methods are functions attached to a type.
//
// A method looks like a function with an extra "receiver" in front
// of its name:
//
//	func (r Rect) Area() float64
//
// The receiver r is a value of type Rect -- a copy of whichever
// Rect the method was called on. Inside the method you use its
// fields like any other struct value: r.W, r.H.
//
// Task: make Area return the rectangle's area: width times height.
package main

import "fmt"

// Rect is an axis-aligned rectangle.
type Rect struct {
	W, H float64
}

// Area returns the area of the rectangle.
func (r Rect) Area() float64 {
	return 0 // TODO: multiply the width by the height
}

func main() {
	r := Rect{W: 3, H: 4}
	fmt.Printf("A %v rect has area %v\n", r, r.Area())
}

// I AM NOT DONE
