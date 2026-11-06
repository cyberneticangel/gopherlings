// A Go string is a read-only sequence of bytes, encoded in UTF-8.
// That means len(s) counts BYTES, not characters: the accented letter
// in the test file takes two bytes, so len reports 6 for that
// five-letter word.
//
// Ranging over a string decodes one rune per step (a rune is a
// Unicode code point), and utf8.RuneCountInString(s) counts them
// directly.
//
// Task: make `runeCount` count runes, not bytes. The test uses a word
// where the two disagree.
package main

import "fmt"

// runeCount returns the number of runes (characters) in s.
func runeCount(s string) int {
	return len(s) // TODO: len counts bytes. Count runes instead.
}

func main() {
	fmt.Println(runeCount("hello"))
	fmt.Println("len says 6 for the accented test word; runeCount should say 5.")
}

// I AM NOT DONE
