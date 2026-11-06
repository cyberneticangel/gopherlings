package main

import "testing"

func TestNearlyEqual(t *testing.T) {
	cases := []struct {
		a, b float64
		want bool
	}{
		{0.1 + 0.2, 0.3, true},
		{0.1 + 0.2, 0.4, false},
		{1.0, 2.0, false},
		{3.14159, 3.14159, true},
		{1.0, 1.0 + 1e-10, true},
		{1.0, 1.0 + 1e-6, false},
	}
	for _, c := range cases {
		if got := nearlyEqual(c.a, c.b); got != c.want {
			t.Errorf("nearlyEqual(%v, %v) = %t, want %t", c.a, c.b, got, c.want)
		}
	}
}
