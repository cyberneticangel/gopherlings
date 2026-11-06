// Types can be generic too: declare a type parameter after the type
// name and use it like any other type inside.
//
//	type Stack[T any] struct {
//		items []T
//	}
//
// Methods on a generic type restate the parameter in their receiver
// -- func (s *Stack[T]) Push(v T) -- which is what lets them use T.
// `any` places no restriction: the right constraint for a container
// that only stores values.
//
// Task: generalize IntStack into Stack[T any]. Rename the type, add
// the [T any] parameter, and make Push take a T and Pop return a T.
// Pop keeps its contract: on an empty stack it returns the zero value
// of T with ok == false (var zero T gives you that zero value).
// Update main to match.
package main

import "fmt"

// TODO: generalize this into Stack[T any]: rename the type, add the
// [T any] parameter, and swap int for T in Push and Pop.

// IntStack is a last-in, first-out stack of ints.
type IntStack struct {
	items []int
}

// Push puts v on top of the stack.
func (s *IntStack) Push(v int) {
	s.items = append(s.items, v)
}

// Pop removes and returns the top value. ok is false when the stack
// is empty; the value is then the zero value.
func (s *IntStack) Pop() (int, bool) {
	if len(s.items) == 0 {
		var zero int
		return zero, false
	}
	last := len(s.items) - 1
	v := s.items[last]
	s.items = s.items[:last]
	return v, true
}

func main() {
	s := &IntStack{}
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

// I AM NOT DONE
