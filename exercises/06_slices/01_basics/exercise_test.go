package main

import "testing"

func TestSecond(t *testing.T) {
	cases := []struct {
		xs   []int
		want int
	}{
		{[]int{10, 20, 30}, 20},
		{[]int{5, 7}, 7},
		{[]int{-1, -2, -3}, -2},
	}
	for _, c := range cases {
		if got := second(c.xs); got != c.want {
			t.Errorf("second(%v) = %d, want %d", c.xs, got, c.want)
		}
	}
}

func TestTotal(t *testing.T) {
	cases := []struct {
		xs   []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{}, 0},
		{nil, 0},
		{[]int{42}, 42},
		{[]int{-5, 5, 10}, 10},
	}
	for _, c := range cases {
		if got := total(c.xs); got != c.want {
			t.Errorf("total(%v) = %d, want %d", c.xs, got, c.want)
		}
	}
}
