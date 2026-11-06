// Structs bundle related values into one named type.
//
// A struct type lists its fields:
//
//	type Point struct {
//		X, Y int
//	}
//
// You build a value with a struct literal. The unkeyed form
// Point{3, 4} must list EVERY field, in order, and it breaks
// silently when the struct gains or reorders fields -- so prefer
// the keyed form Point{X: 3, Y: 4}, which names each field.
//
// This program does not compile: Point{3} uses the unkeyed form
// but supplies only one of two values. Read the compiler error,
// fix the literal, then write samePoint.
//
// Task: fix `p` in main using keyed fields, and make samePoint
// report whether two points have equal X and equal Y.
package main

import "fmt"

// Point is a coordinate on a 2D grid.
type Point struct {
	X, Y int
}

// samePoint reports whether a and b hold the same coordinates.
func samePoint(a, b Point) bool {
	return false // TODO: compare a's and b's X and Y fields
}

func main() {
	p := Point{3} // TODO: keyed fields: Point{X: 3, Y: 4}
	fmt.Println("p is", p)
	fmt.Println("same:", samePoint(Point{X: 1, Y: 2}, Point{X: 1, Y: 2}))
}

// I AM NOT DONE
