// len counts bytes; utf8.RuneCountInString counts runes. A range loop
// over a string also steps one rune at a time.
package main

import (
	"fmt"
	"unicode/utf8"
)

// runeCount returns the number of runes (characters) in s.
func runeCount(s string) int {
	return utf8.RuneCountInString(s)
}

func main() {
	fmt.Println(runeCount("hello"))
	fmt.Println("len says 6 for the accented test word; runeCount should say 5.")
}
