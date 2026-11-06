// select waits on several channel operations and runs whichever is
// ready. The comma-ok form detects close, and assigning nil to a
// channel variable disables its case -- the standard way to merge
// channels until every input is exhausted.
package main

import "fmt"

// fanIn merges two string channels into one. The returned channel is
// closed once both inputs are closed and drained.
func fanIn(a, b <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				out <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				out <- v
			}
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
