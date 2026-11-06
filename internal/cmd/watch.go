package cmd

import (
	"context"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/cyberneticangel/gopherlings/internal/ui"
	"github.com/cyberneticangel/gopherlings/internal/verify"
	"github.com/cyberneticangel/gopherlings/internal/watcher"
)

// runWatch implements `gopherlings watch` (and bare `gopherlings`): verify
// the current exercise at startup and again after every saved change.
func runWatch(rootFS fs.FS, args []string) int {
	flags := flag.NewFlagSet("watch", flag.ContinueOnError)
	timeout := flags.Duration("timeout", verify.DefaultTimeout, "per-exercise verification timeout")
	if code := parseFlags("watch", flags, args); code != 0 {
		return code
	}

	s, code := openSession(rootFS)
	if s == nil {
		return code
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	w, err := watcher.Start(filepath.Join(s.root, "exercises"), 300*time.Millisecond)
	if err != nil {
		fmt.Println(ui.Redf("cannot watch the exercises directory: %v", err))
		fmt.Println(ui.Dim("Use `gopherlings run` to verify manually instead."))
		return 1
	}
	defer w.Close()

	fmt.Println(ui.Boldf("gopherlings %s", Version) +
		ui.Dimf(" — %d exercises, %d done", len(s.store.Ordered()), s.state.DoneCount()))
	fmt.Println()

	sr, _ := s.doStep(s.state.Current, *timeout)
	for sr != stepFinished {
		fmt.Println()
		fmt.Println(ui.Dim("Watching for changes… (Ctrl+C to quit)"))
		select {
		case <-ctx.Done():
			fmt.Println()
			fmt.Println(ui.Dim("Bye! 👋"))
			return 0
		case ev := <-w.Events():
			if ev.Err != nil {
				fmt.Println(ui.Yellowf("watcher: %v", ev.Err))
				continue
			}
			sr, _ = s.doStep(s.state.Current, *timeout)
		}
	}
	return 0
}
