package main

import (
	"strings"
	"testing"
)

func TestDivide(t *testing.T) {
	got, err := divide(6, 3)
	if err != nil {
		t.Fatalf("divide(6, 3) returned error %v, want none", err)
	}
	if got != 2 {
		t.Errorf("divide(6, 3) = %v, want 2", got)
	}

	got, err = divide(1, 0)
	if err == nil {
		t.Fatal("divide(1, 0) returned no error, want one about dividing by zero")
	}
	if !strings.Contains(err.Error(), "divide by zero") {
		t.Errorf("divide(1, 0) error = %q, want it to mention \"divide by zero\"", err)
	}
	if got != 0 {
		t.Errorf("divide(1, 0) = %v, want 0 when there is an error", got)
	}
}
