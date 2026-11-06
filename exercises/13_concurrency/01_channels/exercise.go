// Channels connect goroutines: a send hands a value to a receive, and
// both sides block until they meet -- that synchronization is the
// point. `ch <- v` sends, `v, ok := <-ch` receives.
//
// close(ch) announces "no more values". Once everything sent has been
// received, further receives return the zero value with ok == false.
//
// Task: make send put n on the channel and then close it, so the
// receiver gets 42 first and then sees ok == false.
//
// Tip: run this file before fixing it and Go will stop it with the
// famous "fatal error: all goroutines are asleep - deadlock!" -- the
// runtime noticing that nobody can make progress.
package main

import "fmt"

// send passes n to a receiver on ch, then announces that no more
// values are coming.
func send(n int, ch chan int) {
	_ = ch // TODO: send n on ch, then close ch
}

func main() {
	ch := make(chan int, 1)
	send(42, ch)
	v, ok := <-ch
	fmt.Println("received:", v, "open:", ok)
	v, ok = <-ch
	fmt.Println("received:", v, "open:", ok)
}

// I AM NOT DONE
