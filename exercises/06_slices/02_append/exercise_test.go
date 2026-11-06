package main

import (
	"slices"
	"testing"
)

func TestPush(t *testing.T) {
	cases := []struct {
		xs   []int
		v    int
		want []int
	}{
		{[]int{1, 2}, 3, []int{1, 2, 3}},
		{nil, 1, []int{1}},
		{[]int{}, 42, []int{42}},
		{[]int{7, 8, 9}, 10, []int{7, 8, 9, 10}},
	}
	for _, c := range cases {
		got := push(c.xs, c.v)
		if !slices.Equal(got, c.want) {
			t.Errorf("push(%v, %d) = %v, want %v", c.xs, c.v, got, c.want)
		}
	}
}
