package main

import "testing"

func TestScale(t *testing.T) {
	r := Rect{W: 3, H: 4}
	r.Scale(2)
	if r.W != 6 || r.H != 8 {
		t.Errorf("after Scale(2), rect = %v, want {6 8}", r)
	}
	if got := r.Area(); got != 48 {
		t.Errorf("area after Scale(2) = %v, want 48", got)
	}

	p := Rect{W: 2.5, H: 2}
	p.Scale(3)
	if got := p.Area(); got != 45 {
		t.Errorf("area after Scale(3) = %v, want 45", got)
	}

	small := Rect{W: 1, H: 1}
	small.Scale(0.5)
	if small.W != 0.5 || small.H != 0.5 {
		t.Errorf("after Scale(0.5), rect = %v, want {0.5 0.5}", small)
	}
}
