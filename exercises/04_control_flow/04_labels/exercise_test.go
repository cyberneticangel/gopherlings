package main

import "testing"

func TestFirstNegative(t *testing.T) {
	cases := []struct {
		name string
		grid [][]int
		want int
	}{
		{"negative in the second row", [][]int{{4, 8}, {1, -7, 9}}, -7},
		{"negative in the last spot", [][]int{{1, 2}, {3, 4}, {5, -1}}, -1},
		{"first wins, not the smallest", [][]int{{-2, -9}}, -2},
		{"no negatives", [][]int{{1, 2}, {3, 4}}, 0},
		{"empty grid", nil, 0},
	}
	for _, c := range cases {
		if got := firstNegative(c.grid); got != c.want {
			t.Errorf("firstNegative(%v) = %d, want %d", c.grid, got, c.want)
		}
	}
}
