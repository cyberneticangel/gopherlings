// Variables in Go are statically typed: every variable declares the
// type of the values it can hold, and the compiler checks each
// declaration before your program ever runs.
//
// Three common ways to declare a variable:
//
//	var count int = 5 // explicit type
//	var count = 5     // type inferred from the value
//	count := 5        // short form (inside functions only)
//
// Task: the declarations below name a sensible value but declare the
// WRONG type for it. Fix each declared type so the program compiles,
// and keep the values exactly as they are. TODO comments mark the
// spots.
package main

import "fmt"

var section = "variables" // no type given: Go infers string. Fine!

var count string = 5 // TODO: 5 is not a string -- fix the type

var ready string = true // TODO: true is not a string either

var pi int = 3.14 // TODO: 3.14 is not an int

// sectionName reports which section of gopherlings you are in.
func sectionName() string {
	return section
}

// gopherCount reports how many gophers are here to learn.
func gopherCount() int {
	return count
}

// isReady reports whether the gophers are ready to start.
func isReady() bool {
	return ready
}

// piApprox reports an approximation of pi.
func piApprox() float64 {
	return pi
}

func main() {
	fmt.Printf("%s: %d gophers, ready: %t, pi ~ %v\n",
		sectionName(), gopherCount(), isReady(), piApprox())
}

// I AM NOT DONE
