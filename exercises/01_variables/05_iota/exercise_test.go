package main

import "testing"

func TestWeekdayNumbers(t *testing.T) {
	cases := []struct {
		name string
		day  Weekday
		want Weekday
	}{
		{"Sunday", Sunday, 0},
		{"Monday", Monday, 1},
		{"Tuesday", Tuesday, 2},
		{"Wednesday", Wednesday, 3},
		{"Thursday", Thursday, 4},
		{"Friday", Friday, 5},
		{"Saturday", Saturday, 6},
	}
	for _, c := range cases {
		if c.day != c.want {
			t.Errorf("%s = %d, want %d", c.name, c.day, c.want)
		}
	}
}

func TestIsWeekend(t *testing.T) {
	cases := map[Weekday]bool{
		Sunday:    true,
		Monday:    false,
		Tuesday:   false,
		Wednesday: false,
		Thursday:  false,
		Friday:    false,
		Saturday:  true,
	}
	for day, want := range cases {
		if got := isWeekend(day); got != want {
			t.Errorf("isWeekend(%d) = %t, want %t", day, got, want)
		}
	}
}
