// The strings package is full of small helpers that save you from
// writing loops by hand. Three you will use all the time:
//
//	strings.TrimSpace(s)             strip leading/trailing whitespace
//	strings.ToLower(s)               lowercase
//	strings.ReplaceAll(s, old, new)  replace every occurrence
//
// They chain nicely: the output of one is the input of the next.
//
// Task: make `slugify` turn a string into a slug: trim it, lowercase
// it, and turn every space into a hyphen. Right now it does nothing.
package main

import "fmt"

// slugify returns s trimmed, lowercased, with spaces turned into hyphens.
func slugify(s string) string {
	return s // TODO: chain strings.TrimSpace, strings.ToLower and strings.ReplaceAll
}

func main() {
	fmt.Println(slugify("  Hello World  "))
}

// I AM NOT DONE
