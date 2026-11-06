package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cyberneticangel/gopherlings/internal/cmd"
	"github.com/cyberneticangel/gopherlings/internal/exercise"
	"github.com/cyberneticangel/gopherlings/internal/manifest"
)

// newStore builds a Store from the embedded content.
func newStore(t *testing.T) *exercise.Store {
	t.Helper()
	store, err := exercise.NewStore(content)
	if err != nil {
		t.Fatalf("embedded content is invalid: %v", err)
	}
	return store
}

// scaffold runs `gopherlings init` in a fresh temp directory and returns it.
func scaffold(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	if code := cmd.Main(content, []string{"init", dir}); code != 0 {
		t.Fatalf("gopherlings init failed with exit code %d", code)
	}
	return dir
}

// TestManifestIntegrity checks authoring invariants that the runner relies on:
// every exercise has an exercise.go with the marker and an exercise_test.go,
// and every solution is marker-free.
func TestManifestIntegrity(t *testing.T) {
	store := newStore(t)
	for _, e := range store.Ordered() {
		ex, err := store.Get(e.Key())
		if err != nil {
			t.Errorf("%s: %v", e.Key(), err)
			continue
		}
		exSrc, ok := ex.Files["exercise.go"]
		if !ok {
			t.Errorf("%s: missing exercise.go", e.Key())
			continue
		}
		if !strings.Contains(exSrc, "// I AM NOT DONE") {
			t.Errorf("%s: exercise.go does not contain the I AM NOT DONE marker", e.Key())
		}
		if _, ok := ex.Files["exercise_test.go"]; !ok {
			t.Errorf("%s: missing exercise_test.go", e.Key())
		}
		sol, err := store.SolutionFiles(e.Key())
		if err != nil {
			t.Errorf("%s: no embedded solution: %v", e.Key(), err)
			continue
		}
		solSrc, ok := sol["exercise.go"]
		if !ok {
			t.Errorf("%s: solution is missing exercise.go", e.Key())
			continue
		}
		if strings.Contains(solSrc, "// I AM NOT DONE") {
			t.Errorf("%s: solution still contains the I AM NOT DONE marker", e.Key())
		}
	}
}

// TestSolutionsPass is the big one: for every exercise, scaffold a fresh
// project, assert the exercise starts broken, apply the embedded solution,
// and assert it then passes verification — all through the real CLI.
func TestSolutionsPass(t *testing.T) {
	if _, err := exercise.NewStore(content); err != nil {
		t.Fatalf("embedded content is invalid: %v", err)
	}
	store := newStore(t)
	ordered := store.Ordered()
	if len(ordered) < 3 {
		t.Fatalf("expected a real curriculum, found only %d exercises", len(ordered))
	}
	for _, e := range ordered {
		e := e
		t.Run(e.Key(), func(t *testing.T) {
			dir := scaffold(t)
			t.Chdir(dir)

			// The exercise must be broken before being solved...
			if code := cmd.Main(content, []string{"run", e.ID}); code == 0 {
				t.Errorf("exercise %s already passes before being solved — it must start broken", e.Key())
			}

			// ...and pass once the solution is applied.
			files, err := store.SolutionFiles(e.Key())
			if err != nil {
				t.Fatalf("no solution: %v", err)
			}
			exDir := filepath.Join(dir, "exercises", filepath.FromSlash(e.Key()))
			for name, src := range files {
				if err := os.WriteFile(filepath.Join(exDir, name), []byte(src), 0o644); err != nil {
					t.Fatalf("write solution: %v", err)
				}
			}
			if code := cmd.Main(content, []string{"run", e.ID}); code != 0 {
				t.Errorf("solution for %s does not pass verification", e.Key())
			}
		})
	}
}

// TestSectionOrdering checks that exercises appear in the same order as their
// sections and that every section has at least one exercise.
func TestSectionOrdering(t *testing.T) {
	store := newStore(t)
	man := store.Manifest()
	seen := map[string]bool{}
	lastIdx := -1
	sectionIdx := map[string]int{}
	for i, s := range man.Sections {
		sectionIdx[s.ID] = i
	}
	for _, e := range man.Exercises {
		idx := sectionIdx[e.Section]
		if seen[e.Section] && idx < lastIdx {
			t.Errorf("exercise %s: section %s appears out of order in the manifest", e.Key(), e.Section)
		}
		seen[e.Section] = true
		lastIdx = idx
	}
	for _, s := range man.Sections {
		if !seen[s.ID] {
			t.Errorf("section %s has no exercises", s.ID)
		}
	}
}

var _ = manifest.Exercise{} // keep import if manifest API evolves
