// A 2D slice is a slice of slices: make([][]int, n) makes the n
// rows, and each row is its own slice that needs its own make.
// The identity matrix has a 1 wherever the row index equals the
// column index.
package main

import "fmt"

// identity returns the n x n identity matrix.
func identity(n int) [][]int {
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
		for j := range grid[i] {
			if i == j {
				grid[i][j] = 1
			}
		}
	}
	return grid
}

func main() {
	fmt.Println("identity(3):")
	for _, row := range identity(3) {
		fmt.Println(row)
	}
}
