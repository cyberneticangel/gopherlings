package main

import (
	"testing"
	"time"
)

// guardStuck is the safety net for a parallelSum that deadlocks
// internally (say, goroutines blocked sending on a full channel).
// Without it, a stuck test would hang until the 10-minute testing
// timeout.
func guardStuck(t *testing.T, msg string) {
	t.Helper()
	timer := time.AfterFunc(5*time.Second, func() {
		panic("STUCK: " + msg)
	})
	t.Cleanup(func() { timer.Stop() })
}

func TestParallelSum(t *testing.T) {
	guardStuck(t, "parallelSum blocked: every partial sum needs room on the shared channel, and wg.Wait must come after the fan-out")

	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i + 1
	}

	cases := []struct {
		name  string
		nums  []int
		parts int
		want  int
	}{
		{"1..100 in 4 parts", nums, 4, 5050},
		{"1..100 in 1 part", nums, 1, 5050},
		{"1..10 in 3 parts", nums[:10], 3, 55},
		{"one element", []int{7}, 1, 7},
		{"empty slice", nil, 3, 0},
	}
	for _, c := range cases {
		if got := parallelSum(c.nums, c.parts); got != c.want {
			t.Errorf("parallelSum(%s) = %d, want %d", c.name, got, c.want)
		}
	}
}
