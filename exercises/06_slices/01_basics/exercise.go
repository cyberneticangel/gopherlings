// A slice is an ordered list of values: []int is a slice of ints.
//
//	xs := []int{10, 20, 30} // a slice literal
//	xs[0]                   // the first element: indexing starts at 0
//	len(xs)                 // the number of elements
//
// `range` walks a slice and hands you the index and the value:
//
//	for i, v := range xs { ... } // use _ for what you don't need
//
// Task: make second return the element at index 1 (the second
// element), and make total return the sum of every element,
// walking the slice with range.
package main

import "fmt"

// second returns the second element of xs (index 1).
func second(xs []int) int {
	return 0 // TODO: index into xs
}

// total returns the sum of all elements of xs.
func total(xs []int) int {
	return 0 // TODO: walk xs with range and add the values up
}

func main() {
	nums := []int{10, 20, 30}
	fmt.Println("slice: ", nums)
	fmt.Println("second:", second(nums))
	fmt.Println("total: ", total(nums))
}

// I AM NOT DONE
