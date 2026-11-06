package main

import "testing"

func TestSwap(t *testing.T) {
	x, y := 1, 2
	swap(&x, &y)
	if x != 2 || y != 1 {
		t.Errorf("after swap(&x, &y), x = %d and y = %d, want 2 and 1", x, y)
	}

	a, b := -7, 42
	swap(&a, &b)
	if a != 42 || b != -7 {
		t.Errorf("after swap(&a, &b), a = %d and b = %d, want 42 and -7", a, b)
	}

	n, m := 9, 9
	swap(&n, &m)
	if n != 9 || m != 9 {
		t.Errorf("after swapping two equal values, got %d and %d, want 9 and 9", n, m)
	}
}
