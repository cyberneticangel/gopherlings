// Ranging over a map visits every key:
//
//	for k, v := range m { ... }
//
// But here is the catch: the order is RANDOMIZED. Go deliberately
// shuffles it (even between two runs of the same program) so nobody
// accidentally depends on it. When order matters, collect the keys and
// sort them yourself, with sort.Strings or the one-liner
// slices.Sorted(maps.Keys(m)).
//
// Task: make `sortedKeys` return all keys of m in ascending order.
package main

import "fmt"

// sortedKeys returns every key of m, sorted ascending.
func sortedKeys(m map[string]int) []string {
	return nil // TODO: collect the keys in a range loop, then sort them
}

func main() {
	counts := map[string]int{"zebra": 3, "apple": 1, "mango": 2}
	fmt.Println("map:   ", counts)
	fmt.Println("sorted:", sortedKeys(counts))
}

// I AM NOT DONE
