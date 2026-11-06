// A constant is a name for a value that is fixed when the program is
// compiled. Declare them with the `const` keyword:
//
//	const days = 7
//
// Two rules the compiler enforces:
//
//   - constants have no `:=` -- `const days := 7` is a syntax error,
//   - constants can never be reassigned.
//
// Because a constant's value is known at compile time, you can even
// build new constants out of existing ones.
//
// Task: fix the broken declaration below, then make daysInWeek return
// the constant instead of reassigning it. TODO comments mark the spots.
package main

import "fmt"

const days := 7 // TODO: constants do not use :=

const slogan string = "gophers never quit" // an explicit type is fine

// days * 24 is computed at compile time -- this only works because
// days is a constant.
const hoursInWeek = days * 24

// daysInWeek reports how many days a week has.
func daysInWeek() int {
	days = 8 // TODO: constants cannot be reassigned -- delete this line
	return days
}

func main() {
	fmt.Printf("a week has %d days = %d hours. %s!\n",
		daysInWeek(), hoursInWeek, slogan)
}

// I AM NOT DONE
