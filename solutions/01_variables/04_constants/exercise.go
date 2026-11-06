// Constants are fixed at compile time: declared with `const` (never
// `:=`) and never reassigned. Compile-time values can build other
// constants, like hoursInWeek below.
package main

import "fmt"

const days = 7

const slogan string = "gophers never quit"

const hoursInWeek = days * 24

// daysInWeek reports how many days a week has.
func daysInWeek() int {
	return days
}

func main() {
	fmt.Printf("a week has %d days = %d hours. %s!\n",
		daysInWeek(), hoursInWeek, slogan)
}
