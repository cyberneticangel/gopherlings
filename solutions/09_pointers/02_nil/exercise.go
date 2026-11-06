// A pointer's zero value is nil: pointing at nothing. Dereferencing
// nil panics at runtime, so check for nil before following a
// pointer.
package main

import "fmt"

// safeDeref returns the string p points at, or "" when p is nil.
func safeDeref(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func main() {
	name := "Gopher"
	fmt.Println(safeDeref(&name))
	fmt.Println(safeDeref(nil))
}
