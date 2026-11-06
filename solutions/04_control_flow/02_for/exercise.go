// Go has exactly one loop keyword: for. The classic three-part form is
// `for init; condition; post { ... }`; drop the init and the post and
// the same keyword becomes a while loop.
package main

import "fmt"

// sumTo returns 1 + 2 + ... + n.
func sumTo(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}

// powerOfTwo returns 2 to the power of exp.
func powerOfTwo(exp int) int {
	result := 1
	for exp > 0 {
		result *= 2
		exp--
	}
	return result
}

func main() {
	fmt.Println("sumTo(5) =", sumTo(5))
	fmt.Println("powerOfTwo(10) =", powerOfTwo(10))
}
