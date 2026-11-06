package cmd

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cyberneticangel/gopherlings/internal/exercise"
	"github.com/cyberneticangel/gopherlings/internal/state"
	"github.com/cyberneticangel/gopherlings/internal/ui"
)

const scaffoldReadme = `# gopherlings playground

Learn Go by fixing small, broken programs.

## How it works

Each directory under ` + "`exercises/`" + ` is a small program with something wrong
or missing. Make the tests pass, then delete the ` + "`// I AM NOT DONE`" + ` comment
to move to the next exercise.

Run ` + "`gopherlings`" + ` (or ` + "`gopherlings watch`" + `) in this directory and it will
verify your work every time you save.

## Commands

- ` + "`gopherlings list`" + ` — see the curriculum and your progress
- ` + "`gopherlings hint`" + ` — get a nudge for the current exercise
- ` + "`gopherlings run <name>`" + ` — verify an exercise once
- ` + "`gopherlings reset <name>`" + ` — restore an exercise to its original state

Note: the exercises are *supposed* to contain compile errors, so ` + "`go build ./...`" + `
will complain. That's the point — verify exercises through gopherlings instead.
`

// runInit scaffolds a new gopherlings project with all exercises embedded
// in the binary.
func runInit(rootFS fs.FS, args []string) int {
	flags := flag.NewFlagSet("init", flag.ContinueOnError)
	force := flags.Bool("force", false, "overwrite existing exercise files")
	if code := parseFlags("init", flags, args); code != 0 {
		return code
	}

	store, err := exercise.NewStore(rootFS)
	if err != nil {
		fmt.Println(ui.Redf("broken installation: %v", err))
		return 1
	}

	target := "."
	if flags.NArg() > 0 {
		target = flags.Arg(0)
	}
	abs, err := filepath.Abs(target)
	if err != nil {
		fmt.Println(ui.Redf("invalid directory %q: %v", target, err))
		return 1
	}

	exercisesDir := filepath.Join(abs, "exercises")
	if _, err := os.Stat(exercisesDir); err == nil && !*force {
		fmt.Println(ui.Redf("%s already contains an exercises directory.", abs))
		fmt.Println(ui.Dim("Use --force to overwrite it with a fresh copy."))
		return 1
	}
	if err := os.MkdirAll(abs, 0o755); err != nil {
		fmt.Println(ui.Redf("cannot create %s: %v", abs, err))
		return 1
	}

	// Project module. Exercises use only the standard library, so no
	// dependencies are needed and everything works offline.
	gomod := "module gopherlings_playground\n\ngo " + goDirective + "\n"
	if err := os.WriteFile(filepath.Join(abs, "go.mod"), []byte(gomod), 0o644); err != nil {
		fmt.Println(ui.Redf("cannot write go.mod: %v", err))
		return 1
	}
	_ = os.WriteFile(filepath.Join(abs, ".gitignore"), []byte(".gopherlings/\n"), 0o644)
	_ = os.WriteFile(filepath.Join(abs, "README.md"), []byte(scaffoldReadme), 0o644)

	// Copy every exercise from the embedded copy.
	ordered := store.Ordered()
	for _, e := range ordered {
		ex, err := store.Get(e.Key())
		if err != nil {
			fmt.Println(ui.Redf("broken installation (exercise %s): %v", e.Key(), err))
			return 1
		}
		dir := filepath.Join(abs, "exercises", filepath.FromSlash(e.Key()))
		if err := os.MkdirAll(dir, 0o755); err != nil {
			fmt.Println(ui.Redf("cannot create %s: %v", dir, err))
			return 1
		}
		for name, content := range ex.Files {
			if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
				fmt.Println(ui.Redf("cannot write %s: %v", name, err))
				return 1
			}
		}
	}

	// Fresh progress state pointing at the first exercise.
	if _, err := state.Init(abs, ordered[0].Key()); err != nil {
		fmt.Println(ui.Redf("cannot initialize progress: %v", err))
		return 1
	}

	fmt.Println(ui.Greenf("🐹 Initialized a gopherlings playground in %s", abs))
	fmt.Println()
	fmt.Println(ui.Boldf("  %d exercises across %d sections are waiting for you.",
		len(ordered), len(store.Manifest().Sections)))
	fmt.Println()
	fmt.Println("  To begin:")
	fmt.Println()
	if target != "." {
		fmt.Printf("    %s\n", ui.Cyan("cd "+target))
	}
	fmt.Printf("    %s\n", ui.Cyan("gopherlings"))
	fmt.Println()
	fmt.Println(ui.Dim("  gopherlings will verify each exercise as you save your edits."))
	fmt.Println(ui.Dim("  Good luck — delete every 'I AM NOT DONE' you can find."))

	if _, err := exec.LookPath("go"); err != nil {
		fmt.Println()
		fmt.Println(ui.Yellow("⚠ The go toolchain was not found on PATH."))
		fmt.Println(ui.Yellow("  Install Go 1.24 or newer from https://go.dev/dl/ to verify exercises."))
	}
	return 0
}
