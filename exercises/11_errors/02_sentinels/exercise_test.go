package main

import (
	"errors"
	"testing"
)

func TestFind(t *testing.T) {
	stock := map[string]int{"apple": 3, "pear": 1}

	got, err := find(stock, "apple")
	if err != nil {
		t.Fatalf("find(stock, \"apple\") returned error %v, want none", err)
	}
	if got != 3 {
		t.Errorf("find(stock, \"apple\") = %d, want 3", got)
	}

	if _, err = find(stock, "plum"); err == nil {
		t.Fatal("find(stock, \"plum\") returned no error, want ErrNotFound")
	}
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("find(stock, \"plum\") error = %q, but errors.Is(err, ErrNotFound) = false -- same text, wrong value", err)
	}
}
