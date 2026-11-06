// The encoding/json package turns Go values into JSON and back.
//
// For structs, the `json:"..."` tag on each field decides the JSON key;
// fields without a tag keep their Go name. Fields appear in the output in
// the same order they are declared in the struct. json.Marshal returns
// the JSON as a []byte, ready to be converted with string(data).
//
// Task: make encode marshal u with json.Marshal (you will need to add the
// encoding/json import). Handle the error properly — ignoring it with _
// is NOT ok: return "" if err != nil, otherwise string(data). The test
// expects exactly {"name":"Ada","age":36} for Ada, 36.
package main

import "fmt"

// User is a struct whose JSON shape is decided by its tags.
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// encode returns u as a JSON string, or "" if marshalling fails.
func encode(u User) string {
	return "" // TODO: marshal u, then return "" on error or string(data)
}

func main() {
	fmt.Println("Ada as JSON:", encode(User{Name: "Ada", Age: 36}))
}

// I AM NOT DONE
