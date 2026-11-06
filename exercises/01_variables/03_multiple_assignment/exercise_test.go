package main

import "testing"

func TestSwap(t *testing.T) {
	cases := []struct {
		a, b, wantA, wantB int
	}{
		{1, 2, 2, 1},
		{-5, 7, 7, -5},
		{0, 0, 0, 0},
		{100, -100, -100, 100},
	}
	for _, c := range cases {
		gotA, gotB := swap(c.a, c.b)
		if gotA != c.wantA || gotB != c.wantB {
			t.Errorf("swap(%d, %d) = (%d, %d), want (%d, %d)",
				c.a, c.b, gotA, gotB, c.wantA, c.wantB)
		}
	}
}
