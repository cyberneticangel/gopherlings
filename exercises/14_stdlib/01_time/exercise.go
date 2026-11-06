// Times and durations live in the time package.
//
// A time.Duration counts nanoseconds, and the constants time.Second and
// time.Minute are Durations too — so arithmetic just works:
// time.Duration(s) * time.Second is the Duration for s seconds. Every
// Duration also knows how to format itself: the String method prints
// compactly, e.g. 90 seconds as "1m30s" and 3661 as "1h1m1s".
//
// Task: make secondsLabel(s) build the Duration for s seconds and return
// its .String().
package main

import (
	"fmt"
	"time"
)

// secondsLabel formats a number of seconds as a compact duration string,
// like 90 -> "1m30s".
func secondsLabel(s int) string {
	return "" // TODO: build the Duration and return its .String()
}

func main() {
	fmt.Println("a Duration prints itself:", time.Hour)
	fmt.Println("90 seconds ->", secondsLabel(90))
	fmt.Println("3661 seconds ->", secondsLabel(3661))
}

// I AM NOT DONE
