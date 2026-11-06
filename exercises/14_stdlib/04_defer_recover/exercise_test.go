package main

import "testing"

func TestCatchPanic(t *testing.T) {
	if !catchPanic(func() { panic("boom") }) {
		t.Error(`catchPanic(func() { panic("boom") }) = false, want true`)
	}
}

func TestCatchPanicNoPanic(t *testing.T) {
	if catchPanic(func() {}) {
		t.Error("catchPanic(func() {}) = true, want false")
	}
}
