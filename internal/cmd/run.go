package cmd

import (
	"flag"
	"fmt"
	"io/fs"
	"time"

	"github.com/cyberneticangel/gopherlings/internal/ui"
	"github.com/cyberneticangel/gopherlings/internal/verify"
)

type stepResult int

const (
	stepFail stepResult = iota
	stepWaitingMarker
	stepCompleted
	stepFinished
)

// doStep verifies one exercise, prints the result, and updates progress:
// a pass with the marker still present waits; a pass without the marker
// completes the exercise and moves to the next one.
func (s *session) doStep(key string, timeout time.Duration) (stepResult, *verify.Result) {
	ex, err := s.store.Get(key)
	if err != nil {
		fmt.Println(ui.Redf("cannot load exercise %q: %v", key, err))
		return stepFail, &verify.Result{}
	}

	res := verify.Run(s.root, "exercises/"+key, timeout)
	printResult(ex, res)

	if !res.OK {
		printHintReminder(ex)
		return stepFail, res
	}

	if s.markerPresent(key) {
		fmt.Println()
		fmt.Println(ui.Green("✓ It compiles and the tests pass!"))
		fmt.Println()
		fmt.Printf("  Now delete the %s comment from\n  %s to continue.\n",
			ui.Yellow(marker), ui.Dim("exercises/"+key+"/exercise.go"))
		return stepWaitingMarker, res
	}

	s.state.Complete(key)
	next := s.nextAfter(key)
	if next == nil {
		_ = s.state.Save(s.root)
		printCongrats(s.state.DoneCount())
		return stepFinished, res
	}
	s.state.SetCurrent(next.Key())
	if err := s.state.Save(s.root); err != nil {
		fmt.Println(ui.Yellowf("warning: could not save progress: %v", err))
	}
	fmt.Println()
	fmt.Println(ui.Greenf("✓ Complete: %s", ex.Name))
	fmt.Println()
	s.printBanner(*next)
	return stepCompleted, res
}

// runRun implements `gopherlings run [name]`. Without a name it verifies the
// current exercise and advances progress; with a name it verifies that
// exercise only, leaving progress untouched.
func runRun(rootFS fs.FS, args []string) int {
	flags := flag.NewFlagSet("run", flag.ContinueOnError)
	timeout := flags.Duration("timeout", verify.DefaultTimeout, "per-exercise verification timeout")
	if code := parseFlags("run", flags, args); code != 0 {
		return code
	}

	s, code := openSession(rootFS)
	if s == nil {
		return code
	}

	if flags.NArg() == 0 {
		sr, _ := s.doStep(s.state.Current, *timeout)
		if sr == stepFail {
			return 1
		}
		return 0
	}

	e, err := s.store.Find(flags.Arg(0))
	if err != nil {
		fmt.Println(ui.Red(err.Error()))
		return 1
	}
	ex, err := s.store.Get(e.Key())
	if err != nil {
		fmt.Println(ui.Red(err.Error()))
		return 1
	}
	res := verify.Run(s.root, "exercises/"+e.Key(), *timeout)
	printResult(ex, res)
	if !res.OK {
		return 1
	}
	if s.markerPresent(e.Key()) {
		fmt.Println()
		fmt.Printf("  Delete the %s comment from %s when you're ready.\n",
			ui.Yellow(marker), ui.Dim("exercises/"+e.Key()+"/exercise.go"))
	}
	return 0
}
