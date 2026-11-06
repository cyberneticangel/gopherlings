package main

import "testing"

func TestSign(t *testing.T) {
	cases := []struct {
		x    int
		want string
	}{
		{5, "positive"},
		{1, "positive"},
		{0, "zero"},
		{-3, "negative"},
		{-100, "negative"},
	}
	for _, c := range cases {
		if got := sign(c.x); got != c.want {
			t.Errorf("sign(%d) = %q, want %q", c.x, got, c.want)
		}
	}
}
