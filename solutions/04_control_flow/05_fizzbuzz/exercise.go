// FizzBuzz combines a for loop, an if/else if chain and a slice built
// with append.
package main

import (
	"fmt"
	"strconv"
)

// fizzbuzz returns the FizzBuzz answers for the numbers 1 through n.
func fizzbuzz(n int) []string {
	result := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

func main() {
	fmt.Println(fizzbuzz(15))
}
