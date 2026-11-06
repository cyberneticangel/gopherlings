package main

import "testing"

func TestMax(t *testing.T) {
	check := func(name string, got, want any) {
		t.Helper()
		if got != want {
			t.Errorf("%s: got %v, want %v", name, got, want)
		}
	}

	check("Max(3, 9)", Max(3, 9), 9)
	check("Max(9, 3)", Max(9, 3), 9)
	check("Max(1.5, 2.5)", Max(1.5, 2.5), 2.5)
	check("Max(2.5, 1.5)", Max(2.5, 1.5), 2.5)
	check(`Max("a", "b")`, Max("a", "b"), "b")
	check(`Max("b", "a")`, Max("b", "a"), "b")
}
