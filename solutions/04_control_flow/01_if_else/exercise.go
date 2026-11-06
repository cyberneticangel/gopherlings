// The if statement runs a block only when its condition is true, and
// `else if` / `else` chain further cases behind it.
package main

import "fmt"

// sign classifies x as "positive", "negative" or "zero".
func sign(x int) string {
	if x > 0 {
		return "positive"
	} else if x < 0 {
		return "negative"
	} else {
		return "zero"
	}
}

func main() {
	fmt.Println("sign(-2) =", sign(-2))
}
