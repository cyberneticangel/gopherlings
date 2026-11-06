package main

import "testing"

func TestSumTo(t *testing.T) {
	cases := []struct {
		n, want int
	}{
		{5, 15},
		{0, 0},
		{1, 1},
		{100, 5050},
	}
	for _, c := range cases {
		if got := sumTo(c.n); got != c.want {
			t.Errorf("sumTo(%d) = %d, want %d", c.n, got, c.want)
		}
	}
}

func TestPowerOfTwo(t *testing.T) {
	cases := []struct {
		exp, want int
	}{
		{10, 1024},
		{0, 1},
		{1, 2},
		{8, 256},
	}
	for _, c := range cases {
		if got := powerOfTwo(c.exp); got != c.want {
			t.Errorf("powerOfTwo(%d) = %d, want %d", c.exp, got, c.want)
		}
	}
}
