// A sentinel is one shared error value, returned by name. errors.Is
// checks identity, so it only matches when the function returns the
// sentinel itself -- a fresh errors.New with the same message fails
// the check.
package main

import (
	"errors"
	"fmt"
)

// ErrNotFound is the sentinel error for a missing key.
var ErrNotFound = errors.New("not found")

// find looks k up in m, returning ErrNotFound when it is absent.
func find(m map[string]int, k string) (int, error) {
	v, ok := m[k]
	if !ok {
		return 0, ErrNotFound
	}
	return v, nil
}

func main() {
	stock := map[string]int{"apple": 3}

	n, err := find(stock, "apple")
	if err != nil {
		fmt.Println("oops:", err)
		return
	}
	fmt.Println("apples:", n)

	if _, err := find(stock, "plum"); err != nil {
		fmt.Println("plums:", err)
	}
}
