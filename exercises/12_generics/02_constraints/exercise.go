// A constraint says which types a type parameter may stand for. The
// loosest is `any` -- every type -- but loose is not always useful:
// with T declared as `any`, T could be a struct or a func, and `+` is
// not defined for those. That is why `total + x` does not compile
// when the constraint is `any`.
//
// A UNION constraint lists exactly the types you need, so the
// compiler knows the body's operations are allowed:
//
//	func sum[T int | float64](xs []T) T
//
// Task: make sum do its job. Range over xs, add every element to
// total, and return the total.
package main

import "fmt"

// sum returns the total of every element in xs.
func sum[T int | float64](xs []T) T {
	var total T
	// TODO: add every element of xs to total
	return total
}

func main() {
	fmt.Println("sum ints:", sum([]int{1, 2, 3}))
	fmt.Println("sum floats:", sum([]float64{1.5, 2.5}))
}

// I AM NOT DONE
