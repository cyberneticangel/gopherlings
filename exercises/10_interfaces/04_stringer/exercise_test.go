package main

import (
	"fmt"
	"testing"
)

func TestStringer(t *testing.T) {
	cases := []struct {
		temp Temperature
		want string
	}{
		{21.5, "21.5°C"},
		{0, "0.0°C"},
		{-40, "-40.0°C"},
		{100, "100.0°C"},
	}
	for _, c := range cases {
		if got := fmt.Sprintf("%v", c.temp); got != c.want {
			t.Errorf("Temperature(%v) printed as %q, want %q", float64(c.temp), got, c.want)
		}
	}
}
