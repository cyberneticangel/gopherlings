// A context.Context carries cancellation across API boundaries. A
// function that waits for something should ALSO watch its context:
// ctx.Done() is a channel that closes when the context is cancelled,
// and ctx.Err() then reports why (context.Canceled, or
// context.DeadlineExceeded for a deadline).
//
// The standard shape is a select:
//
//	select {
//	case <-ready:
//		return nil       // what we waited for happened
//	case <-ctx.Done():
//		return ctx.Err() // cancelled first
//	}
//
// Task: make waitForReady follow that shape -- return nil once ready
// is closed, or ctx.Err() if the context is cancelled first. The
// current version ignores the context completely.
package main

import (
	"context"
	"fmt"
)

// waitForReady blocks until ready is closed or ctx is cancelled. It
// returns nil in the first case and ctx.Err() in the second.
func waitForReady(ctx context.Context, ready <-chan struct{}) error {
	<-ready // TODO: select on ready AND ctx.Done(); return ctx.Err() if the context wins
	return nil
}

func main() {
	ready := make(chan struct{})
	close(ready)
	fmt.Println("waitForReady:", waitForReady(context.Background(), ready))
}

// I AM NOT DONE
