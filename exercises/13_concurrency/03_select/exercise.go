// select waits on several channel operations at once and runs the
// first one that is ready:
//
//	select {
//	case v := <-a:
//		fmt.Println("from a:", v)
//	case v := <-b:
//		fmt.Println("from b:", v)
//	}
//
// Receiving as `v, ok := <-ch` also detects close: when ch is closed
// and drained, the case fires with ok == false. Setting the channel
// VARIABLE ITSELF to nil then disables that case -- a nil channel is
// never ready.
//
// Task: make fanIn forward values from BOTH a and b onto out, in
// whatever order they arrive, and close out once both inputs are
// closed and drained. The current version drains only a and drops
// everything from b.
package main

import "fmt"

// fanIn merges two string channels into one. The returned channel is
// closed once both inputs are closed and drained.
func fanIn(a, b <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for v := range a { // TODO: b never gets forwarded -- rewrite with a select over both
			out <- v
		}
	}()
	return out
}

func main() {
	a := make(chan string, 1)
	b := make(chan string, 1)
	a <- "from a"
	b <- "from b"
	close(a)
	close(b)

	for v := range fanIn(a, b) {
		fmt.Println(v)
	}
}

// I AM NOT DONE
