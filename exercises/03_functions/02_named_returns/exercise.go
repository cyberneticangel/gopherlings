// Return values can be named. A named result is declared right in the
// function signature and starts out as its zero value (0 for int).
//
// Inside the function you assign to the names as ordinary variables.
// A `return` with nothing after it — a "naked return" — then sends the
// current values of the named results back to the caller.
//
// Task: `half` is already assigned. Give `rest` a value too, so the
// naked return sends back both parts: half = sum / 2 and
// rest = sum - half.
package main

import "fmt"

// split breaks a total into an even half and what is left over.
func split(sum int) (half int, rest int) {
	half = sum / 2
	// TODO: assign the leftover amount to `rest`
	return
}

func main() {
	half, rest := split(7)
	fmt.Printf("splitting 7: half %d, rest %d\n", half, rest)
}

// I AM NOT DONE
