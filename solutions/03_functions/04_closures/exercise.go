// A closure is a function value that remembers the variables from the
// scope it was created in. Here the returned function keeps a private
// count that nobody else can touch.
package main

import "fmt"

// counter returns a function that returns 1, 2, 3, ... on successive
// calls. Each counter counts independently.
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	next := counter()
	fmt.Println(next(), next(), next())
}
