package main

import "testing"

func TestSum(t *testing.T) {
	if got := sum([]int{1, 2, 3}); got != 6 {
		t.Errorf("sum([]int{1, 2, 3}) = %d, want 6", got)
	}
	if got := sum([]int{10, 20, 30, 40}); got != 100 {
		t.Errorf("sum([]int{10, 20, 30, 40}) = %d, want 100", got)
	}
	if got := sum([]int{}); got != 0 {
		t.Errorf("sum([]int{}) = %d, want 0", got)
	}
	if got := sum([]float64{1.5, 2.5}); got != 4 {
		t.Errorf("sum([]float64{1.5, 2.5}) = %v, want 4", got)
	}
}
