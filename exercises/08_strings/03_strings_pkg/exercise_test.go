package main

import "testing"

func TestSlugify(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"  Hello World  ", "hello-world"},
		{"Go IS fun", "go-is-fun"},
		{"already clean", "already-clean"},
		{"  TRIM me  ", "trim-me"},
	}
	for _, c := range cases {
		if got := slugify(c.in); got != c.want {
			t.Errorf("slugify(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
