// A sentinel error is one well-known, package-level error value that
// callers can check for by identity:
//
//	var ErrNotFound = errors.New("not found")
//
//	if errors.Is(err, ErrNotFound) { ... }
//
// errors.Is compares values, not text -- and that is the trap this
// exercise springs. Returning errors.New("not found") inside the
// function mints a FRESH error on every call: same message,
// different value, so errors.Is says false.
//
// The ErrNotFound sentinel is already declared below. Task: make
// `find` return the SENTINEL itself for missing keys, so callers can
// rely on errors.Is.
package main

import (
	"errors"
	"fmt"
)

// ErrNotFound is the sentinel error for a missing key.
var ErrNotFound = errors.New("not found")

// find looks k up in m, returning an error when it is absent.
func find(m map[string]int, k string) (int, error) {
	v, ok := m[k]
	if !ok {
		return 0, errors.New("not found") // TODO: a fresh error every call -- return the ErrNotFound sentinel instead
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

// I AM NOT DONE
