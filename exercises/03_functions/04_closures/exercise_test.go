package main

import "testing"

func TestCounter(t *testing.T) {
	c := counter()
	for want := 1; want <= 3; want++ {
		if got := c(); got != want {
			t.Errorf("counter() call %d: got %d, want %d", want, got, want)
		}
	}
}

func TestCounterIndependent(t *testing.T) {
	first := counter()
	first()
	first()
	second := counter()
	if got := second(); got != 1 {
		t.Errorf("fresh counter() = %d, want 1", got)
	}
	if got := first(); got != 3 {
		t.Errorf("original counter() = %d, want 3", got)
	}
}
