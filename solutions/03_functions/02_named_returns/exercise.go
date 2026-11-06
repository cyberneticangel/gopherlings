// Return values can be named. Assign to the names inside the function,
// then a naked `return` sends their current values back to the caller.
package main

import "fmt"

// split breaks a total into an even half and what is left over.
func split(sum int) (half int, rest int) {
	half = sum / 2
	rest = sum - half
	return
}

func main() {
	half, rest := split(7)
	fmt.Printf("splitting 7: half %d, rest %d\n", half, rest)
}
