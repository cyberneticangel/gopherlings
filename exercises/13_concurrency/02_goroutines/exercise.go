// The `go` keyword starts a GOROUTINE: a lightweight, independently
// running function. `go f()` returns immediately -- f keeps going in
// the background while your code continues.
//
// A favorite pattern: hand a channel to a producer goroutine that
// sends its values and then closes the channel to say "that is all".
//
// Task: make fillAsync start ONE goroutine (go func() { ... }()) that
// sends every num on ch and then closes ch. The current code sends
// INLINE -- and the test does not start receiving until fillAsync has
// returned, so with the test's unbuffered channel an inline send
// blocks forever: there is nobody to receive yet.
package main

import "fmt"

// fillAsync sends every num on ch from its own goroutine, then closes
// ch.
func fillAsync(ch chan<- int, nums []int) {
	// TODO: move this loop into a goroutine: go func() { ... }()
	for _, n := range nums {
		ch <- n
	}
	close(ch)
}

func main() {
	ch := make(chan int, 3)
	fillAsync(ch, []int{1, 2, 3})
	for v := range ch {
		fmt.Println("got:", v)
	}
}

// I AM NOT DONE
