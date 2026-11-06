// strconv converts between strings and numbers: strconv.Atoi parses a
// string into an int, strconv.Itoa formats an int as a string. Atoi's
// error result is ignored with _ for now -- errors get their own
// section later.
package main

import (
	"fmt"
	"strconv"
)

// doubled parses s as an integer and returns double the value, as a
// string. Unparseable input counts as 0.
func doubled(s string) string {
	n, _ := strconv.Atoi(s)
	return strconv.Itoa(n * 2)
}

func main() {
	fmt.Println("doubled(\"21\") =", doubled("21"))
	fmt.Println("doubled(\"abc\") =", doubled("abc"))
}
