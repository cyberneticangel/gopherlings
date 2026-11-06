// A variadic function takes any number of arguments of one type. Three
// dots before the type collect them all:
//
//	func maxOf(nums ...int) int
//
// Inside the function, `nums` is an ordinary []int slice, so you can
// range over it:
//
//	for _, n := range nums {
//		// ...
//	}
//
// Task: make `maxOf` return the largest number. You may assume at least
// one argument is always passed. Start the max at the first element,
// not at 0 — maxOf must also work when every argument is negative.
package main

import "fmt"

// maxOf returns the largest of its arguments. nums is never empty.
func maxOf(nums ...int) int {
	return 0 // TODO: range over nums and keep track of the largest value
}

func main() {
	fmt.Println("max of 3, 1, 4, 1, 5 is", maxOf(3, 1, 4, 1, 5))
}

// I AM NOT DONE
