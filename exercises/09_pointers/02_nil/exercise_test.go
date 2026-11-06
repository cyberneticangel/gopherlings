package main

import "testing"

func TestSafeDeref(t *testing.T) {
	name := "Gopher"
	if got := safeDeref(&name); got != "Gopher" {
		t.Errorf("safeDeref(&name) = %q, want %q", got, "Gopher")
	}

	empty := ""
	if got := safeDeref(&empty); got != "" {
		t.Errorf("safeDeref(&empty) = %q, want an empty string", got)
	}

	if got := safeDeref(nil); got != "" {
		t.Errorf("safeDeref(nil) = %q, want an empty string", got)
	}
}
