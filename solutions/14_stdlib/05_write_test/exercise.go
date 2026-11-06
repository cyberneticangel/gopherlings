// The flipped exercise, solved: isPalindrome was correct all along — the
// real work was replacing the TODO with a table-driven test in
// exercise_test.go.
package main

import "fmt"

// isPalindrome reports whether s reads the same forwards and backwards.
//
// Contract: runes are compared case-sensitively, so "Level" is not a
// palindrome, and the empty string counts as one.
func isPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func main() {
	for _, w := range []string{"racecar", "level", "Level", "ab"} {
		fmt.Printf("%-9q palindrome? %v\n", w, isPalindrome(w))
	}
}
