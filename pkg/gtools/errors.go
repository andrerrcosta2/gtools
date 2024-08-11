// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtools

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type StackableError interface {
	Trace() string
	Error() string
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

type WrappedError interface {
	Error() string
	Unwrap() error
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

// FlattenError takes an error and returns a slice of errors that represent the stack trace of the error,
// as well as the last error in the stack trace.
//
// Parameters:
// - err: The error to flatten.
//
// Returns:
// - []error: a slice of errors representing the stack trace of the error.
// - Error: The last error in the stack trace.
func FlattenError(err error) ([]error, error) {
	unw := errors.Unwrap(err)
	if unw != nil {
		return FlattenError(unw)
	}

	var stk []error
	stk = append(stk, err)

	// Return the stack trace and the last error
	return stk, stk[len(stk)-1]
}

// ReadTrace takes a slice of errors and returns a formatted string representing the stack trace of those errors.
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
	sb.WriteString("Error trace:\n")

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

func StackErrors(currentStack []error, currentErr error, newErr error) ([]error, error) {
	if newErr != nil {
		tstk, terr := FlattenError(newErr)
		currentStack = append(currentStack, tstk...)
		return currentStack, terr
	}
	return currentStack, currentErr
}

type StackableErrorImpl struct {
	err error
	stk []error
}

func NewStackableError(err error) *StackableErrorImpl {
	return &StackableErrorImpl{
		err: err,
	}
}

func (s *StackableErrorImpl) Error() string {
	return s.err.Error()
}

func (s *StackableErrorImpl) Trace() string {
	return ReadTrace(s.stk)
}

func (s *StackableErrorImpl) Unwrap() error {
	return s.err
}

func (s *StackableErrorImpl) Len() int {
	return len(s.stk)
}

var _ error = (*StackableErrorImpl)(nil)
var _ StackableError = (*StackableErrorImpl)(nil)
var _ WrappedError = (*StackableErrorImpl)(nil)

type ConcurrentStackableError struct {
	mtx sync.RWMutex
	err error
	stk []error
}

func NewConcurrentStackableError(e error) *ConcurrentStackableError {
	err := &ConcurrentStackableError{
		err: e,
		stk: []error{},
		mtx: sync.RWMutex{},
	}
	if e != nil {
		err.stk, err.err = FlattenError(e)
	}
	return err
}

func (s *ConcurrentStackableError) Error() string {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.err.Error()
}

func (s *ConcurrentStackableError) Trace() string {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return ReadTrace(s.stk)
}

func (s *ConcurrentStackableError) Unwrap() error {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.err
}

func (s *ConcurrentStackableError) Stack(err ...error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	for _, e := range err {
		newStack, newErr := StackErrors(s.stk, s.err, e)

		// Update the internal stack
		if newErr != nil {
			s.err = newErr
		}
		s.stk = newStack
	}
}

func (s *ConcurrentStackableError) From(e error) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if e == nil {
		return nil
	}
	stk, err := FlattenError(e)
	if s.err == nil {
		return &ConcurrentStackableError{
			err: err,
			stk: stk,
		}
	}
	s.stk = append(s.stk, stk...)
	s.err = err
	return s
}

func (s *ConcurrentStackableError) Len() int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return len(s.stk)
}

func (s *ConcurrentStackableError) Empty() bool {
	s.mtx.RLock()
	s.mtx.RUnlock()
	return s.Len() == 0 && s.err == nil
}

func (s *ConcurrentStackableError) String() string {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return ReadTrace(s.stk)
}

var _ error = (*ConcurrentStackableError)(nil)
var _ StackableError = (*ConcurrentStackableError)(nil)
var _ WrappedError = (*ConcurrentStackableError)(nil)
