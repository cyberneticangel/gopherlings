// comparable admits exactly the types that support == and !=:
// numbers, strings, booleans, pointers, channels, and arrays or
// structs of such types. That is what lets indexOf compare elements,
// something `any` would not allow.
package main

import "fmt"

// indexOf reports the position of the first element of xs that equals
// v, or -1 if xs does not contain v.
func indexOf[T comparable](xs []T, v T) int {
	for i, x := range xs {
		if x == v {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println("index of b:", indexOf([]string{"a", "b", "c"}, "b"))
	fmt.Println("index of 9:", indexOf([]int{1, 2, 3}, 9))
}
