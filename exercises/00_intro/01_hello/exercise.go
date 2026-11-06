// Welcome to gopherlings! 🐹
//
// This project teaches Go through small, broken programs. Each exercise
// lives in its own directory and comes with a test that checks your work.
//
// Your workflow:
//
//  1. Run `gopherlings` in this directory — it verifies your work every
//     time you save. (`gopherlings run` verifies once instead.)
//  2. Fix the code until the tests pass.
//  3. Delete the `// I AM NOT DONE` comment at the bottom and save to
//     move on to the next exercise.
//
// This first exercise does not compile. Read the error the runner shows
// you, fix it, and watch the test go green. Go is telling you exactly
// what is wrong — learning to read compiler errors is half of learning Go.
package main

import "fmt"

// greeting builds the message gopherlings greets you with.
func greeting() string {
	message := "Hello, Gopher!"
	return Message // <- Go is case-sensitive: `Message` does not exist!
}

func main() {
	fmt.Println(greeting())
}

// I AM NOT DONE
