// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gerrors

import (
	"errors"
	"github.com/andrerrcosta2/gtools/core/domain/data"
)

type ErrorSeverity uint8

// NewErrorLevel creates a new ErrorLevel instance with the provided key and severity.
// It's a value object that represents the level of an error, with a key and a severity.
// The key is used to identify the error, and the severity is used to determine the importance of the error.
// The severity is a value of the ErrorSeverity type.
// The returned ErrorLevel is a value object, so it's safe to use the returned value directly.
func NewErrorLevel(key string, severity ErrorSeverity) ErrorLevel {
	return ErrorLevel{Key: key, Severity: severity}
}

type ErrorLevel struct {
	Key      string
	Severity ErrorSeverity
}

// String returns a string representation of the ErrorLevel.
// It simply returns the key as a string.
func (e ErrorLevel) String() string {
	return e.Key
}

// AsStackable checks if the error is of type StackableError and returns it.
// If the error isn't of type StackableError, it returns nil and false.
// Returns the error as StackableError and true if successful, nil and false otherwise.
func AsStackable(err error) (StackableError, bool) {
	// Declare a variable of type StackableError
	var stackableError StackableError

	// Check if the error is of type StackableError
	if !errors.As(err, &stackableError) {
		// If not, return nil and false
		return nil, false
	}

	// If the error is of type StackableError, return it and true
	return err.(StackableError), true
}

type StackableError interface {
	WrappedError
	// From creates a new StackableError from the given error.
	// It returns a new error with the given error as its underlying error.
	// If the error is already of type StackableError, it returns the error as is.
	// Otherwise, it returns a new error with the given error as its underlying error.
	From(err error) StackableError
	// IsEmpty checks if the error is empty.
	// An error is empty if it has no underlying errors.
	// It returns true if the error is empty, false otherwise.
	IsEmpty() bool
	// Len returns the number of underlying errors.
	// It returns 0 if the receiver has no underlying errors.
	Len() int
	// Output returns the error.
	// It returns nil if the receiver is empty.
	Output() error
	// Stack returns a new error with the given errors as its underlying errors.
	// The underlying errors are returned in the order they were passed to this method.
	// If no errors are passed, it returns the receiver itself.
	// If one error is passed, it returns the error as is.
	// If more than one error is passed, it returns a new error with the given errors as its underlying errors.
	Stack(err error)
	// Trace returns a string representation of the stack trace
	// of this error and all its underlying errors.
	// It's useful for debugging and logging.
	// It's the responsibility of the implementing type to
	// provide a meaningful implementation of this method.
	// If the type does not have a meaningful implementation,
	// it should return an empty string.
	Trace() string
	// Unstack returns a slice of errors representing the underlying errors.
	// It's the responsibility of the implementing type to
	// provide a meaningful implementation of this method.
	// If the type does not have a meaningful implementation,
	// it should return an empty slice.
	Unstack() []error
}

// AsWrapped checks if the given error is of type WrappedError.
// If it is, it returns the error as WrappedError and true.
// If it's not, it returns nil and false.
//
// Parameters:
// - err: The error to check.
//
// Returns:
// - WrappedError: The error as WrappedError if it's of type WrappedError.
// - Bool: True if the error is of type WrappedError, false otherwise.
func AsWrapped(err error) (WrappedError, bool) {
	// Declare a variable of type WrappedError
	var wrappedError WrappedError

	// Check if the error is of type WrappedError
	if !errors.As(err, &wrappedError) {
		// If not, return nil and false
		return nil, false
	}

	// If the error is of type WrappedError, return it and true
	return err.(WrappedError), true
}

type WrappedError interface {
	Cause() error
	// Error returns the error message of the error.
	// It's the responsibility of the implementing type to
	// provide a meaningful implementation of this method.
	// If the type does not have a meaningful implementation,
	// it should return an empty string.
	Error() string
	// Unwrap returns the underlying error.
	// It's the responsibility of the implementing type to
	// provide a meaningful implementation of this method.
	// If the type does not have a meaningful implementation,
	// it should return nil.
	Unwrap() error
}

// AsLeveled checks if the given error is of type LeveledError.
// If it is, it returns the error as LeveledError and true.
// If it's not, it returns nil and false.
//
// Parameters:
// - err: The error to check.
//
// Returns:
// - LeveledError: The error as LeveledError if it's of type LeveledError.
// - Bool: True if the error is of type LeveledError, false otherwise.
func AsLeveled(err error) (LeveledError, bool) {
	// Declare a variable of type LeveledError
	var leveledError LeveledError

	// Check if the error is of type LeveledError
	if !errors.As(err, &leveledError) {
		// If not, return nil and false
		return nil, false
	}

	// If the error is of type LeveledError, return it and true
	return err.(LeveledError), true
}

type LeveledError interface {
	// Error returns the error message of the leveled error.
	//
	// This method is part of the error interface.
	Error() string
	// Level returns the level of the leveled error.
	//
	// This method is part of the LeveledError interface.
	Level() ErrorLevel
	// Severity returns the severity of the leveled error.
	// The bigger the number, the more severe the error is.
	// This method is part of the LeveledError interface.
	Severity() ErrorSeverity
	// SeverityDiff returns the severity difference between the Error
	// instance and the provided LeveledError.
	// It returns an integer representing the severity difference.
	SeverityDiff(err LeveledError) int
}

type OperationalError interface {
	// Error returns the error message of the operational error.
	//
	// This method is part of the error interface.
	Error() string
	// Operation returns the op that caused the error.
	//
	// This method is part of the OperationalError interface.
	Operation() string
}

type TaggableError interface {
	data.Taggable[string]
	// Error returns the error message of the taggable error.
	//
	// This method is part of the error interface.
	Error() string
}
