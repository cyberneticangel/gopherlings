package main

import "testing"

func TestMaxOf(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{3, 1, 4, 1, 5}, 5},
		{[]int{7}, 7},
		{[]int{-3, -1}, -1},
		{[]int{-2, -8, -5}, -2},
	}
	for _, c := range cases {
		if got := maxOf(c.nums...); got != c.want {
			t.Errorf("maxOf(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}
