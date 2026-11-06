package main

import "testing"

func TestDoubled(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"21", "42"},
		{"-5", "-10"},
		{"0", "0"},
		{"abc", "0"},
		{"", "0"},
		{"1000000", "2000000"},
	}
	for _, c := range cases {
		if got := doubled(c.in); got != c.want {
			t.Errorf("doubled(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
