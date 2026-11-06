// Nested loops need care when you want to stop early: a plain `break`
// only exits the loop it sits in. The outer loop keeps going and will
// scan right past what you found.
//
// Go's answer is a labeled break. Name the outer loop with a label,
// then break to that label to leave both loops at once:
//
//	outer:
//	for _, row := range grid {
//		for _, v := range row {
//			if v < 0 {
//				break outer
//			}
//		}
//	}
//
// (`continue` understands labels too.) An early `return` would also
// work here; labels shine when you cannot return immediately.
//
// Task: make `firstNegative` scan the grid row by row, left to right,
// and return the first negative number it finds — or 0 if there is
// none. The scan is already written; it just never stops in time.
package main

import "fmt"

// firstNegative returns the first negative number in the grid, scanned
// row by row, left to right — or 0 when the grid has none.
func firstNegative(grid [][]int) int {
	for _, row := range grid {
		for _, v := range row {
			if v < 0 {
				break // TODO: this only leaves the inner loop — and nothing is remembered!
			}
		}
	}
	return 0
}

func main() {
	grid := [][]int{
		{3, 5, 8},
		{1, -4, 9},
		{2, 7, -1},
	}
	fmt.Println("first negative:", firstNegative(grid))
}

// I AM NOT DONE
