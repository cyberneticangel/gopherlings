package main

import "testing"

func TestDaysInWeek(t *testing.T) {
	if got := daysInWeek(); got != 7 {
		t.Errorf("daysInWeek() = %d, want %d", got, 7)
	}
	if days != 7 {
		t.Errorf("days = %d, want %d", days, 7)
	}
}
