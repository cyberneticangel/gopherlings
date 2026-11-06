// A tagless switch — no expression after the keyword — is a tidier
// if/else chain: each case is a full condition, tested top to bottom,
// and the first true one wins.
package main

import "fmt"

// grade converts a score from 0 to 100 into a letter grade.
func grade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func main() {
	fmt.Println("grade(84) =", grade(84))
}
