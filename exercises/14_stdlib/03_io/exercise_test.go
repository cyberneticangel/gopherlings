package main

import (
	"errors"
	"strings"
	"testing"
)

// errBoom is the error that errReader always fails with.
var errBoom = errors.New("boom")

// errReader is an io.Reader whose Read always fails.
type errReader struct{}

func (errReader) Read([]byte) (int, error) {
	return 0, errBoom
}

func TestCountBytes(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{"gopher", 6},
		{"", 0},
		{"gopher gopher", 13},
	}
	for _, c := range cases {
		n, err := countBytes(strings.NewReader(c.input))
		if err != nil {
			t.Errorf("countBytes(%q) returned error %v, want nil", c.input, err)
			continue
		}
		if n != c.want {
			t.Errorf("countBytes(%q) = %d bytes, want %d", c.input, n, c.want)
		}
	}
}

func TestCountBytesPropagatesError(t *testing.T) {
	_, err := countBytes(errReader{})
	if !errors.Is(err, errBoom) {
		t.Errorf("countBytes(errReader{}) error = %v, want the reader's own error (%v)", err, errBoom)
	}
}
