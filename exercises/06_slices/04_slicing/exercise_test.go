package main

import (
	"slices"
	"testing"
)

func TestMiddle(t *testing.T) {
	cases := []struct {
		xs, want []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4}},
		{[]int{1, 2, 3}, []int{2}},
		{[]int{4, 9}, nil},
		{[]int{1}, nil},
		{nil, nil},
	}
	for _, c := range cases {
		got := middle(c.xs)
		if !slices.Equal(got, c.want) {
			t.Errorf("middle(%v) = %v, want %v", c.xs, got, c.want)
		}
	}
}
