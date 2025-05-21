// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gerrors

import (
	"errors"
	"fmt"
	"strings"
)

// FlattenError takes an error and returns a slice of errors that represent the stack trace of the error,
// as well as the last error in the stack trace.
//
// Parameters:
// - err: The error to flatten.
//
// Returns:
// - []error: a slice of errors representing the stack trace of the error.
// - Error: The last error in the stack trace.
func FlattenError(err error) (stk []error, peek error, isWrapped bool) {
	if err == nil {
		return
	}

	errs := []error{err}
	for err := errors.Unwrap(err); err != nil; err = errors.Unwrap(err) {
		errs = append(errs, err)
	}

	for i, j := 0, len(errs)-1; i < j; i, j = i+1, j-1 {
		errs[i], errs[j] = errs[j], errs[i]
	}
	return errs, err, len(errs) > 1
}

// ReadTrace takes a slice of errors and returns a formatted string representing the stack trace of these errors.
//
// Parameters:
// - stack: a slice of errors representing the stack trace.
//
// Returns:
// - string: a formatted string representing the stack trace.
func ReadTrace(stack []error) string {
	// Initialize a strings.Builder to hold the formatted string
	var sb strings.Builder

	// Write the header of the error trace to the strings.Builder
	sb.WriteString("error trace:\n")

	// Loop through each error in the stack trace
	for i, err := range stack {
		// Check if the current error is not nil
		if err != nil {
			// Write the current error to the strings.Builder with its index in the stack trace
			sb.WriteString(fmt.Sprintf(" %d: %v\n", i+1, err))
		}
	}

	// Return the formatted string representing the stack trace
	return sb.String()
}
