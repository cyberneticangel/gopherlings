// A switch with no expression after the keyword — a "tagless" switch —
// is Go's tidier take on the if/else chain. Every case is a full
// condition, cases are tested from top to bottom, and the first true
// one runs:
//
//	switch {
//	case temp > 30:
//		fmt.Println("hot")
//	case temp > 15:
//		fmt.Println("mild")
//	default:
//		fmt.Println("cold")
//	}
//
// Order matters: once a case matches, the rest are skipped, so put the
// highest range first.
//
// Task: make `grade` convert a score into a letter grade. From top to
// bottom: 90 or more -> "A", 80 or more -> "B", 70 or more -> "C",
// 60 or more -> "D", everything else -> "F" (the default already does
// that part).
package main

import "fmt"

// grade converts a score from 0 to 100 into a letter grade.
func grade(score int) string {
	switch {
	// TODO: add the four cases above default, from top to bottom:
	// score >= 90 -> "A", score >= 80 -> "B",
	// score >= 70 -> "C", score >= 60 -> "D".
	default:
		return "F"
	}
}

func main() {
	fmt.Println("grade(84) =", grade(84))
}

// I AM NOT DONE
