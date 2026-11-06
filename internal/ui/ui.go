// Package ui provides tiny ANSI-color helpers with no external dependencies.
// Colors are disabled when stdout is not a terminal or NO_COLOR is set.
package ui

import (
	"fmt"
	"os"
	"strings"
)

var enabled = true

// Init configures color support based on the environment. Call once.
func Init() {
	if os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb" {
		enabled = false
		return
	}
	fi, err := os.Stdout.Stat()
	if err != nil || fi.Mode()&os.ModeCharDevice == 0 {
		enabled = false
	}
}

const (
	ansiReset  = "\x1b[0m"
	ansiBold   = "\x1b[1m"
	ansiDim    = "\x1b[2m"
	ansiRed    = "\x1b[31m"
	ansiGreen  = "\x1b[32m"
	ansiYellow = "\x1b[33m"
	ansiBlue   = "\x1b[34m"
	ansiCyan   = "\x1b[36m"
)

func wrap(color, s string) string {
	if !enabled || s == "" {
		return s
	}
	return color + s + ansiReset
}

// Bold renders s in bold.
func Bold(s string) string { return wrap(ansiBold, s) }

// Dim renders s dimmed.
func Dim(s string) string { return wrap(ansiDim, s) }

// Red renders s in red.
func Red(s string) string { return wrap(ansiRed, s) }

// Green renders s in green.
func Green(s string) string { return wrap(ansiGreen, s) }

// Yellow renders s in yellow.
func Yellow(s string) string { return wrap(ansiYellow, s) }

// Blue renders s in blue.
func Blue(s string) string { return wrap(ansiBlue, s) }

// Cyan renders s in cyan.
func Cyan(s string) string { return wrap(ansiCyan, s) }

// Boldf renders formatted s in bold.
func Boldf(format string, a ...any) string { return Bold(fmt.Sprintf(format, a...)) }

// Dimf renders formatted s dimmed.
func Dimf(format string, a ...any) string { return Dim(fmt.Sprintf(format, a...)) }

// Redf renders formatted s in red.
func Redf(format string, a ...any) string { return Red(fmt.Sprintf(format, a...)) }

// Greenf renders formatted s in green.
func Greenf(format string, a ...any) string { return Green(fmt.Sprintf(format, a...)) }

// Yellowf renders formatted s in yellow.
func Yellowf(format string, a ...any) string { return Yellow(fmt.Sprintf(format, a...)) }

// Indent prefixes every line of s with pad.
func Indent(s, pad string) string {
	lines := strings.Split(strings.TrimRight(s, "\n"), "\n")
	for i, line := range lines {
		lines[i] = pad + line
	}
	return strings.Join(lines, "\n")
}
