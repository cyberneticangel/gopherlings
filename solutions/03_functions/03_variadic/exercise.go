// A variadic function takes any number of arguments of one type. Three
// dots before the type collect them all; inside the function, `nums` is
// an ordinary []int slice.
package main

import "fmt"

// maxOf returns the largest of its arguments. nums is never empty.
func maxOf(nums ...int) int {
	largest := nums[0]
	for _, n := range nums {
		if n > largest {
			largest = n
		}
	}
	return largest
}

func main() {
	fmt.Println("max of 3, 1, 4, 1, 5 is", maxOf(3, 1, 4, 1, 5))
}
