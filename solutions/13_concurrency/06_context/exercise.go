// A context carries cancellation: ctx.Done() closes when the context
// is cancelled, and ctx.Err() says why. Waiting code selects on what
// it is waiting for AND on ctx.Done(), so it gives up as soon as the
// caller cancels.
package main

import (
	"context"
	"fmt"
)

// waitForReady blocks until ready is closed or ctx is cancelled. It
// returns nil in the first case and ctx.Err() in the second.
func waitForReady(ctx context.Context, ready <-chan struct{}) error {
	select {
	case <-ready:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ready := make(chan struct{})
	close(ready)
	fmt.Println("waitForReady:", waitForReady(context.Background(), ready))
}
