// Go has exactly one loop keyword: for. It covers every loop form you
// know from other languages.
//
// The classic three-part form — init; condition; post:
//
//	for i := 1; i <= n; i++ {
//		total += i
//	}
//
// Leave out the init and the post and you have a while loop. There is
// no `while` keyword in Go:
//
//	for value > 1 {
//		value = value / 2
//	}
//
// Task: fill in both functions.
//
//   - sumTo adds up the numbers 1 through n, using the three-part form.
//   - powerOfTwo returns 2^exp by doubling a result exp times, using
//     the while-style form. Count exp down inside the loop, or it will
//     never end!
package main

import "fmt"

// sumTo returns 1 + 2 + ... + n.
func sumTo(n int) int {
	return 0 // TODO: loop with `for i := 1; i <= n; i++`
}

// powerOfTwo returns 2 to the power of exp.
func powerOfTwo(exp int) int {
	return 0 // TODO: loop with `for exp > 0`, doubling each round
}

func main() {
	fmt.Println("sumTo(5) =", sumTo(5))
	fmt.Println("powerOfTwo(10) =", powerOfTwo(10))
}

// I AM NOT DONE
