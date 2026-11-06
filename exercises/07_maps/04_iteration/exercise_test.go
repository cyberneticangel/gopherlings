package main

import "testing"

func TestSortedKeys(t *testing.T) {
	cases := []struct {
		m    map[string]int
		want []string
	}{
		{map[string]int{"bob": 27, "alice": 31}, []string{"alice", "bob"}},
		{map[string]int{"zoe": 3, "ada": 1, "mike": 2}, []string{"ada", "mike", "zoe"}},
		{map[string]int{}, nil},
	}
	for _, c := range cases {
		got := sortedKeys(c.m)
		if len(got) != len(c.want) {
			t.Errorf("sortedKeys(%v) = %v, want %v", c.m, got, c.want)
			continue
		}
		for i := range c.want {
			if got[i] != c.want[i] {
				t.Errorf("sortedKeys(%v) = %v, want %v", c.m, got, c.want)
				break
			}
		}
	}
}
