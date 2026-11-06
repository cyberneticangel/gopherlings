package main

import (
	"slices"
	"testing"
)

func TestIdentity(t *testing.T) {
	cases := []struct {
		n    int
		want [][]int
	}{
		{1, [][]int{{1}}},
		{2, [][]int{{1, 0}, {0, 1}}},
		{3, [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}},
		{4, [][]int{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}},
		{0, nil},
	}
	for _, c := range cases {
		got := identity(c.n)
		if len(got) != len(c.want) {
			t.Errorf("identity(%d) has %d rows, want %d", c.n, len(got), len(c.want))
			continue
		}
		for i := range c.want {
			if !slices.Equal(got[i], c.want[i]) {
				t.Errorf("identity(%d) row %d = %v, want %v", c.n, i, got[i], c.want[i])
			}
		}
	}
}
