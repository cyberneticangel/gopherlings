// A type's method set decides which interfaces it satisfies. A
// value-receiver method belongs to both T and *T, but a
// pointer-receiver method belongs to *T ONLY. So this Counter has
// Get in both method sets, while Inc lives only on *Counter.
//
// Go is not being pedantic here: not every value has an address you
// can take (a map entry, for one), so a value cannot honestly
// promise a method that mutates through a pointer.
//
// Storing a plain Counter{} in an Incrementer therefore refuses to
// compile, with the error pointing right at the problem:
//
//	Counter does not implement Incrementer (method Inc has
//	pointer receiver)
//
// Task: fix the marked line by storing the ADDRESS of a counter in
// the interface -- &c works.
package main

import "fmt"

// Counter counts up.
type Counter struct {
	n int
}

// Inc bumps the count; it mutates, so it takes a pointer receiver.
func (c *Counter) Inc() {
	c.n++
}

// Get reports the count; it only reads, so a value receiver is fine.
func (c Counter) Get() int {
	return c.n
}

// Incrementer is satisfied by anything you can call Inc on.
type Incrementer interface {
	Inc()
}

func main() {
	c := Counter{}
	var inc Incrementer = c // TODO: Counter does not implement Incrementer -- Inc has a pointer receiver
	inc.Inc()
	inc.Inc()
	fmt.Println("counted to", c.Get())
}

// I AM NOT DONE
