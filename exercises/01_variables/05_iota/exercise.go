// Go has no enums, but it has something better: `iota`. Inside a const
// block, iota counts up from 0 on each line, and its expression is
// repeated for lines that omit one:
//
//	const (
//		First  = iota // 0
//		Second        // 1
//		Third         // 2
//	)
//
// Combined with a named type, this is Go's idiom for enumerations.
//
// Task: the weekday constants below were numbered by hand -- and they
// are off by one. Rewrite the const block with iota so that Sunday is 0
// and Saturday is 6 (no hand-written numbers). Then fix isWeekend so it
// reports true for Saturday OR Sunday -- use the || operator.
package main

import "fmt"

// Weekday counts the days of the week starting at Sunday = 0.
type Weekday int

const (
	Sunday    = 1 // TODO: replace these hand-written numbers with iota
	Monday    = 2
	Tuesday   = 3
	Wednesday = 4
	Thursday  = 5
	Friday    = 6
	Saturday  = 7
)

// isWeekend reports whether d is a weekend day (Saturday or Sunday).
func isWeekend(d Weekday) bool {
	return d == Sunday // TODO: Saturday is a weekend day too -- use ||
}

func main() {
	fmt.Printf("Wednesday is day %d of the week (0-indexed). Weekend? %t\n",
		Wednesday, isWeekend(Wednesday))
}

// I AM NOT DONE
