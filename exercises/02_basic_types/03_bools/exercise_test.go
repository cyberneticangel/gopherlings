package main

import "testing"

func TestInRange(t *testing.T) {
	cases := []struct {
		x, lo, hi int
		want      bool
	}{
		{0, 1, 10, false},   // below the range
		{1, 1, 10, true},    // on the low boundary
		{5, 1, 10, true},    // inside
		{10, 1, 10, true},   // on the high boundary
		{11, 1, 10, false},  // above the range
		{-5, -10, -1, true}, // negative ranges work too
	}
	for _, c := range cases {
		if got := inRange(c.x, c.lo, c.hi); got != c.want {
			t.Errorf("inRange(%d, %d, %d) = %t, want %t",
				c.x, c.lo, c.hi, got, c.want)
		}
	}
}
