package main

import "testing"

func TestGrade(t *testing.T) {
	cases := []struct {
		score int
		want  string
	}{
		{95, "A"},
		{90, "A"},
		{84, "B"},
		{73, "C"},
		{61, "D"},
		{60, "D"},
		{12, "F"},
		{0, "F"},
	}
	for _, c := range cases {
		if got := grade(c.score); got != c.want {
			t.Errorf("grade(%d) = %q, want %q", c.score, got, c.want)
		}
	}
}
