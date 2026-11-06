// Package cmd implements the gopherlings command-line interface.
package cmd

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/cyberneticangel/gopherlings/internal/exercise"
	"github.com/cyberneticangel/gopherlings/internal/manifest"
	"github.com/cyberneticangel/gopherlings/internal/state"
	"github.com/cyberneticangel/gopherlings/internal/ui"
	"github.com/cyberneticangel/gopherlings/internal/verify"
)

// Version is the gopherlings version.
const Version = "0.1.0"

// goDirective is the `go` line written into scaffolded projects. Exercises
// use only features available in this toolchain release.
const goDirective = "1.24"

// marker is the comment learners delete to move on.
const marker = "// I AM NOT DONE"

// Main dispatches a command and returns the process exit code.
func Main(rootFS fs.FS, args []string) int {
	ui.Init()
	if len(args) == 0 {
		args = []string{"watch"} // bare `gopherlings` behaves like watch
	}
	switch args[0] {
	case "init":
		return runInit(rootFS, args[1:])
	case "watch":
		return runWatch(rootFS, args[1:])
	case "run":
		return runRun(rootFS, args[1:])
	case "list":
		return runList(rootFS, args[1:])
	case "hint":
		return runHint(rootFS, args[1:])
	case "reset":
		return runReset(rootFS, args[1:])
	case "solution":
		return runSolution(rootFS, args[1:])
	case "version", "--version", "-v":
		fmt.Println("gopherlings", Version)
		return 0
	case "help", "-h", "--help":
		printUsage(os.Stdout)
		return 0
	default:
		fmt.Println(ui.Redf("unknown command %q", args[0]))
		fmt.Println()
		printUsage(os.Stdout)
		return 1
	}
}

func printUsage(w *os.File) {
	fmt.Fprintln(w, ui.Bold("gopherlings")+ui.Dimf(" %s — learn Go by fixing small exercises", Version))
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, ui.Cyan("  gopherlings init [dir]")+"      create a new gopherlings project")
	fmt.Fprintln(w, ui.Cyan("  gopherlings")+"                 watch mode: verify on every save")
	fmt.Fprintln(w, ui.Cyan("  gopherlings run [name]")+"     verify the current (or named) exercise once")
	fmt.Fprintln(w, ui.Cyan("  gopherlings list")+"           show the curriculum and your progress")
	fmt.Fprintln(w, ui.Cyan("  gopherlings hint [name]")+"    show hints for an exercise")
	fmt.Fprintln(w, ui.Cyan("  gopherlings reset [name]")+"   restore an exercise to its original state")
	fmt.Fprintln(w, ui.Cyan("  gopherlings solution [name]")+"  show a completed solution (no peeking!)")
	fmt.Fprintln(w, ui.Cyan("  gopherlings version")+"        print the version")
}

// session bundles everything needed to work inside a scaffolded project.
type session struct {
	root  string
	store *exercise.Store
	state *state.State
}

// openSession locates the enclosing gopherlings project and loads it.
// It returns (nil, exitCode) when the project can't be opened.
func openSession(rootFS fs.FS) (*session, int) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(ui.Redf("cannot determine working directory: %v", err))
		return nil, 1
	}
	root, ok := state.FindProject(cwd)
	if !ok {
		fmt.Println(ui.Red("Not inside a gopherlings project (no .gopherlings directory found)."))
		fmt.Println(ui.Dim("Run `gopherlings init` to create one, or `cd` into your project directory."))
		return nil, 1
	}
	store, err := exercise.NewStore(rootFS)
	if err != nil {
		fmt.Println(ui.Redf("broken installation: %v", err))
		return nil, 1
	}
	st, err := state.Load(root)
	if err != nil {
		fmt.Println(ui.Redf("cannot load progress: %v", err))
		fmt.Println(ui.Dim("Delete the .gopherlings directory and run `gopherlings init` again to start over."))
		return nil, 1
	}
	if _, err := store.Get(st.Current); err != nil {
		fmt.Println(ui.Redf("progress file points at unknown exercise %q.", st.Current))
		fmt.Println(ui.Dim("Fix or delete .gopherlings/progress.json to recover."))
		return nil, 1
	}
	return &session{root: root, store: store, state: st}, 0
}

// exDir is the on-disk directory of an exercise within the project.
func (s *session) exDir(key string) string {
	return filepath.Join(s.root, "exercises", filepath.FromSlash(key))
}

// markerPresent reports whether any Go file in the exercise still contains
// the "I AM NOT DONE" marker.
func (s *session) markerPresent(key string) bool {
	dir := s.exDir(key)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			continue
		}
		if strings.Contains(string(data), marker) {
			return true
		}
	}
	return false
}

// nextAfter returns the exercise following key, or nil if key is the last one.
func (s *session) nextAfter(key string) *manifest.Exercise {
	list := s.store.Ordered()
	for i, e := range list {
		if e.Key() == key {
			if i+1 < len(list) {
				n := list[i+1]
				return &n
			}
			return nil
		}
	}
	return nil
}

// printResult renders a verification result for one exercise.
func printResult(ex exercise.Exercise, res *verify.Result) {
	fmt.Println(ui.Dimf("── exercises/%s · %s", ex.Key(), ex.Name))

	switch {
	case res.TimedOut:
		fmt.Println(ui.Yellowf("⏱  Verification timed out. An infinite loop, or a channel that never unblocks?"))
		return
	case res.GoError != "":
		fmt.Println(ui.Redf("✗ go failed: %s", res.GoError))
		return
	}

	if len(res.BuildErrors) > 0 || len(res.OtherBuildOutput) > 0 {
		fmt.Println(ui.Redf("✗ It doesn't compile yet:"))
		shown := res.BuildErrors
		if len(shown) > 12 {
			shown = shown[:12]
			defer fmt.Println(ui.Dimf("   … and %d more", len(res.BuildErrors)-12))
		}
		for _, ce := range shown {
			fmt.Printf("  %s  %s\n", ui.Yellow(ce.File+":"+ce.Line+":"+ce.Col), ce.Msg)
		}
		for _, line := range res.OtherBuildOutput {
			fmt.Println("  " + ui.Dim(line))
		}
		return
	}

	if res.PanicOutput != "" {
		fmt.Println(ui.Red("✗ The program panicked:"))
		fmt.Println(ui.Indent(res.PanicOutput, "  "+ui.Dim("|")))
		return
	}

	if len(res.FailedTests) > 0 {
		fmt.Println(ui.Redf("✗ %d of %d tests failed:", len(res.FailedTests), res.TotalTests))
		for _, ft := range res.FailedTests {
			fmt.Printf("  %s\n", ui.Yellow("✗ "+ft.Test))
			if ft.Output != "" {
				fmt.Println(ui.Indent(ft.Output, "      "))
			}
		}
		return
	}

	if res.TotalTests > 0 {
		fmt.Println(ui.Greenf("✓ All %d tests pass! (%.1fs)", res.TotalTests, res.Duration.Seconds()))
	} else {
		fmt.Println(ui.Greenf("✓ It compiles. (%.1fs)", res.Duration.Seconds()))
	}
}

func printHintReminder(ex exercise.Exercise) {
	if len(ex.Hints) > 0 {
		fmt.Println()
		fmt.Println(ui.Dim("Stuck? Run " + ui.Cyan("gopherlings hint") + ui.Dim(" for a nudge.")))
	}
}

// printBanner announces the next exercise.
func (s *session) printBanner(e manifest.Exercise) {
	fmt.Println(ui.Boldf("Next up: %s", e.Name))
	fmt.Println(ui.Dim("  exercises/" + e.Key() + "/exercise.go"))
	fmt.Println(ui.Dimf("  progress: %d/%d", s.state.DoneCount(), len(s.store.Ordered())))
}

// printCongrats celebrates finishing the whole curriculum.
func printCongrats(done int) {
	fmt.Println()
	fmt.Println(ui.Green("      .-\"\"\"-."))
	fmt.Println(ui.Green("     /       \\"))
	fmt.Println(ui.Green("    | .-. .-. |   ") + ui.Bold("gopherlings complete!"))
	fmt.Println(ui.Green("    | |o| |o| |   ") + ui.Greenf("You finished all %d exercises.", done))
	fmt.Println(ui.Green("    |    \"    |   ") + ui.Dim("Welcome to the burrow, Gopher."))
	fmt.Println(ui.Green("     \\  ---  /"))
	fmt.Println(ui.Green("      '-----'"))
	fmt.Println()
}

// parseFlags is a small wrapper giving consistent flag error handling.
func parseFlags(name string, flags *flag.FlagSet, args []string) int {
	if err := flags.Parse(args); err != nil {
		fmt.Println(ui.Redf("%s: %v", name, err))
		return 2
	}
	return 0
}
