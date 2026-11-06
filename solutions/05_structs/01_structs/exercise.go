// Struct literals bundle field values into a struct. The keyed
// form Point{X: 3, Y: 4} is preferred over the unkeyed Point{3, 4}:
// it survives reordering and new fields, and it reads clearly.
package main

import "fmt"

// Point is a coordinate on a 2D grid.
type Point struct {
	X, Y int
}

// samePoint reports whether a and b hold the same coordinates.
func samePoint(a, b Point) bool {
	return a.X == b.X && a.Y == b.Y
}

func main() {
	p := Point{X: 3, Y: 4}
	fmt.Println("p is", p)
	fmt.Println("same:", samePoint(Point{X: 1, Y: 2}, Point{X: 1, Y: 2}))
}
