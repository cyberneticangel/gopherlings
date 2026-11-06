// A type switch inspects an interface value's dynamic type in one
// statement: case int makes the int branch run, case nil matches a
// nil value, and default catches everything else. Bind a variable,
// switch x := v.(type), when you also need the value -- x then holds
// v converted to the matched case's type.
package main

import "fmt"

// describe names the kind of value v holds.
func describe(v any) string {
	switch v.(type) {
	case int:
		return "number"
	case string:
		return "text"
	case bool:
		return "flag"
	case nil:
		return "nothing"
	default:
		return "unknown"
	}
}

func main() {
	for _, v := range []any{42, "hi", true, nil, 3.14} {
		fmt.Println(describe(v))
	}
}
