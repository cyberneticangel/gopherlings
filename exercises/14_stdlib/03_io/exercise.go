// io.Reader is Go's universal input abstraction: one tiny interface — a
// Read method — shared by files, network connections, compressed streams,
// even in-memory strings (strings.NewReader). Helpers like io.ReadAll
// drain any Reader until EOF for you.
//
// Reading can fail, and Go's rule is simple: pass the error along.
// Swallowing it hides real problems from your caller.
//
// Task: make countBytes read everything from r (io.ReadAll works, or
// io.Copy(io.Discard, r)) and return how many bytes it read. If reading
// fails, return the error unchanged — the test includes a reader that
// always fails and expects its exact error back.
package main

import (
	"fmt"
	"io"
	"strings"
)

// countBytes reads everything from r and returns how many bytes it read.
// If reading fails, the error travels to the caller unchanged.
func countBytes(r io.Reader) (int, error) {
	return 0, nil // TODO: read all of r and return the byte count and the error
}

func main() {
	n, err := countBytes(strings.NewReader("gophers love io"))
	fmt.Printf("countBytes read %d bytes (err = %v)\n", n, err)
}

// I AM NOT DONE
