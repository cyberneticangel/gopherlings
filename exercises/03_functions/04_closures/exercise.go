// A closure is a function value that remembers the variables from the
// scope it was created in — even after that scope has finished.
//
// `counter` declares a local variable and returns a function literal
// that uses it. Every call to the returned function keeps seeing (and
// changing) that same variable, and each call of `counter` creates a
// fresh, independent one.
//
// Task: make the returned function count 1, 2, 3, ... on successive
// calls. Capture a local variable in `counter` and increment it.
package main

import "fmt"

// counter returns a function that returns 1, 2, 3, ... on successive
// calls. Each counter counts independently.
func counter() func() int {
	return func() int {
		return 1 // TODO: count up instead of always returning 1
	}
}

func main() {
	next := counter()
	fmt.Println(next(), next(), next())
}

// I AM NOT DONE
