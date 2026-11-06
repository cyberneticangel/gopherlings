// defer runs a call when the function exits, in LIFO order; recover()
// stops a panic when called directly inside a deferred function.
// Together they turn a panic into an ordinary boolean.
package main

import "fmt"

// catchPanic runs f and reports whether it panicked.
func catchPanic(f func()) (panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	f()
	return panicked
}

func main() {
	bomb := func() { panic("boom") }
	fmt.Println("catchPanic(bomb) =", catchPanic(bomb))
	fmt.Println("catchPanic(quiet) =", catchPanic(func() {}))
}
