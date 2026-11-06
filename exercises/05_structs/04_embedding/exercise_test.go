package main

import "testing"

func TestDescribe(t *testing.T) {
	cases := []struct {
		car  Car
		want string
	}{
		{Car{Make: "Zephyr", Engine: Engine{Horsepower: 210}}, "Zephyr with 210 HP"},
		{Car{Make: "Pinto", Engine: Engine{Horsepower: 88}}, "Pinto with 88 HP"},
		{Car{Make: "Rustbucket", Engine: Engine{Horsepower: 45}}, "Rustbucket with 45 HP"},
	}
	for _, c := range cases {
		if got := c.car.Describe(); got != c.want {
			t.Errorf("Car{Make: %q, ...}.Describe() = %q, want %q", c.car.Make, got, c.want)
		}
	}
}
