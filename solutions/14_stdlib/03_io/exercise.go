// io.Reader is Go's universal input abstraction. io.ReadAll drains any
// Reader until EOF; whatever error surfaces must be passed along to the
// caller untouched.
package main

import (
	"fmt"
	"io"
	"strings"
)

// countBytes reads everything from r and returns how many bytes it read.
// If reading fails, the error travels to the caller unchanged.
func countBytes(r io.Reader) (int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	return len(data), nil
}

func main() {
	n, err := countBytes(strings.NewReader("gophers love io"))
	fmt.Printf("countBytes read %d bytes (err = %v)\n", n, err)
}
