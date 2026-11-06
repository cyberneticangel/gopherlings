package main

import "testing"

func TestStackStrings(t *testing.T) {
	s := &Stack[string]{}
	s.Push("a")
	s.Push("b")

	if v, ok := s.Pop(); v != "b" || !ok {
		t.Errorf("first Pop() = (%q, %t), want (\"b\", true)", v, ok)
	}
	if v, ok := s.Pop(); v != "a" || !ok {
		t.Errorf("second Pop() = (%q, %t), want (\"a\", true)", v, ok)
	}
	if v, ok := s.Pop(); v != "" || ok {
		t.Errorf("Pop() on empty stack = (%q, %t), want (\"\", false)", v, ok)
	}
}

func TestStackInts(t *testing.T) {
	s := &Stack[int]{}
	for i := 1; i <= 3; i++ {
		s.Push(i * 10)
	}
	for want := 30; want >= 10; want -= 10 {
		if v, ok := s.Pop(); v != want || !ok {
			t.Errorf("Pop() = (%d, %t), want (%d, true)", v, ok, want)
		}
	}
	if v, ok := s.Pop(); v != 0 || ok {
		t.Errorf("Pop() on empty stack = (%d, %t), want (0, false)", v, ok)
	}
}
