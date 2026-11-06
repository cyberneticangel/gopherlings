// Reading m[k] on a missing key returns the zero value of the value
// type (0 here) with no error and no panic. That makes it impossible
// to tell "the key is absent" from "the key is present with value 0".
//
// The comma-ok idiom fixes that:
//
//	v, ok := m[k] // ok is true only if the key was present
//
// Task: make `lookup` report whether the key really exists, using the
// comma-ok idiom. Right now it always claims the key exists.
package main

import "fmt"

// lookup returns the value for k and whether k was present in m.
func lookup(m map[string]int, k string) (int, bool) {
	return m[k], true // TODO: this always says the key exists. Use the comma-ok idiom.
}

func main() {
	ages := map[string]int{"alice": 31, "bob": 27}
	fmt.Println(lookup(ages, "alice"))
	fmt.Println(lookup(ages, "carol"))
}

// I AM NOT DONE
