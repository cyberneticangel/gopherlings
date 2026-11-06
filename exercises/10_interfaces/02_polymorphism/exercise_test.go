package main

import (
	"slices"
	"testing"
)

func TestSpeakAll(t *testing.T) {
	cases := []struct {
		speakers []Speaker
		want     []string
	}{
		{[]Speaker{Dog{}, Cat{}}, []string{"Woof", "Meow"}},
		{[]Speaker{Dog{}, Cat{}, Robot{}}, []string{"Woof", "Meow", "Beep"}},
		{[]Speaker{Robot{}, Dog{}}, []string{"Beep", "Woof"}},
	}
	for _, c := range cases {
		if got := speakAll(c.speakers); !slices.Equal(got, c.want) {
			t.Errorf("speakAll(%v) = %v, want %v", c.speakers, got, c.want)
		}
	}
}
