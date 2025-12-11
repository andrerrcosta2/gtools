// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gerrors

import (
	"errors"

	"github.com/andrerrcosta2/gtools/core/domain/data"
)

type Severity uint8

// NewLevel creates a new Level instance with the provided key and severity.
// It's a value object that represents the level of an error, with a key and a severity.
// The key is used to identify the error, and the severity is used to determine the importance of the error.
// The severity is a value of the Severity type.
// The returned Level is a value object, so it's safe to use the returned value directly.
func NewLevel(key string, severity Severity) Level {
	return Level{Key: key, Severity: severity}
}

type Level struct {
	Key      string
	Severity Severity
}

// String returns a string representation of the Level.
// It simply returns the key as a string.
func (e Level) String() string {
	return e.Key
}

// AsStackable checks if the error is of type Stackable and returns it.
// If the error isn't of type Stackable, it returns nil and false.
// Returns the error as Stackable and true if successful, nil and false otherwise.
func AsStackable(err error) (Stackable, bool) {
	// Declare a variable of type Stackable
	var stackableError Stackable

	// Check if the error is of type Stackable
	if !errors.As(err, &stackableError) {
		// If not, return nil and false
		return nil, false
	}

	// If the error is of type Stackable, return it and true
	return err.(Stackable), true
}

type Stackable interface {
	Wrapped
	// From creates a new Stackable from the given error.
	// It returns a new error with the given error as its underlying error.
	// If the error is already of type Stackable, it returns the error as is.
	// Otherwise, it returns a new error with the given error as its underlying error.
	From(err error) Stackable
	// Is reports whether the error is the same as the given error
	Is(err error) bool
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
}

// AsWrapped checks if the given error is of type Wrapped.
// If it is, it returns the error as Wrapped and true.
// If it's not, it returns nil and false.
//
// Parameters:
// - err: The error to check.
//
// Returns:
// - Wrapped: The error as Wrapped if it's of type Wrapped.
// - Bool: True if the error is of type Wrapped, false otherwise.
func AsWrapped(err error) (Wrapped, bool) {
	// Declare a variable of type Wrapped
	var wrappedError Wrapped

	// Check if the error is of type Wrapped
	if !errors.As(err, &wrappedError) {
		// If not, return nil and false
		return nil, false
	}

	// If the error is of type Wrapped, return it and true
	return err.(Wrapped), true
}

type Wrapped interface {
	error
	Cause() error
	// Unwrap returns the underlying errors.
	// It's the responsibility of the implementing type to
	// provide a meaningful implementation of this method.
	// If the type does not have a meaningful implementation,
	// it should return nil.
	Unwrap() []error
}

// AsLeveled checks if the given error is of type Leveled.
// If it is, it returns the error as Leveled and true.
// If it's not, it returns nil and false.
//
// Parameters:
// - err: The error to check.
//
// Returns:
// - Leveled: The error as Leveled if it's of type Leveled.
// - Bool: True if the error is of type Leveled, false otherwise.
func AsLeveled(err error) (Leveled, bool) {
	// Declare a variable of type Leveled
	var leveledError Leveled

	// Check if the error is of type Leveled
	if !errors.As(err, &leveledError) {
		// If not, return nil and false
		return nil, false
	}

	// If the error is of type Leveled, return it and true
	return err.(Leveled), true
}

type Leveled interface {
	// Error returns the error message of the leveled error.
	//
	// This method is part of the error interface.
	Error() string
	// Level returns the level of the leveled error.
	//
	// This method is part of the Leveled interface.
	Level() Level
	// Severity returns the severity of the leveled error.
	// The bigger the number, the more severe the error is.
	// This method is part of the Leveled interface.
	Severity() Severity
	// SeverityDiff returns the severity difference between the defaultErr
	// instance and the provided Leveled.
	// It returns an integer representing the severity difference.
	SeverityDiff(err Leveled) int
}

type Operational interface {
	// Error returns the error message of the operational error.
	//
	// This method is part of the error interface.
	Error() string
	// Operation returns the op that caused the error.
	//
	// This method is part of the Operational interface.
	Operation() string
}

type Taggable interface {
	data.Taggable[string]
	// Error returns the error message of the taggable error.
	//
	// This method is part of the error interface.
	Error() string
}
