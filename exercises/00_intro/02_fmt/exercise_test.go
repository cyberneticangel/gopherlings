package main

import "testing"

func TestWelcome(t *testing.T) {
	cases := map[string]string{
		"Gopher": "Welcome to gopherlings, Gopher!",
		"Ada":    "Welcome to gopherlings, Ada!",
		"world":  "Welcome to gopherlings, world!",
	}
	for name, want := range cases {
		if got := welcome(name); got != want {
			t.Errorf("welcome(%q) = %q, want %q", name, got, want)
		}
	}
}
