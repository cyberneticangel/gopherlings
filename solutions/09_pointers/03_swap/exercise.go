// Go is call-by-value, so a function that must change its caller's
// variables takes pointers. With *int parameters and multiple
// assignment, the exchange is one line: *a, *b = *b, *a.
package main

import "fmt"

// swap exchanges the values pointed at by a and b.
func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println("after swap:", x, y)
}
