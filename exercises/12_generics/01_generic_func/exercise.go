// Generics let one function work on many types. Until Go 1.18 the only
// way was duplication: MaxInt for ints, MaxFloat for floats -- exactly
// the pair sitting below.
//
// A generic function declares type parameters in square brackets:
//
//	func Max[T int | float64](a, b T) T
//
// T is the type parameter, and `int | float64` is its CONSTRAINT: the
// set of types T may stand for. The standard library's cmp package
// ships a ready-made constraint, cmp.Ordered, which covers every type
// that supports < -- all the numbers plus strings.
//
// Task: replace MaxInt and MaxFloat with ONE generic function
//
//	func Max[T cmp.Ordered](a, b T) T
//
// that returns the larger of a and b. (You may also spell the
// constraint yourself as a union: [T int | float64 | string] -- the
// test calls Max with strings too.) Remember to import "cmp" and to
// update main.
package main

import "fmt"

// TODO: replace MaxInt and MaxFloat below with one generic Max.

// MaxInt returns the larger of two ints.
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxFloat returns the larger of two float64s.
func MaxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("max int:", MaxInt(3, 9))
	fmt.Println("max float:", MaxFloat(1.5, 2.5))
}

// I AM NOT DONE
