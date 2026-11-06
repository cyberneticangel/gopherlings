package main

import "testing"

func TestDivMod(t *testing.T) {
	cases := []struct {
		a, b         int
		wantQ, wantR int
	}{
		{17, 5, 3, 2},
		{8, 2, 4, 0},
		{3, 7, 0, 3},
		{9, 3, 3, 0},
	}
	for _, c := range cases {
		gotQ, gotR := divMod(c.a, c.b)
		if gotQ != c.wantQ || gotR != c.wantR {
			t.Errorf("divMod(%d, %d) = (%d, %d), want (%d, %d)",
				c.a, c.b, gotQ, gotR, c.wantQ, c.wantR)
		}
	}
}
