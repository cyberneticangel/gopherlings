package main

import "testing"

func TestRuneCount(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"go", 2},
		{"hello", 5},
		{"héllo", 5}, // é is TWO bytes in UTF-8, so len says 6
		{"", 0},
	}
	for _, c := range cases {
		if got := runeCount(c.in); got != c.want {
			t.Errorf("runeCount(%q) = %d, want %d (len(%q) = %d bytes)", c.in, got, c.want, c.in, len(c.in))
		}
	}
}
