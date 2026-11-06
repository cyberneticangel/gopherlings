// A generic type carries its type parameter everywhere the type is
// used: Stack[T any] stores a []T, and its methods restate [T] in the
// receiver so they can take and return T. One implementation now
// serves stacks of strings, ints, or anything else.
package main

import "fmt"

// Stack is a last-in, first-out stack of Ts.
type Stack[T any] struct {
	items []T
}

// Push puts v on top of the stack.
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

// Pop removes and returns the top value. ok is false when the stack
// is empty; the value is then the zero value of T.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	last := len(s.items) - 1
	v := s.items[last]
	s.items = s.items[:last]
	return v, true
}

func main() {
	s := &Stack[int]{}
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	words := &Stack[string]{}
	words.Push("gopher")
	fmt.Println(words.Pop())
}
