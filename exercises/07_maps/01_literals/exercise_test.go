package main

import "testing"

func TestAges(t *testing.T) {
	got := ages()
	want := map[string]int{"alice": 31, "bob": 27}

	if len(got) != len(want) {
		t.Fatalf("ages() = %v, want %v", got, want)
	}
	for name, wantAge := range want {
		if age, ok := got[name]; !ok || age != wantAge {
			t.Errorf("ages()[%q] = %d, want %d", name, age, wantAge)
		}
	}

	// Reading a missing key is not an error: it gives the zero value.
	if age := got["nobody"]; age != 0 {
		t.Errorf(`ages()["nobody"] = %d, want 0 (missing keys return the zero value)`, age)
	}
}
