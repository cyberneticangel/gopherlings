// append(xs, v) grows xs by one element and returns the resulting
// slice. Because growing may allocate a new backing array, the
// result can be a different slice value -- always reassign it:
// xs = append(xs, v).
package main

import "fmt"

// push returns xs with v appended at the end.
func push(xs []int, v int) []int {
	return append(xs, v)
}

func main() {
	nums := []int{1, 2}
	nums = push(nums, 3)
	fmt.Println("after push:", nums, "length:", len(nums))
}
