// strings.TrimSpace trims the edges, strings.ToLower lowercases and
// strings.ReplaceAll swaps every space for a hyphen. Chain them.
package main

import (
	"fmt"
	"strings"
)

// slugify returns s trimmed, lowercased, with spaces turned into hyphens.
func slugify(s string) string {
	return strings.ReplaceAll(strings.ToLower(strings.TrimSpace(s)), " ", "-")
}

func main() {
	fmt.Println(slugify("  Hello World  "))
}
