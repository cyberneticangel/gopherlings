package main

import "testing"

func TestAnnounce(t *testing.T) {
	cases := []struct {
		s    Speaker
		want string
	}{
		{Dog{}, "It says: Woof"},
		{Cat{}, "It says: Meow"},
	}
	for _, c := range cases {
		if got := announce(c.s); got != c.want {
			t.Errorf("announce(%T{}) = %q, want %q", c.s, got, c.want)
		}
	}
}
