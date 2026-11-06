// A 2D slice is a slice of slices: [][]int. The outer slice holds
// one entry per row, and each row is its own slice, allocated
// separately:
//
//	grid := make([][]int, rows) // rows of nothing yet
//	for i := range grid {
//		grid[i] = make([]int, cols) // each row needs its own make
//	}
//
// Task: build the n x n identity matrix with make and nested
// loops: a 1 on the main diagonal (where the row index equals the
// column index) and 0 everywhere else. identity(3) should be
// [[1 0 0] [0 1 0] [0 0 1]].
package main

import "fmt"

// identity returns the n x n identity matrix.
func identity(n int) [][]int {
	return nil // TODO: make the outer slice, then each row, then set the 1s
}

func main() {
	fmt.Println("identity(3):")
	for _, row := range identity(3) {
		fmt.Println(row)
	}
}

// I AM NOT DONE
