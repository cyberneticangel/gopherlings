package main

import "testing"

func TestSamePoint(t *testing.T) {
	cases := []struct {
		a, b Point
		want bool
	}{
		{Point{X: 1, Y: 2}, Point{X: 1, Y: 2}, true},
		{Point{X: 1, Y: 2}, Point{X: 2, Y: 1}, false},
		{Point{X: 0, Y: 0}, Point{X: 0, Y: 0}, true},
		{Point{X: -1, Y: 5}, Point{X: -1, Y: 5}, true},
		{Point{X: -1, Y: 5}, Point{X: -1, Y: -5}, false},
	}
	for _, c := range cases {
		if got := samePoint(c.a, c.b); got != c.want {
			t.Errorf("samePoint(%v, %v) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
