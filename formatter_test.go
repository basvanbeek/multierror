// Copyright (c) Bas van Beek 2024.

package multierror

import (
	"errors"
	"testing"
)

func TestLineFormatFuncSingle(t *testing.T) {
	expected := `foo`

	errors := []error{
		errors.New("foo"),
	}

	actual := LineErrorFormatFunc(errors)
	if actual != expected {
		t.Fatalf("bad: %#v", actual)
	}
}

func TestLineFormatFuncMultiple(t *testing.T) {
	expected := `2 errors occurred: foo; bar`

	errors := []error{
		errors.New("foo"),
		errors.New("bar"),
	}

	actual := LineErrorFormatFunc(errors)
	if actual != expected {
		t.Fatalf("bad: %#v", actual)
	}
}

func TestSetFormatter(t *testing.T) {
	var mErr *Error

	if want, have := (*Error)(nil), SetFormatter(mErr, ListFormatFunc); want != have {
		t.Errorf("bad: %#v", have)
	}
}
