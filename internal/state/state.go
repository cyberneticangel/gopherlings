// Package state persists learner progress in the project directory under
// .gopherlings/progress.json.
package state

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	// DirName is the hidden directory holding project state.
	DirName = ".gopherlings"
	// FileName is the progress file inside DirName.
	FileName = "progress.json"
)

// State tracks which exercise the learner is on and which are complete.
type State struct {
	// Current is the key of the exercise being worked on.
	Current string `json:"current"`
	// Completed lists completed exercise keys in completion order.
	Completed []string `json:"completed"`
	// StartedAt is when the project was initialized.
	StartedAt time.Time `json:"started_at"`

	completedSet map[string]bool
}

// ProgressFile returns the path of the progress file for project root.
func ProgressFile(root string) string {
	return filepath.Join(root, DirName, FileName)
}

// Exists reports whether root (or an ancestor) is a gopherlings project.
// It returns the project root if found.
func FindProject(start string) (string, bool) {
	dir := start
	for i := 0; i < 32; i++ {
		if _, err := os.Stat(ProgressFile(dir)); err == nil {
			return dir, true
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", false
		}
		dir = parent
	}
	return "", false
}

// Init creates fresh state for a new project with firstKey as current.
func Init(root, firstKey string) (*State, error) {
	s := &State{
		Current:      firstKey,
		Completed:    []string{},
		StartedAt:    time.Now(),
		completedSet: map[string]bool{},
	}
	if err := s.Save(root); err != nil {
		return nil, err
	}
	return s, nil
}

// Load reads state from root's progress file.
func Load(root string) (*State, error) {
	data, err := os.ReadFile(ProgressFile(root))
	if err != nil {
		return nil, fmt.Errorf("read progress: %w", err)
	}
	var s State
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("parse %s: %w", ProgressFile(root), err)
	}
	if s.Current == "" {
		return nil, fmt.Errorf("%s is corrupt: no current exercise", ProgressFile(root))
	}
	if s.Completed == nil {
		s.Completed = []string{}
	}
	s.completedSet = make(map[string]bool, len(s.Completed))
	for _, k := range s.Completed {
		s.completedSet[k] = true
	}
	return &s, nil
}

// Save writes state back to root's progress file.
func (s *State) Save(root string) error {
	if err := os.MkdirAll(filepath.Join(root, DirName), 0o755); err != nil {
		return fmt.Errorf("create state dir: %w", err)
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	path := ProgressFile(root)
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write progress: %w", err)
	}
	return nil
}

// IsCompleted reports whether the exercise key has been completed.
func (s *State) IsCompleted(key string) bool { return s.completedSet[key] }

// Complete marks the exercise key as completed.
func (s *State) Complete(key string) {
	if s.completedSet[key] {
		return
	}
	s.completedSet[key] = true
	s.Completed = append(s.Completed, key)
}

// Uncomplete removes the exercise key from the completed set.
func (s *State) Uncomplete(key string) {
	if !s.completedSet[key] {
		return
	}
	delete(s.completedSet, key)
	out := s.Completed[:0:0]
	for _, k := range s.Completed {
		if k != key {
			out = append(out, k)
		}
	}
	s.Completed = out
}

// SetCurrent updates the current exercise.
func (s *State) SetCurrent(key string) { s.Current = key }

// DoneCount returns the number of completed exercises.
func (s *State) DoneCount() int { return len(s.Completed) }
