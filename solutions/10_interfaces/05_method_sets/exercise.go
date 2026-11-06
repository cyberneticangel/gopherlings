// Pointer-receiver methods live in the method set of *T only, so
// only *Counter satisfies Incrementer. Value-receiver methods are in
// both method sets -- which is why Get can be called on the Counter
// value c directly, and on the pointer through the interface.
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
	var inc Incrementer = &c
	inc.Inc()
	inc.Inc()
	fmt.Println("counted to", c.Get())
}
