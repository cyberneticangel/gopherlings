package main

import "testing"

func TestIndexOf(t *testing.T) {
	if got := indexOf([]string{"a", "b", "c"}, "b"); got != 1 {
		t.Errorf(`indexOf([]string{"a", "b", "c"}, "b") = %d, want 1`, got)
	}
	if got := indexOf([]int{1, 2}, 9); got != -1 {
		t.Errorf("indexOf([]int{1, 2}, 9) = %d, want -1", got)
	}
	if got := indexOf([]int{5, 5, 5}, 5); got != 0 {
		t.Errorf("indexOf([]int{5, 5, 5}, 5) = %d, want 0 (the first match wins)", got)
	}
	if got := indexOf([]int{}, 1); got != -1 {
		t.Errorf("indexOf([]int{}, 1) = %d, want -1", got)
	}
}
