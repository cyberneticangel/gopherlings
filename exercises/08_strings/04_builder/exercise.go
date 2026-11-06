// Strings are immutable: every s += x copies the whole string, so
// building one with repeated + gets slower and slower, O(n^2) for n
// pieces. strings.Builder accumulates everything in a single buffer
// and hands you the result once, at the end: O(n).
//
//	var b strings.Builder
//	b.WriteString("hello")
//	b.WriteString(" world")
//	s := b.String()
//
// Task: make `joinLines` join the lines with "\n" between them (no
// trailing newline) using strings.Builder. Empty or nil input gives "".
package main

import "fmt"

// joinLines joins lines with a "\n" between them and no trailing newline.
func joinLines(lines []string) string {
	return "" // TODO: build the result with a strings.Builder
}

func main() {
	fmt.Println(joinLines([]string{"first", "second", "third"}))
}

// I AM NOT DONE
