// Map literals pair keys with values: map[string]int{"alice": 31, ...}.
// Reading a missing key returns the zero value of the value type.
package main

import "fmt"

// ages returns the ages of a few gophers.
func ages() map[string]int {
	return map[string]int{
		"alice": 31,
		"bob":   27,
	}
}

func main() {
	fmt.Println(ages())
}
