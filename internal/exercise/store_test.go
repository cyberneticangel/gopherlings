package exercise

import (
	"testing"
	"testing/fstest"
)

const testManifest = `{
  "sections": [
    {"id": "00_intro", "title": "Intro"},
    {"id": "01_vars", "title": "Variables"}
  ],
  "exercises": [
    {"id": "01_hello", "section": "00_intro", "name": "Hello", "kind": "compile"},
    {"id": "02_world", "section": "00_intro", "name": "World", "kind": "test"},
    {"id": "01_vars", "section": "01_vars", "name": "Vars", "kind": "test"}
  ]
}`

func mapFS() fstest.MapFS {
	return fstest.MapFS{
		"exercises/manifest.json":                      &fstest.MapFile{Data: []byte(testManifest)},
		"exercises/00_intro/01_hello/exercise.go":      &fstest.MapFile{Data: []byte("package main\n")},
		"exercises/00_intro/01_hello/exercise_test.go": &fstest.MapFile{Data: []byte("package main\n")},
		"exercises/00_intro/02_world/exercise.go":      &fstest.MapFile{Data: []byte("package main\n")},
		"exercises/00_intro/02_world/exercise_test.go": &fstest.MapFile{Data: []byte("package main\n")},
		"exercises/01_vars/01_vars/exercise.go":        &fstest.MapFile{Data: []byte("package main\n")},
		"exercises/01_vars/01_vars/exercise_test.go":   &fstest.MapFile{Data: []byte("package main\n")},
		"solutions/00_intro/01_hello/exercise.go":      &fstest.MapFile{Data: []byte("package main\n")},
	}
}

func TestNewStoreValidation(t *testing.T) {
	if _, err := NewStore(mapFS()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNewStoreMissingDir(t *testing.T) {
	fsys := mapFS()
	delete(fsys, "exercises/01_vars/01_vars/exercise.go")
	delete(fsys, "exercises/01_vars/01_vars/exercise_test.go")
	if _, err := NewStore(fsys); err == nil {
		t.Fatal("expected error for missing exercise dir")
	}
}

func TestFind(t *testing.T) {
	store, err := NewStore(mapFS())
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		query string
		want  string
	}{
		{"00_intro/01_hello", "00_intro/01_hello"},
		{"01_hello", "00_intro/01_hello"},
		{"01_vars", "01_vars/01_vars"}, // suffix match wins over prefix
		{"01_he", "00_intro/01_hello"}, // unique prefix
		{"./01_hello", "00_intro/01_hello"},
		{"exercises/01_hello", "00_intro/01_hello"},
	}
	for _, c := range cases {
		e, err := store.Find(c.query)
		if err != nil {
			t.Errorf("Find(%q): %v", c.query, err)
			continue
		}
		if e.Key() != c.want {
			t.Errorf("Find(%q) = %q, want %q", c.query, e.Key(), c.want)
		}
	}
	if _, err := store.Find("zzz"); err == nil {
		t.Error("expected error for unknown exercise")
	}
	if _, err := store.Find(""); err == nil {
		t.Error("expected error for empty query")
	}
}

func TestGetAndSolutionFiles(t *testing.T) {
	store, err := NewStore(mapFS())
	if err != nil {
		t.Fatal(err)
	}
	ex, err := store.Get("00_intro/01_hello")
	if err != nil {
		t.Fatal(err)
	}
	if len(ex.Files) != 2 {
		t.Errorf("want 2 files, got %d", len(ex.Files))
	}
	sol, err := store.SolutionFiles("00_intro/01_hello")
	if err != nil {
		t.Fatal(err)
	}
	if len(sol) != 1 {
		t.Errorf("want 1 solution file, got %d", len(sol))
	}
	if _, err := store.SolutionFiles("01_vars/01_vars"); err == nil {
		t.Error("expected error for missing solution")
	}
}
