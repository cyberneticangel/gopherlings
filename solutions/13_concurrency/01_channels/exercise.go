// Channels connect goroutines: a send hands a value to a receive, and
// both sides block until they meet. close(ch) announces "no more
// values" -- once what was sent has been received, further receives
// return the zero value with ok == false.
package main

import "fmt"

// send passes n to a receiver on ch, then announces that no more
// values are coming.
func send(n int, ch chan int) {
	ch <- n
	close(ch)
}

func main() {
	ch := make(chan int, 1)
	send(42, ch)
	v, ok := <-ch
	fmt.Println("received:", v, "open:", ok)
	v, ok = <-ch
	fmt.Println("received:", v, "open:", ok)
}
