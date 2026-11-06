// A slice has a fixed length, but it can GROW with the built-in
// append function:
//
//	xs = append(xs, v)
//
// append tucks v onto the end. If the slice's backing array has no
// spare room, append allocates a bigger array and returns a NEW
// slice pointing at it -- which is why you always reassign the
// result. Discard it, and the new element is gone.
//
// Task: make push return xs with v added at the end. The caller
// must reassign too: nums = push(nums, 3).
package main

import "fmt"

// push returns xs with v appended at the end.
func push(xs []int, v int) []int {
	return xs // TODO: append v and return the result
}

func main() {
	nums := []int{1, 2}
	nums = push(nums, 3)
	fmt.Println("after push:", nums, "length:", len(nums))
}

// I AM NOT DONE
