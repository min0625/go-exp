// Package errorx provides extensions for error handling.
package errorx

import "errors"

// AsType finds the first error in err's chain that matches target, and if so, sets
// target to that error value and returns true. Otherwise, it returns false.
func AsType[E error](err error) (E, bool) {
	var e E

	if errors.As(err, &e) {
		return e, true
	}

	var zero E

	return zero, false
}

// Unwrapper is an interface for errors that can unwrap to a single error.
type Unwrapper interface {
	Unwrap() error
}

// MultiUnwrapper is an interface for errors that can unwrap to multiple errors.
type MultiUnwrapper interface {
	Unwrap() []error
}
