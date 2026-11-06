package main

import "testing"

func TestFizzBuzz15(t *testing.T) {
	want := []string{
		"1", "2", "Fizz", "4", "Buzz",
		"Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz",
	}
	got := fizzbuzz(15)
	if len(got) != len(want) {
		t.Fatalf("fizzbuzz(15) returned %d entries, want %d: %v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("fizzbuzz(15)[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestFizzBuzzShorter(t *testing.T) {
	cases := []struct {
		n    int
		want []string
	}{
		{1, []string{"1"}},
		{3, []string{"1", "2", "Fizz"}},
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
	}
	for _, c := range cases {
		got := fizzbuzz(c.n)
		if len(got) != len(c.want) {
			t.Errorf("fizzbuzz(%d) = %v, want %v", c.n, got, c.want)
			continue
		}
		for i := range c.want {
			if got[i] != c.want[i] {
				t.Errorf("fizzbuzz(%d)[%d] = %q, want %q", c.n, i, got[i], c.want[i])
			}
		}
	}
}
