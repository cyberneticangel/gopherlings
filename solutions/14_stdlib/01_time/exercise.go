// time.Duration counts nanoseconds, and time.Duration(s) * time.Second is
// the Duration for s seconds. Duration.String formats it compactly:
// "1m30s", "1h1m1s".
package main

import (
	"fmt"
	"time"
)

// secondsLabel formats a number of seconds as a compact duration string,
// like 90 -> "1m30s".
func secondsLabel(s int) string {
	d := time.Duration(s) * time.Second
	return d.String()
}

func main() {
	fmt.Println("a Duration prints itself:", time.Hour)
	fmt.Println("90 seconds ->", secondsLabel(90))
	fmt.Println("3661 seconds ->", secondsLabel(3661))
}
