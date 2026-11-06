// gopherlings is a Rustlings/Ziglings-style trainer for Go: a set of small,
// deliberately broken exercises plus a CLI that verifies your fixes and
// walks you through the curriculum.
package main

import (
	"embed"
	"os"

	"github.com/cyberneticangel/gopherlings/internal/cmd"
)

// content holds the entire curriculum (exercises + solutions), embedded at
// build time so the binary is fully self-contained.
//
//go:embed all:exercises all:solutions
var content embed.FS

func main() {
	os.Exit(cmd.Main(content, os.Args[1:]))
}
