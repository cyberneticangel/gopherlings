// encoding/json marshals structs to JSON: the `json:"..."` tags set the
// keys, fields keep their struct order, and json.Marshal hands back the
// bytes. Check the error, then convert with string(data).
package main

import (
	"encoding/json"
	"fmt"
)

// User is a struct whose JSON shape is decided by its tags.
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// encode returns u as a JSON string, or "" if marshalling fails.
func encode(u User) string {
	data, err := json.Marshal(u)
	if err != nil {
		return ""
	}
	return string(data)
}

func main() {
	fmt.Println("Ada as JSON:", encode(User{Name: "Ada", Age: 36}))
}
