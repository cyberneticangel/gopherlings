package main

import (
	"errors"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	age, err := load("ada")
	if err != nil {
		t.Fatalf("load(\"ada\") returned error %v, want none", err)
	}
	if age != 36 {
		t.Errorf("load(\"ada\") = %d, want 36", age)
	}

	_, err = load("missing")
	if err == nil {
		t.Fatal("load(\"missing\") returned no error, want one wrapping ErrNotFound")
	}
	if !strings.Contains(err.Error(), "missing") {
		t.Errorf("load(\"missing\") error = %q, want it to mention the key", err)
	}
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("load(\"missing\") error = %q, but errors.Is(err, ErrNotFound) = false -- the chain got lost", err)
	}
}
