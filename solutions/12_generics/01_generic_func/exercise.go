// Generics let one function work on many types. Max declares a type
// parameter T constrained by cmp.Ordered -- every type that supports
// < (all the numbers plus strings). Constraints can also be spelled
// as unions, e.g. [T int | float64 | string], when no named
// constraint fits.
package main

import (
	"cmp"
	"fmt"
)

// Max returns the larger of a and b, for any ordered type.
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println("max int:", Max(3, 9))
	fmt.Println("max float:", Max(1.5, 2.5))
	fmt.Println("max string:", Max("a", "b"))
}
