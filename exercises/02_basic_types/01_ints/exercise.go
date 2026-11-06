// Go is strict about types: an operator's behavior depends on its
// operands. When both sides of `/` are integers, `/` is INTEGER
// division -- it truncates toward zero and throws the remainder away:
//
//	7 / 2          // 3, not 3.5
//	float64(7) / 2 // 3.5 -- convert first, then divide
//
// Task: fix average so it computes the mean as a float64. You must
// convert the values BEFORE dividing; converting afterward is too late,
// the fractional part is already gone.
package main

import "fmt"

// average returns the mean of a and b.
func average(a, b int) float64 {
	return float64((a + b) / 2) // TODO: converting too late -- convert before dividing
}

func main() {
	fmt.Printf("average(1, 2) = %v\n", average(1, 2))
}

// I AM NOT DONE
