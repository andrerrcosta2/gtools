// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gerrors

import (
	"errors"
)

// IsErrorOfAny checks if err is one of the errors in targets.
//
// It wraps errors.Is, but it's more convenient to use when you need to check against multiple errors.
func IsErrorOfAny(err error, targets ...error) bool {
	// Iterate over the targets and check if err is one of them using errors.Is.
	for _, target := range targets {
		if errors.Is(err, target) {
			return true
		}
	}
	// If err is not one of the targets, return false.
	return false
}

// AllErrorsOf checks if all the errors in the stack of err are one of the targets.
//
// It takes a gtools.Stackable and a variable number of targets as arguments.
//
// It returns true if all the errors in the stack of err are one of the targets, false otherwise.
func AllErrorsOf(err error, targets ...error) bool {
	if se, ok := AsStackable(err); ok {
		// Iterate over the slice of errors and check if each error is one of the targets.
		// If any error is not one of the targets, return false.
		for _, e := range se.Unwrap() {
			if !IsErrorOfAny(e, targets...) {
				return false
			}
		}

		// If all errors are one of the targets, return true.
		return true
	}

	flat := FlattenError(err)

	for _, e := range flat {
		if !IsErrorOfAny(e, targets...) {
			return false
		}
	}
	return true
}

// ContainsError checks if any of the errors in the stack of err is one of the targets.
//
// It takes a gtools.Stackable and a variable number of targets as arguments.
//
// It returns true if any of the errors in the stack of err is one of the targets, false otherwise.
func ContainsError(err error, targets ...error) bool {
	// Iterate over the slice of errors and check if each error is one of the targets.
	// If any error is one of the targets, return true.
	if se, ok := AsStackable(err); ok {
		for _, target := range targets {
			for _, e := range se.Unwrap() {
				if errors.Is(e, target) {
					return true
				}
			}
		}
		return false
	}

	flat := FlattenError(err)
	for _, e := range flat {
		for _, target := range targets {
			if errors.Is(e, target) {
				return true
			}
		}
	}
	return false
}
