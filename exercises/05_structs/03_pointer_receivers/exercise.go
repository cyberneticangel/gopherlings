// A value receiver hands the method a COPY of the struct, so any
// change it makes is lost when the method returns:
//
//	func (r Rect) Scale(f float64) // edits a copy: never sticks
//
// To modify the caller's struct, the method needs a pointer
// receiver:
//
//	func (r *Rect) Scale(f float64)
//
// Now r points at the caller's Rect, so field writes like r.W *= f
// land in the original. Go auto-dereferences pointers to structs:
// you write r.W, not (*r).W.
//
// Task: change Scale's receiver from Rect to *Rect so the scaling
// reaches the caller. One little character!
package main

import "fmt"

// Rect is an axis-aligned rectangle.
type Rect struct {
	W, H float64
}

// Scale multiplies both dimensions of the rectangle by f.
func (r Rect) Scale(f float64) {
	r.W *= f // TODO: this edits a copy -- make the receiver a *Rect
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

// I AM NOT DONE
