package main

import (
	"errors"
	"testing"
)

func TestValidateAge(t *testing.T) {
	if err := validateAge(30); err != nil {
		t.Errorf("validateAge(30) = %v, want nil", err)
	}
	if err := validateAge(0); err != nil {
		t.Errorf("validateAge(0) = %v, want nil", err)
	}

	err := validateAge(-3)
	if err == nil {
		t.Fatal("validateAge(-3) returned no error, want a *ValidationError")
	}
	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Fatalf("validateAge(-3) error = %v (type %T), want errors.As(err, &ve) = true", err, err)
	}
	if ve.Field != "age" {
		t.Errorf("validation error field = %q, want %q", ve.Field, "age")
	}
}
