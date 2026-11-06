// A method is a function with a receiver: func (r Rect) Area().
// The receiver is the value the method is called on -- here a copy
// of the Rect, whose fields it multiplies.
package main

import "fmt"

// Rect is an axis-aligned rectangle.
type Rect struct {
	W, H float64
}

// Area returns the area of the rectangle.
func (r Rect) Area() float64 {
	return r.W * r.H
}

func main() {
	r := Rect{W: 3, H: 4}
	fmt.Printf("A %v rect has area %v\n", r, r.Area())
}
