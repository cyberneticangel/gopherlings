// Slices are ordered lists: index them with xs[i] (starting at 0)
// and walk them with `for i, v := range xs`. len(xs) is the number
// of elements.
package main

import "fmt"

// second returns the second element of xs (index 1).
func second(xs []int) int {
	return xs[1]
}

// total returns the sum of all elements of xs.
func total(xs []int) int {
	sum := 0
	for _, v := range xs {
		sum += v
	}
	return sum
}

func main() {
	nums := []int{10, 20, 30}
	fmt.Println("slice: ", nums)
	fmt.Println("second:", second(nums))
	fmt.Println("total: ", total(nums))
}
