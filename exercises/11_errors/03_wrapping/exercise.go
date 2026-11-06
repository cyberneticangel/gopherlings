// "not found" -- found WHAT? Lower-level errors are often too vague
// on their own. Wrapping adds context while keeping the original
// error reachable:
//
//	fmt.Errorf("load %q: %w", key, err)
//
// The %w verb embeds the wrapped error in the message AND links it
// into a chain; errors.Is follows that chain, so identity checks
// keep working across layers. Using %v instead formats the same
// text but SEVERS the chain: the message looks right, yet errors.Is
// says false.
//
// Task: make `load` wrap the ErrNotFound sentinel with fmt.Errorf
// and %w, so the error mentions the key AND still satisfies
// errors.Is(err, ErrNotFound).
package main

import (
	"errors"
	"fmt"
)

// ErrNotFound is the sentinel error for a missing key.
var ErrNotFound = errors.New("not found")

// ages is the fixed data that load reads.
var ages = map[string]int{"ada": 36, "grace": 45}

// load looks a name up, wrapping failures with the key as context.
func load(key string) (int, error) {
	age, ok := ages[key]
	if !ok {
		return 0, fmt.Errorf("load %q: %v", key, ErrNotFound) // TODO: %v formats the text but loses the error -- wrap with %w instead
	}
	return age, nil
}

func main() {
	age, err := load("grace")
	if err != nil {
		fmt.Println("oops:", err)
		return
	}
	fmt.Println("grace is", age)

	if _, err := load("bob"); err != nil {
		fmt.Println("oops:", err)
	}
}

// I AM NOT DONE
