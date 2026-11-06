// The slice expression xs[low:high] selects the elements from low
// up to, but not including, high -- a half-open range. It copies
// nothing: the result shares xs's backing array.
package main

import "fmt"

// middle returns xs without the first and last element.
func middle(xs []int) []int {
	if len(xs) < 3 {
		return nil
	}
	return xs[1 : len(xs)-1]
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("nums:  ", nums)
	fmt.Println("middle:", middle(nums))
	fmt.Println("short: ", middle([]int{1}))
}
