package main

import "testing"

func TestLookup(t *testing.T) {
	cases := []struct {
		m    map[string]int
		k    string
		want int
		ok   bool
	}{
		{map[string]int{"alice": 31, "bob": 27}, "alice", 31, true},
		{map[string]int{"alice": 31, "bob": 27}, "bob", 27, true},
		{map[string]int{"alice": 31, "bob": 27}, "carol", 0, false},
		{map[string]int{"carol": 0}, "carol", 0, true},
		{map[string]int{}, "alice", 0, false},
	}
	for _, c := range cases {
		got, ok := lookup(c.m, c.k)
		if got != c.want || ok != c.ok {
			t.Errorf("lookup(%v, %q) = (%d, %t), want (%d, %t)", c.m, c.k, got, ok, c.want, c.ok)
		}
	}
}
