// The comma-ok idiom: v, ok := m[k]. ok is true only when the key was
// present, even if the stored value is the zero value.
package main

import "fmt"

// lookup returns the value for k and whether k was present in m.
func lookup(m map[string]int, k string) (int, bool) {
	v, ok := m[k]
	return v, ok
}

func main() {
	ages := map[string]int{"alice": 31, "bob": 27}
	fmt.Println(lookup(ages, "alice"))
	fmt.Println(lookup(ages, "carol"))
}
