package main

import "testing"

func TestSplit(t *testing.T) {
	cases := []struct {
		sum      int
		wantHalf int
		wantRest int
	}{
		{7, 3, 4},
		{10, 5, 5},
		{0, 0, 0},
		{1, 0, 1},
	}
	for _, c := range cases {
		gotHalf, gotRest := split(c.sum)
		if gotHalf != c.wantHalf || gotRest != c.wantRest {
			t.Errorf("split(%d) = (%d, %d), want (%d, %d)",
				c.sum, gotHalf, gotRest, c.wantHalf, c.wantRest)
		}
	}
}
