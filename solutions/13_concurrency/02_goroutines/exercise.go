// `go f()` runs f in its own goroutine and returns immediately. The
// producer pattern: a goroutine sends its values, then closes the
// channel so receivers know the stream has ended.
package main

import "fmt"

// fillAsync sends every num on ch from its own goroutine, then closes
// ch.
func fillAsync(ch chan<- int, nums []int) {
	go func() {
		defer close(ch)
		for _, n := range nums {
			ch <- n
		}
	}()
}

func main() {
	ch := make(chan int, 3)
	fillAsync(ch, []int{1, 2, 3})
	for v := range ch {
		fmt.Println("got:", v)
	}
}
