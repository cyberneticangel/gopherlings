// A slice expression xs[low:high] picks out the elements from
// index low up to but NOT including high -- a "half-open" range:
//
//	xs := []int{1, 2, 3, 4, 5}
//	xs[1:4] // [2 3 4]
//	xs[:2]  // [1 2]  (omit low: start at 0)
//	xs[3:]  // [4 5]  (omit high: run to the end)
//
// A slice expression does not copy: the result shares the backing
// array with xs (remember cloneOf from the last exercise? Reach
// for slices.Clone when you need the pieces to be independent).
//
// Task: make middle return xs without its first and last element:
// the slice from index 1 to len(xs)-1. For slices shorter than 3
// there is no middle: return an empty slice (nil or []int{}).
package main

import "fmt"

// middle returns xs without the first and last element.
func middle(xs []int) []int {
	return xs // TODO: slice from index 1 up to len(xs)-1
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("nums:  ", nums)
	fmt.Println("middle:", middle(nums))
	fmt.Println("short: ", middle([]int{1}))
}

// I AM NOT DONE
