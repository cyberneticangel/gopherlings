package main

import (
	"testing"
	"time"
)

// guardStuck is the safety net for blocks the receive guards cannot
// reach (the test calls send directly). Without it, a stuck test
// would hang until the 10-minute testing timeout.
func guardStuck(t *testing.T, msg string) {
	t.Helper()
	timer := time.AfterFunc(5*time.Second, func() {
		panic("STUCK: " + msg)
	})
	t.Cleanup(func() { timer.Stop() })
}

// receiveInt waits up to 2s for one value from ch, so a missing send
// is reported as a timeout instead of hanging the test.
func receiveInt(t *testing.T, ch <-chan int) (int, bool) {
	t.Helper()
	select {
	case v, ok := <-ch:
		return v, ok
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting to receive from ch (2s) -- did send do its job?")
	}
	return 0, false
}

func TestSend(t *testing.T) {
	guardStuck(t, "send blocked: on this buffered channel it must return immediately")

	ch := make(chan int, 1)
	send(42, ch)

	v, ok := receiveInt(t, ch)
	if v != 42 || !ok {
		t.Fatalf("first receive from ch = (%d, %t), want (42, true)", v, ok)
	}

	v, ok = receiveInt(t, ch)
	if v != 0 || ok {
		t.Errorf("second receive from ch = (%d, %t), want the channel closed (0, false)", v, ok)
	}
}
