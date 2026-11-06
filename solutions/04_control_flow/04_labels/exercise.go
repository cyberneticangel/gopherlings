// A plain `break` inside nested loops only exits the inner loop. Label
// the outer loop and `break outer` to leave both loops at once.
package main

import "fmt"

// firstNegative returns the first negative number in the grid, scanned
// row by row, left to right — or 0 when the grid has none.
func firstNegative(grid [][]int) int {
	first := 0
outer:
	for _, row := range grid {
		for _, v := range row {
			if v < 0 {
				first = v
				break outer
			}
		}
	}
	return first
}

func main() {
	grid := [][]int{
		{3, 5, 8},
		{1, -4, 9},
		{2, 7, -1},
	}
	fmt.Println("first negative:", firstNegative(grid))
}
