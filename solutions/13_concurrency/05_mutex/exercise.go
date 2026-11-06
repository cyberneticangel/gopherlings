// A sync.Mutex lets one goroutine at a time into the critical
// section: Lock before touching shared data, Unlock when done --
// defer mu.Unlock() right after Lock is the idiomatic form. Without
// it, concurrent writes to a map crash the program.
package main

import (
	"fmt"
	"sync"
)

// SafeCounter counts hits per key. Safe for concurrent use.
type SafeCounter struct {
	mu     sync.Mutex
	counts map[string]int
}

// newSafeCounter returns an empty, ready-to-use counter.
func newSafeCounter() *SafeCounter {
	return &SafeCounter{counts: make(map[string]int)}
}

// Increment adds one to the count for key.
func (c *SafeCounter) Increment(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[key]++
}

// Get returns the current count for key.
func (c *SafeCounter) Get(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counts[key]
}

func main() {
	c := newSafeCounter()
	c.Increment("hits")
	c.Increment("hits")
	fmt.Println("hits:", c.Get("hits"))
}
