package main

import "testing"

func TestJoinLines(t *testing.T) {
	cases := []struct {
		lines []string
		want  string
	}{
		{[]string{"a", "b"}, "a\nb"},
		{[]string{"first", "second", "third"}, "first\nsecond\nthird"},
		{[]string{"only"}, "only"},
		{[]string{}, ""},
		{nil, ""},
	}
	for _, c := range cases {
		if got := joinLines(c.lines); got != c.want {
			t.Errorf("joinLines(%q) = %q, want %q", c.lines, got, c.want)
		}
	}
}
