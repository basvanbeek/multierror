// Copyright (c) Bas van Beek 2024.

package multierror

import (
	"fmt"
	"strings"
)

// SetFormatter if provided with a multierror will update the multierror
// serialization function with the provided ErrorFormatFunc.
//
// This method is not safe to be called concurrently.
func SetFormatter(err error, fn ErrorFormatFunc) error {
	if e, ok := err.(*Error); ok {
		// check for typed nil
		if e != nil {
			e.ErrorFormat = fn
		}
	}
	return err
}

// LineErrorFormatFunc is a basic formatter that outputs the number of errors
// that occurred along with all errors on a single line.
//
// Typically, this format is used for logs.
func LineErrorFormatFunc(es []error) string {
	if len(es) == 1 {
		return es[0].Error()
	}

	var b strings.Builder
	for _, err := range es {
		b.WriteString(err.Error())
		b.WriteString("; ")
	}
	return fmt.Sprintf("%d errors occurred: %s", len(es), b.String()[:b.Len()-2])
}
