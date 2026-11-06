// The delete builtin removes a key from a map. Deleting an absent key
// is a no-op: delete(m, key).
package main

import "fmt"

// removeKey deletes k from m and returns the map.
func removeKey(m map[string]int, k string) map[string]int {
	delete(m, k)
	return m
}

func main() {
	ages := map[string]int{"alice": 31, "bob": 27}
	fmt.Println(removeKey(ages, "bob"))
}
