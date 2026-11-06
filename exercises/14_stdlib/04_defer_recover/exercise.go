// defer and recover are Go's safety net.
//
// A deferred function call runs just before the surrounding function
// exits — even if a panic is unwinding through it — and deferred calls
// run in LIFO order (last deferred, first run). recover() stops a panic,
// but ONLY when called directly inside a deferred function; anywhere else
// it just returns nil.
//
// Task: make catchPanic run f without letting a panic escape. Add a
// deferred closure that calls recover() and, when it returns non-nil,
// sets the named return value `panicked` to true.
package main

import "fmt"

// catchPanic runs f and reports whether it panicked.
func catchPanic(f func()) (panicked bool) {
	f() // TODO: a panic in f must not escape — guard it with defer + recover
	return panicked
}

func main() {
	bomb := func() { panic("boom") }
	fmt.Println("catchPanic(bomb) =", catchPanic(bomb))
	fmt.Println("catchPanic(quiet) =", catchPanic(func() {}))
}

// I AM NOT DONE
