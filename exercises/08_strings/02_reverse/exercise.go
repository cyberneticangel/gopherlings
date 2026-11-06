// Reversing a string looks easy: swap the bytes from both ends until
// they meet. And for plain ASCII it works! But multi-byte characters
// get shredded: the bytes of one character end up beside the wrong
// bytes, and the result is garbage.
//
// The fix: convert to []rune first, so you are swapping whole
// characters, then convert back with string(runes).
//
// Task: make `reverse` work for ANY string. The swapping loop is
// already here; it is just operating on the wrong type.
package main

import "fmt"

// reverse returns s with its runes in the opposite order.
func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b) // TODO: this reverses bytes and garbles multi-byte runes. Swap runes instead.
}

func main() {
	fmt.Println(reverse("go"))
	fmt.Println("The test reverses a word with a two-byte letter. Bytes will garble it.")
}

// I AM NOT DONE
