// Package watcher reports filesystem activity under the exercises tree,
// coalesced with a small debounce so editor save bursts trigger one verify.
package watcher

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

// Event is a batch of changed file paths (or a watcher error).
type Event struct {
	Paths []string
	Err   error
}

// Watcher watches root recursively for changes to Go source files.
type Watcher struct {
	events chan Event
	stop   chan struct{}
	done   chan struct{}
}

// Start begins watching root (recursively). Only changes to .go files are
// reported; everything is coalesced per debounce window.
func Start(root string, debounce time.Duration) (*Watcher, error) {
	notif, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	if err := addRecursive(notif, root); err != nil {
		notif.Close()
		return nil, err
	}
	w := &Watcher{
		events: make(chan Event, 4),
		stop:   make(chan struct{}),
		done:   make(chan struct{}),
	}
	go w.loop(notif, debounce)
	return w, nil
}

// Events returns the event stream. Events may be dropped if the consumer
// falls behind; the verify loop always re-checks the current state anyway.
func (w *Watcher) Events() <-chan Event { return w.events }

// Close stops watching and releases resources.
func (w *Watcher) Close() {
	close(w.stop)
	<-w.done
}

func (w *Watcher) loop(notif *fsnotify.Watcher, debounce time.Duration) {
	defer close(w.done)
	defer notif.Close()

	var mu sync.Mutex
	pending := map[string]bool{}
	var timer *time.Timer

	flush := func() {
		mu.Lock()
		paths := make([]string, 0, len(pending))
		for p := range pending {
			paths = append(paths, p)
		}
		pending = map[string]bool{}
		timer = nil
		mu.Unlock()
		select {
		case w.events <- Event{Paths: paths}:
		default: // consumer busy: drop; the next event re-verifies anyway
		}
	}

	record := func(path string) {
		if !strings.HasSuffix(path, ".go") {
			return
		}
		mu.Lock()
		pending[path] = true
		if timer != nil {
			timer.Stop()
		}
		timer = time.AfterFunc(debounce, flush)
		mu.Unlock()
	}

	for {
		select {
		case <-w.stop:
			mu.Lock()
			if timer != nil {
				timer.Stop()
			}
			mu.Unlock()
			return
		case err := <-notif.Errors:
			select {
			case w.events <- Event{Err: err}:
			default:
			}
		case ev, ok := <-notif.Events:
			if !ok {
				return
			}
			if ev.Has(fsnotify.Create) {
				if fi, err := os.Stat(ev.Name); err == nil && fi.IsDir() {
					if err := addRecursive(notif, ev.Name); err != nil {
						select {
						case w.events <- Event{Err: err}:
						default:
						}
					}
				}
			}
			record(ev.Name)
		}
	}
}

// addRecursive adds dir and all its subdirectories to the watch list.
func addRecursive(notif *fsnotify.Watcher, root string) error {
	return filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			// Skip hidden dirs like .git.
			if strings.HasPrefix(d.Name(), ".") && path != root {
				return filepath.SkipDir
			}
			return notif.Add(path)
		}
		return nil
	})
}
