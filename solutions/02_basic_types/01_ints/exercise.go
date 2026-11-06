// int / int is integer division: convert the operands to float64
// BEFORE dividing, or the fractional part is already lost.
package main

import "fmt"

// average returns the mean of a and b.
func average(a, b int) float64 {
	return float64(a+b) / 2
}

func main() {
	fmt.Printf("average(1, 2) = %v\n", average(1, 2))
}
