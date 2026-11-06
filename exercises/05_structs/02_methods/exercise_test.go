package main

import "testing"

func TestArea(t *testing.T) {
	cases := []struct {
		r    Rect
		want float64
	}{
		{Rect{W: 3, H: 4}, 12},
		{Rect{W: 2.5, H: 2}, 5},
		{Rect{W: 1, H: 1}, 1},
		{Rect{W: 0, H: 7}, 0},
	}
	for _, c := range cases {
		if got := c.r.Area(); got != c.want {
			t.Errorf("Rect%v.Area() = %v, want %v", c.r, got, c.want)
		}
	}
}
