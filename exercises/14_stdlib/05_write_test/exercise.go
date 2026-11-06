// Flip the script: this time the CODE is correct and the TEST is broken.
//
// isPalindrome below is finished and needs no fixing. Your job is in
// exercise_test.go: replace the TODO test with a real table-driven test
// that covers at least these cases:
//
//	"racecar" -> true        "aba" -> true         "abba" -> true
//	"level"   -> true        ""    -> true         "ab"   -> false
//	"Level"   -> false       (the check is case-sensitive!)
//
// Write the test, run it, and once it passes delete the marker comment at
// the bottom of this file to finish the section.
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

// I AM NOT DONE
