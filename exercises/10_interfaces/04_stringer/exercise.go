// fmt does not know how YOUR types should look, so %v falls back to
// printing the raw value: Temperature(21.5) shows up as just "21.5".
//
// The fmt.Stringer interface lets a type fix that:
//
//	type Stringer interface {
//		String() string
//	}
//
// Whenever fmt prints a value with %v (or Println, Sprint, ...), it
// checks whether the value has a String() method and calls it. Note
// what happened there: fmt was written long before your type, in a
// package that does not import yours -- implicit satisfaction at
// work.
//
// Task: give Temperature a String() method so it prints as a
// temperature: for 21.5, fmt.Sprintf("%.1f°C", float64(t)) gives
// "21.5°C". The receiver t holds a Temperature, so convert it to
// float64 before formatting.
package main

import "fmt"

// Temperature is a temperature in degrees Celsius.
type Temperature float64

func main() {
	fmt.Printf("room temperature: %v\n", Temperature(21.5))
}

// I AM NOT DONE
