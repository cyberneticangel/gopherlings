// Package exercise gives typed access to the exercises and solutions that
// are embedded in the gopherlings binary.
package exercise

import (
	"fmt"
	"io/fs"
	"path"
	"sort"
	"strings"

	"github.com/cyberneticangel/gopherlings/internal/manifest"
)

// Exercise is a manifest entry together with the files that make it up,
// as they exist in the embedded (pristine) copy.
type Exercise struct {
	manifest.Exercise
	// Files maps paths relative to the exercise directory (e.g.
	// "exercise.go", "exercise_test.go") to their contents.
	Files map[string]string
}

// Store provides read access to the embedded exercises and solutions.
type Store struct {
	man       *manifest.Manifest
	exercises fs.FS // rooted at exercises/
	solutions fs.FS // rooted at solutions/
}

// NewStore builds a Store from a filesystem rooted at the directory that
// contains the exercises/ and solutions/ trees.
func NewStore(root fs.FS) (*Store, error) {
	exFS, err := fs.Sub(root, "exercises")
	if err != nil {
		return nil, fmt.Errorf("locate embedded exercises: %w", err)
	}
	solFS, err := fs.Sub(root, "solutions")
	if err != nil {
		return nil, fmt.Errorf("locate embedded solutions: %w", err)
	}
	man, err := manifest.Load(exFS)
	if err != nil {
		return nil, err
	}
	s := &Store{man: man, exercises: exFS, solutions: solFS}
	for _, e := range man.Exercises {
		if _, err := fs.Stat(exFS, e.Dir()); err != nil {
			return nil, fmt.Errorf("manifest references missing exercise dir %q: %w", e.Dir(), err)
		}
	}
	return s, nil
}

// Manifest returns the underlying manifest.
func (s *Store) Manifest() *manifest.Manifest { return s.man }

// Ordered returns the exercises in curriculum order.
func (s *Store) Ordered() []manifest.Exercise { return s.man.Exercises }

// Get returns the pristine files of the exercise with the given key.
func (s *Store) Get(key string) (Exercise, error) {
	e, ok := s.man.Find(key)
	if !ok {
		return Exercise{}, fmt.Errorf("unknown exercise %q", key)
	}
	files, err := readDirFiles(s.exercises, e.Dir())
	if err != nil {
		return Exercise{}, fmt.Errorf("read exercise %q: %w", key, err)
	}
	return Exercise{Exercise: e, Files: files}, nil
}

// SolutionFiles returns the solved (non-test) files for an exercise.
func (s *Store) SolutionFiles(key string) (map[string]string, error) {
	e, ok := s.man.Find(key)
	if !ok {
		return nil, fmt.Errorf("unknown exercise %q", key)
	}
	return readDirFiles(s.solutions, e.Dir())
}

// Find resolves a user-supplied name to an exercise. It accepts a full key
// ("00_intro/01_hello"), an exercise ID ("01_hello"), or a unique prefix of
// either ("01_he").
func (s *Store) Find(query string) (manifest.Exercise, error) {
	q := strings.TrimSpace(query)
	q = strings.TrimPrefix(q, "./")
	q = strings.TrimPrefix(q, "exercises/")
	q = strings.TrimSuffix(q, "/")
	if q == "" {
		return manifest.Exercise{}, fmt.Errorf("empty exercise name")
	}
	if e, ok := s.man.Find(q); ok {
		return e, nil
	}
	// Suffix match on key ("intro/01_hello" or "01_hello").
	var matches []manifest.Exercise
	for _, e := range s.man.Exercises {
		k := e.Key()
		if strings.HasSuffix(k, "/"+q) || k == q {
			matches = append(matches, e)
		}
	}
	if len(matches) == 0 {
		// Prefix match on ID.
		for _, e := range s.man.Exercises {
			if strings.HasPrefix(e.ID, q) {
				matches = append(matches, e)
			}
		}
	}
	switch len(matches) {
	case 1:
		return matches[0], nil
	case 0:
		return manifest.Exercise{}, fmt.Errorf("no exercise matches %q (see `gopherlings list`)", query)
	default:
		names := make([]string, 0, len(matches))
		for _, m := range matches {
			names = append(names, m.Key())
		}
		return manifest.Exercise{}, fmt.Errorf("%q is ambiguous: %s", query, strings.Join(names, ", "))
	}
}

// readDirFiles reads every regular file in dir (non-recursively) of fsys.
func readDirFiles(fsys fs.FS, dir string) (map[string]string, error) {
	entries, err := fs.ReadDir(fsys, dir)
	if err != nil {
		return nil, err
	}
	files := make(map[string]string, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		data, err := fs.ReadFile(fsys, path.Join(dir, entry.Name()))
		if err != nil {
			return nil, err
		}
		files[entry.Name()] = string(data)
	}
	return files, nil
}

// SortKeys is a small helper used by tests.
func SortKeys(m map[string]string) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
