# Authoring gopherlings exercises

This guide defines how exercises for gopherlings are written, structured, and
validated. Read it fully before authoring.

## The idea

Each exercise is a **small, deliberately broken Go program** with a test that
checks the fix. The learner makes the test pass, then deletes the
`// I AM NOT DONE` comment to move on. Exercises teach exactly one idea each
and ramp in difficulty within their section.

## Layout

Every exercise is its own directory and its own package:

```
exercises/<section>/<NN_topic>/exercise.go        # the broken program (package main)
exercises/<section>/<NN_topic>/exercise_test.go   # the tests that check the fix
solutions/<section>/<NN_topic>/exercise.go        # the solved program (no marker!)
```

- `<section>` is the section id, e.g. `06_slices`.
- `<NN_topic>` starts with a two-digit sequence number, e.g. `03_append`.
- IDs must be **globally unique and descriptive** across the whole curriculum
  (never a generic `01_basics` — a learner may type just the id).
- Only `exercise.go` (plus, rarely, `exercise_test.go`) lives in the exercise
  dir. Extra helper files are not allowed without a very good reason.

## Rules

1. **`package main` always.** Include a small `main()` that prints something
   related to the task, so learners can also run their code with
   `go run ./exercises/<section>/<NN_topic>/`.
2. **The broken state must fail.** `go test ./exercises/<section>/<NN_topic>/`
   must fail on the untouched exercise — with a compile error or a failing
   test. Both are fine; pick whichever teaches the lesson better.
   - `kind: "compile"` if the initial state does not compile.
   - `kind: "test"` if it compiles but the tests fail.
   - The broken state must fail **fast**. Never ship an initial state that
     busy-loops or hangs; a full runtime deadlock ("all goroutines are
     asleep") is acceptable, a spin loop is not.
3. **The solution must pass.** With the solution's `exercise.go` copied over
   the exercise's, `go test` must pass and `go vet` must be clean (`go test`
   runs vet automatically).
4. **Standard library only.** No external imports, no network, no file I/O
   outside `t.TempDir()`, no time/rand dependence unless seeded, no
   interactivity. Deterministic and fast (solved state finishes well under
   2 seconds).
5. **The marker.** `exercise.go` ends with the line `// I AM NOT DONE`.
   Solutions must not contain the marker — not even in comments. Never
   mention the literal string `// I AM NOT DONE` inside a solution.
6. **Tests are the contract.** Tests must fail with helpful messages using
   `t.Errorf("f(%v) = %v, want %v", ...)` style. Table-driven tests
   (`[]struct{...}` cases or maps) are preferred.
7. **Doc comments teach.** `exercise.go` starts with a doc comment that
   explains the concept being taught and states the task precisely. Concise,
   friendly, imperative. Mark the work with `// TODO:` comments at the exact
   spots that need fixing.
8. **Formatting.** All files must be gofmt-clean (`gofmt -l` prints nothing).
9. **Difficulty ramp.** Within a section, exercises get harder. No exercise
   should require knowledge from a *later* section.

## The marker flow (for reference)

The runner verifies with `go test -count=1 -json ./exercises/<section>/<NN_topic>`.
A pass with the marker present prints "delete the comment to continue"; a pass
without the marker completes the exercise and advances.

## Manifest fragments

Authors do **not** edit `exercises/manifest.json` (it is merged centrally).
Instead, provide one fragment file per section: a JSON **array** of exercise
objects, in curriculum order:

```json
[
  {
    "id": "03_append",
    "section": "06_slices",
    "name": "Growing Slices",
    "kind": "test",
    "blurb": "One short sentence shown nowhere yet, keep it.",
    "hints": [
      "First hint: a nudge, not the answer.",
      "Second hint: the shape of the solution.",
      "Third hint: nearly the answer."
    ]
  }
]
```

- `kind` is `"compile"` or `"test"`.
- 1–3 hints per exercise. Progressive: nudge → shape → answer.

## Validation workflow (mandatory, per exercise)

Work in a scratch module so the intentionally-broken files never interfere
with the repo build:

```bash
SCRATCH=/tmp/opencode/scratch-<section>
mkdir -p "$SCRATCH" && cd "$SCRATCH"
printf 'module scratch\n\ngo 1.24\n' > go.mod

# for each exercise:
cp -r /home/dylan/Projects/gophers/exercises/<section>/<NN_topic> .
timeout 60 go test ./<NN_topic>/ ; echo "broken-state exit: $?"   # must FAIL
cp /home/dylan/Projects/gophers/solutions/<section>/<NN_topic>/exercise.go <NN_topic>/exercise.go
timeout 60 go test ./<NN_topic>/ ; echo "solved exit: $?"         # must PASS
gofmt -l .                                                        # must print nothing
```

Also check the broken state **fails fast** (the `timeout 60` guard catches
hangs). If an exercise uses a solution-side `exercise_test.go` override,
copy that too in the solved check.

Report every exercise with: broken-state result, solved result, gofmt result.

## Canonical example

`exercises/00_intro/01_hello/exercise.go`:

```go
// Welcome to gopherlings! 🐹
//
// This project teaches Go through small, broken programs. ...
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
```

with `exercise_test.go`:

```go
package main

import "testing"

func TestGreeting(t *testing.T) {
	want := "Hello, Gopher!"
	if got := greeting(); got != want {
		t.Errorf("greeting() = %q, want %q", got, want)
	}
}
```

and the solution replacing `Message` with `message` and dropping the marker.

## Curriculum overview (planned)

| Section | Title | Theme |
|---------|-------|-------|
| 00_intro | Hello, Gopherlings | workflow, fmt |
| 01_variables | Variables & Constants | var, :=, zero values, const, iota |
| 02_basic_types | Basic Types | ints, floats, bools, conversions |
| 03_functions | Functions | returns, variadic, closures |
| 04_control_flow | Control Flow | if, for, switch, labels |
| 05_structs | Structs & Methods | literals, receivers, embedding |
| 06_slices | Slices | append, aliasing, slicing, 2D |
| 07_maps | Maps | comma-ok, delete, iteration |
| 08_strings | Strings, Runes & Bytes | UTF-8, []rune, strings pkg, Builder |
| 09_pointers | Pointers | &, *, nil, swap |
| 10_interfaces | Interfaces | satisfaction, type switches, Stringer, method sets |
| 11_errors | Errors | sentinels, wrapping, errors.Is/As, custom types |
| 12_generics | Generics | type params, constraints, generic types |
| 13_concurrency | Concurrency | channels, select, WaitGroup, Mutex, context |
| 14_stdlib | Standard Library | time, json, io, defer/recover, writing tests |
