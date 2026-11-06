package main

import (
	"sync"
	"testing"
	"time"
)

// guardStuck is the safety net for a counter whose methods can block
// (say, a Lock without an Unlock). Without it, a stuck test would
// hang until the 10-minute testing timeout.
func guardStuck(t *testing.T, msg string) {
	t.Helper()
	timer := time.AfterFunc(5*time.Second, func() {
		panic("STUCK: " + msg)
	})
	t.Cleanup(func() { timer.Stop() })
}

func TestSafeCounterConcurrent(t *testing.T) {
	guardStuck(t, "a counter method blocked: every Lock needs an Unlock")

	const goroutines = 200
	counter := newSafeCounter()

	var wg sync.WaitGroup
	wg.Add(goroutines)
	for range goroutines {
		go func() {
			defer wg.Done()
			counter.Increment("hits")
		}()
	}
	wg.Wait()

	if got := counter.Get("hits"); got != goroutines {
		t.Errorf("Get(\"hits\") after %d concurrent Increments = %d, want %d (updates were lost)", goroutines, got, goroutines)
	}
}

func TestSafeCounterBasics(t *testing.T) {
	guardStuck(t, "a counter method blocked: every Lock needs an Unlock")

	counter := newSafeCounter()
	if got := counter.Get("nope"); got != 0 {
		t.Errorf("Get on a fresh counter = %d, want 0", got)
	}
	counter.Increment("hits")
	counter.Increment("hits")
	counter.Increment("misses")
	if got := counter.Get("hits"); got != 2 {
		t.Errorf("Get(\"hits\") = %d, want 2", got)
	}
	if got := counter.Get("misses"); got != 1 {
		t.Errorf("Get(\"misses\") = %d, want 1", got)
	}
}
