package main

import "testing"

func TestCounterThroughInterface(t *testing.T) {
	c := &Counter{}
	var inc Incrementer = c
	inc.Inc()
	inc.Inc()
	if got := c.Get(); got != 2 {
		t.Errorf("after two inc.Inc() calls, c.Get() = %d, want 2", got)
	}
}
