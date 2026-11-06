// Variables are statically typed: a variable's declared type must match
// the type of the value it holds (5 is an int, true is a bool, 3.14 is
// a float64). With no type given, Go infers it from the value.
package main

import "fmt"

var section = "variables"

var count int = 5

var ready bool = true

var pi float64 = 3.14

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
