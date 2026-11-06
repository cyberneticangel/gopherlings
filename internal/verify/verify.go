// Package verify runs `go test -json` against a single exercise package and
// turns the output into a structured result. Both styles of exercise funnel
// through the same command: packages that don't compile report build errors,
// packages that compile but misbehave report failing tests.
package verify

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

// DefaultTimeout caps a single verification run. Concurrency exercises can
// deadlock, so a timeout with a friendly message matters.
const DefaultTimeout = 45 * time.Second

// CompileError is a single compiler diagnostic.
type CompileError struct {
	File string
	Line string
	Col  string
	Msg  string
}

func (e CompileError) String() string {
	return fmt.Sprintf("%s:%s:%s: %s", e.File, e.Line, e.Col, e.Msg)
}

// TestFailure is one failing test and its captured output.
type TestFailure struct {
	Test   string
	Output string
}

// Result is the outcome of verifying one exercise.
type Result struct {
	OK bool

	Package string

	// BuildErrors are parsed compiler diagnostics, if the package didn't compile.
	BuildErrors []CompileError
	// OtherBuildOutput holds build-output lines we couldn't parse as diagnostics.
	OtherBuildOutput []string

	// FailedTests lists failing tests in the order they ran.
	FailedTests []TestFailure
	// TotalTests is how many tests ran (0 for packages with no test files).
	TotalTests int

	// PanicOutput is the tail of package output when the program panicked.
	PanicOutput string

	// TimedOut is true when verification was killed after the timeout —
	// usually an infinite loop or a channel deadlock.
	TimedOut bool

	// GoError is set when the go command itself failed (missing toolchain,
	// bad environment, ...), as opposed to the exercise failing.
	GoError string

	Duration time.Duration
}

// Run verifies the exercise package (relative to root, e.g.
// "exercises/00_intro/01_hello") by running go test in root.
func Run(root, pkgPath string, timeout time.Duration) *Result {
	res := &Result{Package: pkgPath}
	start := time.Now()
	if timeout <= 0 {
		timeout = DefaultTimeout
	}

	goBin, err := exec.LookPath("go")
	if err != nil {
		res.GoError = "the go toolchain was not found on PATH — install Go from https://go.dev/dl/"
		res.Duration = time.Since(start)
		return res
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, goBin, "test", "-count=1", "-json", "./"+pkgPath)
	cmd.Dir = root
	// Isolate the learner's project from ambient Go configuration.
	cmd.Env = append(os.Environ(), "GOWORK=off", "GOTOOLCHAIN=local")

	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	runErr := cmd.Run()
	res.Duration = time.Since(start)

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		res.TimedOut = true
		return res
	}
	// A non-nil runErr that is not an ExitError means we never got a proper
	// test run (bad flags, unreadable dir, ...). ExitError is normal when
	// tests fail or the package doesn't compile.
	var exitErr *exec.ExitError
	if runErr != nil && !errors.As(runErr, &exitErr) {
		res.GoError = strings.TrimSpace(stderr.String())
		if res.GoError == "" {
			res.GoError = runErr.Error()
		}
		return res
	}

	parse(stdout.String(), res)

	if !res.parsedAnything() && stderr.Len() > 0 {
		// `go test -json` prints machine output to stdout; anything substantial
		// on stderr means go itself complained (e.g. missing go.mod).
		res.GoError = strings.TrimSpace(stderr.String())
	}

	if res.TimedOut || res.GoError != "" {
		res.OK = false
	}
	return res
}

func (r *Result) parsedAnything() bool {
	return len(r.BuildErrors) > 0 || len(r.OtherBuildOutput) > 0 ||
		len(r.FailedTests) > 0 || r.TotalTests > 0
}

// testEvent mirrors the JSON events emitted by `go test -json`.
type testEvent struct {
	Action  string  `json:"Action"`
	Package string  `json:"Package"`
	Test    string  `json:"Test"`
	Output  string  `json:"Output"`
	Elapsed float64 `json:"Elapsed"`
}

var reCompileError = regexp.MustCompile(`^(?:\./)?([^\s:]+\.go):(\d+):(\d+):\s*(.*)$`)

// noiseLines are removed from failing-test output before display.
var noiseMatchers = []string{
	"=== RUN", "=== PAUSE", "=== CONT", "=== NAME", "=== BENCH",
	"--- FAIL: ", "--- PASS: ", "--- SKIP: ",
}

func parse(out string, res *Result) {
	var (
		pkgStatus string
		buildOut  strings.Builder
		pkgOut    strings.Builder
		testOut   = map[string]*strings.Builder{}
		order     []string // tests in first-seen order
		status    = map[string]string{}
	)

	sc := bufio.NewScanner(strings.NewReader(out))
	sc.Buffer(make([]byte, 0, 64*1024), 4*1024*1024)
	for sc.Scan() {
		line := sc.Bytes()
		if len(bytes.TrimSpace(line)) == 0 {
			continue
		}
		var ev testEvent
		if err := json.Unmarshal(line, &ev); err != nil {
			continue // non-JSON chatter; ignore
		}
		switch ev.Action {
		case "build-output":
			buildOut.WriteString(ev.Output)
		case "output":
			if ev.Test != "" {
				b := testOut[ev.Test]
				if b == nil {
					b = &strings.Builder{}
					testOut[ev.Test] = b
				}
				b.WriteString(ev.Output)
			} else {
				pkgOut.WriteString(ev.Output)
			}
		case "run":
			if ev.Test != "" && status[ev.Test] == "" {
				order = append(order, ev.Test)
			}
		case "pass", "fail", "skip":
			if ev.Test != "" {
				status[ev.Test] = ev.Action
			} else if pkgStatus == "" {
				pkgStatus = ev.Action
			} else {
				pkgStatus = ev.Action // keep the last package-level verdict
			}
		}
	}

	res.TotalTests = len(status)

	// Parse build output into diagnostics.
	for _, line := range strings.Split(strings.TrimRight(buildOut.String(), "\n"), "\n") {
		line = strings.TrimRight(line, "\r")
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if m := reCompileError.FindStringSubmatch(line); m != nil {
			res.BuildErrors = append(res.BuildErrors, CompileError{File: m[1], Line: m[2], Col: m[3], Msg: m[4]})
			continue
		}
		res.OtherBuildOutput = append(res.OtherBuildOutput, line)
	}

	// Collect failing tests in run order.
	for _, t := range order {
		if status[t] != "fail" {
			continue
		}
		f := TestFailure{Test: t, Output: cleanTestOutput(testOut[t])}
		res.FailedTests = append(res.FailedTests, f)
	}

	// Detect panics and deadlocks in package-level output.
	so := pkgOut.String()
	if strings.Contains(so, "panic: ") || strings.Contains(so, "fatal error: ") {
		res.PanicOutput = tail(so, 50)
	}

	// The verdict comes from the package-level result event plus whatever
	// failure modes we found above. Run() additionally rules out timeouts
	// and go-command errors.
	res.OK = pkgStatus == "pass" &&
		len(res.BuildErrors) == 0 &&
		len(res.OtherBuildOutput) == 0 &&
		len(res.FailedTests) == 0 &&
		res.PanicOutput == ""
}

func cleanTestOutput(b *strings.Builder) string {
	if b == nil {
		return ""
	}
	var kept []string
	for _, line := range strings.Split(b.String(), "\n") {
		line = strings.TrimRight(line, "\r")
		if strings.TrimSpace(line) == "" {
			continue
		}
		noisy := false
		for _, n := range noiseMatchers {
			if strings.HasPrefix(line, n) {
				noisy = true
				break
			}
		}
		switch {
		case noisy:
			continue
		case line == "PASS" || line == "FAIL":
			continue
		case strings.HasPrefix(line, "ok "), strings.HasPrefix(line, "FAIL\t"), strings.HasPrefix(line, "exit status"):
			continue
		}
		kept = append(kept, line)
	}
	return strings.Join(kept, "\n")
}

// tail returns the last n lines of s.
func tail(s string, n int) string {
	lines := strings.Split(strings.TrimRight(s, "\n"), "\n")
	if len(lines) > n {
		lines = lines[len(lines)-n:]
	}
	return strings.Join(lines, "\n")
}
