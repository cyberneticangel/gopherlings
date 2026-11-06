// Map iteration order is randomized, so collect the keys and sort them
// when a stable order matters. sort.Strings sorts a []string in place.
package main

import (
	"fmt"
	"sort"
)

// sortedKeys returns every key of m, sorted ascending.
func sortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func main() {
	counts := map[string]int{"zebra": 3, "apple": 1, "mango": 2}
	fmt.Println("map:   ", counts)
	fmt.Println("sorted:", sortedKeys(counts))
}
