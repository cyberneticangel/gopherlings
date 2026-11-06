package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"racecar", true},
		{"", true},
		{"ab", false},
		{"aba", true},
		{"abba", true},
		{"level", true},
		{"Level", false},
		// a few extras for good measure
		{"a", true},
		{"abcba", true},
		{"gopher", false},
		{"Hello", false},
	}
	for _, c := range cases {
		if got := isPalindrome(c.in); got != c.want {
			t.Errorf("isPalindrome(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}
