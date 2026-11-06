package main

import "testing"

func TestSecondsLabel(t *testing.T) {
	cases := []struct {
		seconds int
		want    string
	}{
		{90, "1m30s"},
		{59, "59s"},
		{3661, "1h1m1s"},
		{0, "0s"},
		{60, "1m0s"},
	}
	for _, c := range cases {
		if got := secondsLabel(c.seconds); got != c.want {
			t.Errorf("secondsLabel(%d) = %q, want %q", c.seconds, got, c.want)
		}
	}
}
