package main

import (
	"slices"
	"testing"
	"time"
)

func TestFillAsync(t *testing.T) {
	ch := make(chan int) // unbuffered: every send must meet a receiver

	// fillAsync itself must return immediately: the sending belongs to
	// a goroutine it starts. We deliberately do not receive from ch
	// until fillAsync has returned, so an inline send has nobody to
	// hand a value to and shows up here as a timeout.
	returned := make(chan struct{})
	go func() {
		fillAsync(ch, []int{1, 2, 3})
		close(returned)
	}()

	select {
	case <-returned:
	case <-time.After(2 * time.Second):
		t.Fatal("fillAsync did not return within 2s: it sends inline, which blocks on the unbuffered channel -- move the sends into a goroutine")
	}

	got := []int{}
	for range 3 {
		select {
		case v, ok := <-ch:
			if !ok {
				t.Fatalf("ch closed after %d values, want 3 values before the close", len(got))
			}
			got = append(got, v)
		case <-time.After(2 * time.Second):
			t.Fatalf("timed out with %d of 3 values received: %v -- is the goroutine sending on ch?", len(got), got)
		}
	}
	if !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("received %v, want [1 2 3]", got)
	}

	select {
	case v, ok := <-ch:
		if ok {
			t.Errorf("receive after the 3 values = (%d, true), want the channel closed", v)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for ch to be closed -- did the goroutine close it?")
	}
}
