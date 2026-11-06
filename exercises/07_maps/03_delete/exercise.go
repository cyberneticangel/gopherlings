// Removing an entry from a map is a one-liner with the delete builtin:
//
//	delete(m, key)
//
// Deleting a key that is not there is a harmless no-op. Maps are
// reference types, so delete changes the map the caller sees too;
// that is why removeKey can mutate m and return the very same map.
//
// Task: make `removeKey` actually delete the key. Right now it returns
// the map untouched.
package main

import "fmt"

// removeKey deletes k from m and returns the map.
func removeKey(m map[string]int, k string) map[string]int {
	return m // TODO: delete k from m before returning it
}

func main() {
	ages := map[string]int{"alice": 31, "bob": 27}
	fmt.Println(removeKey(ages, "bob"))
}

// I AM NOT DONE
