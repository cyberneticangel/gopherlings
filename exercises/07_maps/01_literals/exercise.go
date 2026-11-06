// Maps store key/value pairs. The quickest way to create one is a map
// literal: write the type, then the pairs inside braces.
//
//	ages := map[string]int{"alice": 31, "bob": 27}
//
// You can also start empty with `make(map[string]int)` and add entries
// with `ages["carol"] = 44`.
//
// One surprise: reading a MISSING key does not panic and is not an
// error. It just returns the zero value of the value type (0 for int).
//
// Task: make `ages` return a map literal holding "alice": 31 and
// "bob": 27.
package main

import "fmt"

// ages returns the ages of a few gophers.
func ages() map[string]int {
	return nil // TODO: return a map literal with "alice": 31 and "bob": 27
}

func main() {
	fmt.Println(ages())
}

// I AM NOT DONE
