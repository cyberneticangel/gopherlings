// FizzBuzz — the classic capstone. It combines a for loop, an
// if/else if chain, and a slice that you build as you go.
//
// For every number i from 1 to n, one string goes into the result:
//
//   - divisible by 3 AND 5 -> "FizzBuzz"
//   - divisible by 3       -> "Fizz"
//   - divisible by 5       -> "Buzz"
//   - otherwise            -> the number as text, via strconv.Itoa(i)
//
// A slice is a growable list. Start it empty, then append to it — and
// remember, append returns the new slice, so keep the assignment:
//
//	result := make([]string, 0, n)
//	result = append(result, "one entry")
//
// Task: make `fizzbuzz` return all n entries, in order.
package main

import "fmt"

// fizzbuzz returns the FizzBuzz answers for the numbers 1 through n.
func fizzbuzz(n int) []string {
	return nil // TODO: build the list from "1" up to n, Fizzes and Buzzes included
}

func main() {
	fmt.Println(fizzbuzz(15))
}

// I AM NOT DONE
