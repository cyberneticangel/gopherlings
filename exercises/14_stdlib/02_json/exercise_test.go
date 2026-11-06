package main

import "testing"

func TestEncode(t *testing.T) {
	cases := []struct {
		user User
		want string
	}{
		{User{Name: "Ada", Age: 36}, `{"name":"Ada","age":36}`},
		{User{}, `{"name":"","age":0}`},
	}
	for _, c := range cases {
		if got := encode(c.user); got != c.want {
			t.Errorf("encode(%+v) = %s, want %s", c.user, got, c.want)
		}
	}
}
