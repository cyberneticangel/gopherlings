package main

import "testing"

func TestDouble(t *testing.T) {
	x := 21
	double(&x)
	if x != 42 {
		t.Errorf("after double(&x), x = %d, want 42", x)
	}

	y := -5
	double(&y)
	if y != -10 {
		t.Errorf("after double(&y), y = %d, want -10", y)
	}

	z := 0
	double(&z)
	if z != 0 {
		t.Errorf("after double(&z), z = %d, want 0", z)
	}
}
