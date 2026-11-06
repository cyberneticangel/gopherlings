// Repeated + concatenation copies the string every time (O(n^2));
// strings.Builder appends into one buffer (O(n)) and hands you the
// finished string at the end.
package main

import (
	"fmt"
	"strings"
)

// joinLines joins lines with a "\n" between them and no trailing newline.
func joinLines(lines []string) string {
	var b strings.Builder
	for i, line := range lines {
		if i > 0 {
			b.WriteString("\n")
		}
		b.WriteString(line)
	}
	return b.String()
}

func main() {
	fmt.Println(joinLines([]string{"first", "second", "third"}))
}
