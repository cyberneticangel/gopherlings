package main

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestWaitForReadyRespectsCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancelled before waitForReady even runs

	ready := make(chan struct{}) // never closed
	done := make(chan error, 1)
	go func() {
		done <- waitForReady(ctx, ready)
	}()

	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Errorf("waitForReady with a cancelled context = %v, want %v", err, context.Canceled)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("did not respect context: waitForReady is still waiting on a cancelled ctx")
	}
}

func TestWaitForReadyReturnsOnReady(t *testing.T) {
	ready := make(chan struct{})
	close(ready)

	done := make(chan error, 1)
	go func() {
		done <- waitForReady(context.Background(), ready)
	}()

	select {
	case err := <-done:
		if err != nil {
			t.Errorf("waitForReady with a closed ready channel = %v, want nil", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for waitForReady to notice the closed ready channel")
	}
}
