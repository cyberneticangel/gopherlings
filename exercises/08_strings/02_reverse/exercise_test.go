package main

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"go", "og"},
		{"héllo", "olléh"},
		{"abcd", "dcba"},
		{"x", "x"},
		{"", ""},
	}
	for _, c := range cases {
		if got := reverse(c.in); got != c.want {
			t.Errorf("reverse(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
