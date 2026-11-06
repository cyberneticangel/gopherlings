// A pointer receiver (r *Rect) lets a method modify the caller's
// struct: r points at the original, so field writes stick. Reach
// for a pointer receiver whenever a method must mutate its
// receiver, or to avoid copying a large struct.
package main

import "fmt"

// Rect is an axis-aligned rectangle.
type Rect struct {
	W, H float64
}

// Scale multiplies both dimensions of the rectangle by f.
func (r *Rect) Scale(f float64) {
	r.W *= f
	r.H *= f
}

// Area returns the area of the rectangle.
func (r Rect) Area() float64 {
	return r.W * r.H
}

func main() {
	r := Rect{W: 3, H: 4}
	r.Scale(2)
	fmt.Printf("%v, area %v\n", r, r.Area())
}
