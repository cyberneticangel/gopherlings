package main

import "testing"

func TestDeclarations(t *testing.T) {
	if got := sectionName(); got != "variables" {
		t.Errorf("sectionName() = %q, want %q", got, "variables")
	}
	if got := gopherCount(); got != 5 {
		t.Errorf("gopherCount() = %d, want %d", got, 5)
	}
	if !isReady() {
		t.Errorf("isReady() = false, want true")
	}
	if got := piApprox(); got != 3.14 {
		t.Errorf("piApprox() = %v, want %v", got, 3.14)
	}
}
