package main

import "testing"

func TestZeroValues(t *testing.T) {
	gotN, gotS, gotB := zeroValues()
	if gotN != 0 {
		t.Errorf("zeroValues() int = %d, want 0", gotN)
	}
	if gotS != "" {
		t.Errorf("zeroValues() string = %q, want %q", gotS, "")
	}
	if gotB {
		t.Errorf("zeroValues() bool = %t, want false", gotB)
	}
}
