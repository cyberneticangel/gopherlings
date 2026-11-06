// iota counts up from 0 inside a const block, one step per line. Give
// the first constant the expression `iota` and the rest repeat it --
// Go's idiom for numbered enumerations.
package main

import "fmt"

// Weekday counts the days of the week starting at Sunday = 0.
type Weekday int

const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

// isWeekend reports whether d is a weekend day (Saturday or Sunday).
func isWeekend(d Weekday) bool {
	return d == Saturday || d == Sunday
}

func main() {
	fmt.Printf("Wednesday is day %d of the week (0-indexed). Weekend? %t\n",
		Wednesday, isWeekend(Wednesday))
}
