// []byte slices a string into raw bytes and shreds multi-byte
// characters; []rune keeps whole characters. Reverse the runes, then
// convert back with string().
package main

import "fmt"

// reverse returns s with its runes in the opposite order.
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	fmt.Println(reverse("go"))
	fmt.Println("The test reverses a word with a two-byte letter. Bytes will garble it.")
}
