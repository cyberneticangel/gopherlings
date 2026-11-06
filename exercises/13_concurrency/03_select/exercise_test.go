package main

import (
	"testing"
	"time"
)

// guardStuck is the safety net for blocks the receive guards cannot
// reach (the test calls fanIn directly). Without it, a stuck test
// would hang until the 10-minute testing timeout.
func guardStuck(t *testing.T, msg string) {
	t.Helper()
	timer := time.AfterFunc(5*time.Second, func() {
		panic("STUCK: " + msg)
	})
	t.Cleanup(func() { timer.Stop() })
}

// collect reads from out until it closes, guarding every receive with
// a 2s timeout so a stuck fanIn is reported, not hung.
func collect(t *testing.T, out <-chan string) map[string]bool {
	t.Helper()
	got := map[string]bool{}
	for {
		select {
		case v, ok := <-out:
			if !ok {
				return got
			}
			got[v] = true
		case <-time.After(2 * time.Second):
			t.Fatalf("timed out collecting from fanIn; got %v so far", got)
		}
	}
}

func TestFanInMergesBoth(t *testing.T) {
	guardStuck(t, "fanIn blocked: it must return a channel right away and do the merging in a goroutine")

	a := make(chan string, 1)
	b := make(chan string, 1)
	a <- "a1"
	b <- "b1"
	close(a)
	close(b)

	got := collect(t, fanIn(a, b))
	if len(got) != 2 || !got["a1"] || !got["b1"] {
		t.Errorf("fanIn forwarded %v, want exactly \"a1\" and \"b1\"", got)
	}
}

func TestFanInDrainsBoth(t *testing.T) {
	guardStuck(t, "fanIn blocked: it must return a channel right away and do the merging in a goroutine")

	a := make(chan string, 2)
	b := make(chan string, 2)
	a <- "a1"
	a <- "a2"
	b <- "b1"
	b <- "b2"
	close(a)
	close(b)

	got := collect(t, fanIn(a, b))
	for _, want := range []string{"a1", "a2", "b1", "b2"} {
		if !got[want] {
			t.Errorf("fanIn output is missing %q (got %v)", want, got)
		}
	}
	if len(got) != 4 {
		t.Errorf("fanIn forwarded %d distinct values (%v), want 4", len(got), got)
	}
}
