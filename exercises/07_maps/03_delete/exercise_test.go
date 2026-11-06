package main

import "testing"

func TestRemoveKey(t *testing.T) {
	m := map[string]int{"alice": 31, "bob": 27, "carol": 44}
	got := removeKey(m, "bob")

	if _, ok := got["bob"]; ok {
		t.Errorf(`removeKey(m, "bob"): "bob" is still in the map, want it deleted`)
	}
	if len(got) != 2 {
		t.Errorf(`len(removeKey(m, "bob")) = %d, want 2`, len(got))
	}
	if got["alice"] != 31 || got["carol"] != 44 {
		t.Errorf(`removeKey(m, "bob") changed other entries: %v`, got)
	}

	// Deleting a key that is not there must be a harmless no-op.
	got = removeKey(got, "nobody")
	if len(got) != 2 {
		t.Errorf("len after deleting a missing key = %d, want 2", len(got))
	}
}
