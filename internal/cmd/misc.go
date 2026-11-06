package cmd

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/cyberneticangel/gopherlings/internal/ui"
)

// runList shows the curriculum with per-exercise status.
func runList(rootFS fs.FS, args []string) int {
	flags := flag.NewFlagSet("list", flag.ContinueOnError)
	if code := parseFlags("list", flags, args); code != 0 {
		return code
	}
	s, code := openSession(rootFS)
	if s == nil {
		return code
	}

	total := len(s.store.Ordered())
	fmt.Println(ui.Bold("gopherlings") + ui.Dimf(" %s", Version) +
		ui.Greenf("  %d/%d complete", s.state.DoneCount(), total))

	maxID := 0
	for _, e := range s.store.Ordered() {
		if len(e.ID) > maxID {
			maxID = len(e.ID)
		}
	}

	lastSection := ""
	for _, e := range s.store.Ordered() {
		if e.Section != lastSection {
			lastSection = e.Section
			title := ""
			if sec, ok := s.store.Manifest().SectionOf(e); ok {
				title = sec.Title
			}
			fmt.Println()
			fmt.Println(ui.Boldf("%s", e.Section) + ui.Dimf("  %s", title))
		}
		var status string
		switch {
		case s.state.IsCompleted(e.Key()):
			status = ui.Green("✓")
		case e.Key() == s.state.Current:
			status = ui.Blue("▶")
		default:
			status = ui.Dim("·")
		}
		name := e.Name
		if e.Key() == s.state.Current {
			name = ui.Bold(name)
		}
		fmt.Printf("  %s %-*s  %s\n", status, maxID, e.ID, name)
	}
	fmt.Println()
	return 0
}

// runHint prints the hints for the current (or named) exercise.
func runHint(rootFS fs.FS, args []string) int {
	flags := flag.NewFlagSet("hint", flag.ContinueOnError)
	if code := parseFlags("hint", flags, args); code != 0 {
		return code
	}
	s, code := openSession(rootFS)
	if s == nil {
		return code
	}
	key := s.state.Current
	if flags.NArg() > 0 {
		e, err := s.store.Find(flags.Arg(0))
		if err != nil {
			fmt.Println(ui.Red(err.Error()))
			return 1
		}
		key = e.Key()
	}
	ex, err := s.store.Get(key)
	if err != nil {
		fmt.Println(ui.Red(err.Error()))
		return 1
	}
	if len(ex.Hints) == 0 {
		fmt.Println(ui.Dim("No hints for this one — you've got this."))
		return 0
	}
	fmt.Println(ui.Boldf("Hints for %s:", key))
	fmt.Println()
	for i, h := range ex.Hints {
		fmt.Printf("  %s %s\n", ui.Yellow(fmt.Sprintf("%d.", i+1)), h)
	}
	return 0
}

// runReset restores an exercise to its pristine embedded state.
func runReset(rootFS fs.FS, args []string) int {
	flags := flag.NewFlagSet("reset", flag.ContinueOnError)
	if code := parseFlags("reset", flags, args); code != 0 {
		return code
	}
	s, code := openSession(rootFS)
	if s == nil {
		return code
	}
	key := s.state.Current
	if flags.NArg() > 0 {
		e, err := s.store.Find(flags.Arg(0))
		if err != nil {
			fmt.Println(ui.Red(err.Error()))
			return 1
		}
		key = e.Key()
	}
	ex, err := s.store.Get(key)
	if err != nil {
		fmt.Println(ui.Red(err.Error()))
		return 1
	}
	dir := s.exDir(key)

	// Remove stray Go files the learner added.
	if entries, err := os.ReadDir(dir); err == nil {
		for _, entry := range entries {
			if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
				continue
			}
			if _, ok := ex.Files[entry.Name()]; !ok {
				_ = os.Remove(filepath.Join(dir, entry.Name()))
			}
		}
	}
	for name, content := range ex.Files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
			fmt.Println(ui.Redf("reset failed: %v", err))
			return 1
		}
	}

	wasCompleted := s.state.IsCompleted(key)
	if wasCompleted {
		s.state.Uncomplete(key)
	}
	if key != s.state.Current {
		s.state.SetCurrent(key)
	}
	if err := s.state.Save(s.root); err != nil {
		fmt.Println(ui.Yellowf("warning: could not save progress: %v", err))
	}
	fmt.Println(ui.Greenf("Reset %s to its original state.", key))
	if wasCompleted {
		fmt.Println(ui.Dim("It was marked complete; it is now your current exercise again."))
	}
	return 0
}

// runSolution prints the embedded solution for an exercise.
func runSolution(rootFS fs.FS, args []string) int {
	flags := flag.NewFlagSet("solution", flag.ContinueOnError)
	if code := parseFlags("solution", flags, args); code != 0 {
		return code
	}
	s, code := openSession(rootFS)
	if s == nil {
		return code
	}
	key := s.state.Current
	if flags.NArg() > 0 {
		e, err := s.store.Find(flags.Arg(0))
		if err != nil {
			fmt.Println(ui.Red(err.Error()))
			return 1
		}
		key = e.Key()
	}
	files, err := s.store.SolutionFiles(key)
	if err != nil {
		fmt.Println(ui.Redf("no solution embedded for %s", key))
		return 1
	}
	fmt.Println(ui.Dim("No shame in a peek — but try `gopherlings hint` first next time."))
	fmt.Println()
	names := make([]string, 0, len(files))
	for n := range files {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, n := range names {
		fmt.Println(ui.Cyan("── exercises/" + key + "/" + n + " (solution)"))
		fmt.Println(files[n])
	}
	return 0
}
