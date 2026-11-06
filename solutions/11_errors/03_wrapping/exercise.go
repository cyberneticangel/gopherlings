// The %w verb wraps an error: the message gains context and the
// original error stays linked in a chain. errors.Is unwraps layer by
// layer, so sentinel checks survive any amount of wrapping. %v would
// print the same text -- and throw the chain away.
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
		return 0, fmt.Errorf("load %q: %w", key, ErrNotFound)
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
