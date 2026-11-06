package main

import "testing"

func TestDescribe(t *testing.T) {
	cases := []struct {
		v    any
		want string
	}{
		{42, "number"},
		{-7, "number"},
		{"hello", "text"},
		{"", "text"},
		{true, "flag"},
		{false, "flag"},
		{nil, "nothing"},
		{3.14, "unknown"},
		{[]int{1, 2}, "unknown"},
		{struct{}{}, "unknown"},
	}
	for _, c := range cases {
		if got := describe(c.v); got != c.want {
			t.Errorf("describe(%#v) = %q, want %q", c.v, got, c.want)
		}
	}
}
