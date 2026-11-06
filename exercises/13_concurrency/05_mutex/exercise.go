// When several goroutines touch the same data, plain access is a data
// race. A sync.Mutex lets one goroutine at a time into the critical
// section:
//
//	mu.Lock()
//	// ... use the shared data ...
//	mu.Unlock()
//
// Maps are the classic victim: two concurrent writes to one map crash
// the program with "fatal error: concurrent map writes".
//
// Task: give SafeCounter a sync.Mutex field and use it in Increment
// and Get -- lock before touching the map, and unlock before
// returning (defer mu.Unlock() right after Lock is the idiomatic
// form).
package main

import "fmt"

// SafeCounter counts hits per key. Safe for concurrent use -- after
// your fix.
type SafeCounter struct {
	counts map[string]int // TODO: shared without a lock -- add a sync.Mutex field
}

// newSafeCounter returns an empty, ready-to-use counter.
func newSafeCounter() *SafeCounter {
	return &SafeCounter{counts: make(map[string]int)}
}

// Increment adds one to the count for key.
func (c *SafeCounter) Increment(key string) {
	c.counts[key]++ // TODO: guard this write with the mutex
}

// Get returns the current count for key.
func (c *SafeCounter) Get(key string) int {
	return c.counts[key] // TODO: guard this read too
}

func main() {
	c := newSafeCounter()
	c.Increment("hits")
	c.Increment("hits")
	fmt.Println("hits:", c.Get("hits"))
}

// I AM NOT DONE
