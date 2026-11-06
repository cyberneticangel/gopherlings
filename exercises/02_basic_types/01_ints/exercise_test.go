package main

import "testing"

func TestAverage(t *testing.T) {
	cases := []struct {
		a, b int
		want float64
	}{
		{1, 2, 1.5},
		{3, 4, 3.5},
		{2, 4, 3},
		{-3, 2, -0.5},
		{0, 0, 0},
	}
	for _, c := range cases {
		if got := average(c.a, c.b); got != c.want {
			t.Errorf("average(%d, %d) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
