// Your first line of Go: return the sum with the + operator.
package main

import "fmt"

// add returns the sum of a and b.
func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("2 + 3 =", add(2, 3))
}
