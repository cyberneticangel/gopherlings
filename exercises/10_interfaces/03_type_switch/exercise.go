// A type assertion, v.(int), pulls the value out if v holds an int.
// Checking one type at a time gets clumsy fast, so Go has a switch
// for it:
//
//	switch x := v.(type) {
//	case int:
//		// x is an int here, equal to v
//	case string:
//		// x is a string here
//	}
//
// The variable x holds v converted to the matched case's type. Bind
// it only when you need the value -- an unused binding is a compile
// error, so a plain switch v.(type) is enough when you only need the
// kind. A `case nil:` matches a nil v, and default catches the rest.
//
// Task: make `describe` return "number" for ints, "text" for
// strings, "flag" for bools, "nothing" for nil, and "unknown" for
// anything else.
package main

import "fmt"

// describe names the kind of value v holds.
func describe(v any) string {
	return "" // TODO: inspect v's dynamic type with a type switch
}

func main() {
	for _, v := range []any{42, "hi", true, nil, 3.14} {
		fmt.Println(describe(v))
	}
}

// I AM NOT DONE
