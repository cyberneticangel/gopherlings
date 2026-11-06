// A union constraint like [T int | float64] lists exactly the types a
// type parameter may be. The compiler then knows `+` is legal in the
// body -- with `any` it would not be, since an `any` could be a
// struct or a func, which have no +.
package main

import "fmt"

// sum returns the total of every element in xs.
func sum[T int | float64](xs []T) T {
	var total T
	for _, x := range xs {
		total += x
	}
	return total
}

func main() {
	fmt.Println("sum ints:", sum([]int{1, 2, 3}))
	fmt.Println("sum floats:", sum([]float64{1.5, 2.5}))
}
