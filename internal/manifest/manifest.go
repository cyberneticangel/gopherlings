// Package manifest defines the metadata that describes every exercise.
// The manifest is authored once, embedded into the gopherlings binary, and
// is the single source of truth for exercise ordering, names, and hints.
package manifest

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"sort"
)

// Kind describes how an exercise is primarily meant to be solved. It only
// affects the wording of the UI: verification is identical for both kinds
// (compile errors and failing tests both surface through `go test -json`).
type Kind string

const (
	// KindCompile exercises start with code that does not compile.
	KindCompile Kind = "compile"
	// KindTest exercises compile but fail their tests until completed.
	KindTest Kind = "test"
)

// Exercise is one step in the curriculum.
type Exercise struct {
	// ID is unique within its section, e.g. "01_hello".
	ID string `json:"id"`
	// Section is the section ID the exercise belongs to, e.g. "00_intro".
	Section string `json:"section"`
	// Name is a human-readable title, e.g. "Hello, Gopher".
	Name string `json:"name"`
	// Kind affects UI copy only.
	Kind Kind `json:"kind"`
	// Blurb is a one-line description shown in `gopherlings list`.
	Blurb string `json:"blurb,omitempty"`
	// Hints are shown progressively by `gopherlings hint`.
	Hints []string `json:"hints,omitempty"`
}

// Key is the canonical, unique identifier of an exercise: "<section>/<id>".
func (e Exercise) Key() string { return e.Section + "/" + e.ID }

// Dir is the path of the exercise directory relative to the project's
// exercises/ folder, e.g. "00_intro/01_hello".
func (e Exercise) Dir() string { return e.Key() }

// Section groups exercises into a chapter of the curriculum.
type Section struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Blurb string `json:"blurb,omitempty"`
}

// Manifest is the full curriculum definition.
type Manifest struct {
	Sections  []Section  `json:"sections"`
	Exercises []Exercise `json:"exercises"`
}

// Load reads and validates manifest.json from fsys (rooted at exercises/).
func Load(fsys fs.FS) (*Manifest, error) {
	data, err := fs.ReadFile(fsys, "manifest.json")
	if err != nil {
		return nil, fmt.Errorf("read manifest: %w", err)
	}
	var m Manifest
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("parse manifest.json: %w", err)
	}
	if err := m.Validate(); err != nil {
		return nil, err
	}
	return &m, nil
}

// Validate checks internal consistency: unique keys, known sections, valid kinds.
func (m *Manifest) Validate() error {
	if len(m.Exercises) == 0 {
		return fmt.Errorf("manifest contains no exercises")
	}
	sectionIDs := make(map[string]bool, len(m.Sections))
	for _, s := range m.Sections {
		if s.ID == "" || s.Title == "" {
			return fmt.Errorf("section %q: id and title are required", s.ID)
		}
		if sectionIDs[s.ID] {
			return fmt.Errorf("duplicate section %q", s.ID)
		}
		sectionIDs[s.ID] = true
	}
	seen := make(map[string]bool, len(m.Exercises))
	for _, e := range m.Exercises {
		k := e.Key()
		if e.ID == "" || e.Section == "" || e.Name == "" {
			return fmt.Errorf("exercise %q: id, section and name are required", k)
		}
		if seen[k] {
			return fmt.Errorf("duplicate exercise %q", k)
		}
		seen[k] = true
		if !sectionIDs[e.Section] {
			return fmt.Errorf("exercise %q references unknown section %q", k, e.Section)
		}
		if e.Kind != KindCompile && e.Kind != KindTest {
			return fmt.Errorf("exercise %q: unknown kind %q (want %q or %q)", k, e.Kind, KindCompile, KindTest)
		}
	}
	return nil
}

// Find returns the exercise with the given key.
func (m *Manifest) Find(key string) (Exercise, bool) {
	for _, e := range m.Exercises {
		if e.Key() == key {
			return e, true
		}
	}
	return Exercise{}, false
}

// SectionOf returns the section containing the given exercise.
func (m *Manifest) SectionOf(e Exercise) (Section, bool) {
	for _, s := range m.Sections {
		if s.ID == e.Section {
			return s, true
		}
	}
	return Section{}, false
}

// SortedSections returns sections in ID order.
func (m *Manifest) SortedSections() []Section {
	out := append([]Section(nil), m.Sections...)
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}
