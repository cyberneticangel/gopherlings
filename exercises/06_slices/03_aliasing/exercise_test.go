package main

import (
	"slices"
	"testing"
)

func TestCloneOf(t *testing.T) {
	orig := []int{1, 2, 3}
	res := cloneOf(orig)
	if !slices.Equal(res, orig) {
		t.Fatalf("cloneOf(%v) = %v, want a copy with the same elements", orig, res)
	}
	res[0] = 99
	if orig[0] != 1 {
		t.Errorf("writing res[0] = 99 changed the original: orig = %v, want [1 2 3]", orig)
	}
	orig[1] = 55
	if res[1] != 2 {
		t.Errorf("writing orig[1] = 55 changed the clone: res = %v, want [99 2 3]", res)
	}

	empty := []int{}
	clone := cloneOf(empty)
	clone = append(clone, 7)
	if len(empty) != 0 {
		t.Errorf("appending to the clone changed the original: %v, want []", empty)
	}
}
