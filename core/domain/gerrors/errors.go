// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gerrors

import (
	"errors"
	"reflect"
	"sync"
)

// TODO: some errors are implementing more methods they declare on the interface.
// must also define a if they will retrieve its own interface or just an error
// TODO: do a full review on that package.

// Stackable creates a new stackable error from the given error.
// It implements the gtools.StackableError interface.
// It returns a StackableError that can be used to stack errors.
// It isn't thread safe.
func Stackable(err error) StackableError {
	return &stackableError{
		err: err,
	}
}

type stackableError struct {
	err error
	stk []error
}

// Error returns the error message of the underlying error.
// It implements the error interface.
func (s *stackableError) Error() string {
	return s.err.Error()
}

func (s *stackableError) From(e error) StackableError {
	if e == nil {
		return nil
	}
	stk, err := FlattenError(e)
	if s.err == nil {
		return &stackableError{
			err: err,
			stk: stk,
		}
	}
	s.stk = append(s.stk, stk...)
	s.err = err
	return s
}

// IsEmpty checks if the error is empty.
// It implements the gtools.StackableError interface.
// It returns true if the error is empty, false otherwise.
func (s *stackableError) IsEmpty() bool {
	return s.err == nil
}

// Trace returns a string representing the stack trace of the error.
// It implements the gtools.StackableError interface.
// It returns a string that represents the stack trace of the error.
func (s *stackableError) Trace() string {
	// ReadTrace takes a slice of errors and returns a formatted string representing the stack trace of these errors.
	// It takes the stack trace of the error and returns a string representing it.
	return ReadTrace(s.stk)
}

// Unwrap returns the underlying error.
// It implements the xerrors.Wrapper interface.
func (s *stackableError) Unwrap() error {
	return s.err
}

// Len returns the number of errors in the stack trace of the error.
// It implements the gtools.StackableError interface.
// It returns the number of errors in the stack trace of the error.
func (s *stackableError) Len() int {
	return len(s.stk)
}

// Stack appends the given error to the stack trace of the error and returns the new stackable error.
// It implements the gtools.StackableError interface.
// It takes an error and appends it to the stack trace of the error.
// It returns the new stackable error.
func (s *stackableError) Stack(e error) {
	if e == nil {
		return
	}
	// Flatten the given error to extract its stack trace
	stk, err := FlattenError(e)
	// Append the stack trace to the stack trace of the current error
	s.stk = append(s.stk, stk...)
	// Set the underlying error to the flattened error
	s.err = err
}

func (s *stackableError) Unstack() []error {
	return s.stk
}

var _ error = (*stackableError)(nil)
var _ StackableError = (*stackableError)(nil)
var _ WrappedError = (*stackableError)(nil)

// ConcurrentStackable returns a new concurrent stackable error from the given error.
// It implements the gtools.StackableError interface.
// It wraps the given error and provides a thread-safe stack of errors.
func ConcurrentStackable(e error) StackableError {
	// Create a new concurrent stackable error
	err := &concurrentStackableError{
		err: e,
		stk: []error{},
		mtx: sync.RWMutex{},
	}
	// If the error is not nil, flatten it and store it in the stack
	if e != nil {
		err.stk, err.err = FlattenError(e)
	}
	// Return the new concurrent stackable error
	return err
}

type concurrentStackableError struct {
	mtx sync.RWMutex
	err error
	stk []error
}

func (s *concurrentStackableError) Error() string {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.err.Error()
}

func (s *concurrentStackableError) Trace() string {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return ReadTrace(s.stk)
}

func (s *concurrentStackableError) Unwrap() error {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.err
}

func (s *concurrentStackableError) Stack(e error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if e == nil {
		return
	}
	// Flatten the given error to extract its stack trace
	stk, err := FlattenError(e)
	// Append the stack trace to the stack trace of the current error
	s.stk = append(s.stk, stk...)
	// Set the underlying error to the flattened error
	s.err = err
}

func (s *concurrentStackableError) From(e error) StackableError {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if e == nil {
		return nil
	}
	stk, err := FlattenError(e)
	if s.err == nil {
		return &concurrentStackableError{
			err: err,
			stk: stk,
		}
	}
	s.stk = append(s.stk, stk...)
	s.err = err
	return s
}

func (s *concurrentStackableError) Len() int {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return len(s.stk)
}

func (s *concurrentStackableError) IsEmpty() bool {
	s.mtx.RLock()
	s.mtx.RUnlock()
	return s.Len() == 0 && s.err == nil
}

func (s *concurrentStackableError) String() string {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return ReadTrace(s.stk)
}

func (s *concurrentStackableError) Return() error {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if s.IsEmpty() {
		return nil
	}
	return s
}

func (s *concurrentStackableError) Unstack() []error {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	return s.stk
}

var _ error = (*concurrentStackableError)(nil)
var _ StackableError = (*concurrentStackableError)(nil)
var _ WrappedError = (*concurrentStackableError)(nil)

// Operational creates an OperationalError from the given operation and error.
//
// An OperationalError is an error that is associated with an operation.
// It implements the gtools.OperationalError interface.
func Operational(op string, err error) OperationalError {
	return &operationalError{
		op:  op,
		err: err,
	}
}

// operationalError is an implementation of gtools.OperationalError interface.
//
// It contains the operation that caused the error as well as the error itself.
//
// The error returned by the Error method is the same as the error returned by the
// underlying error.
type operationalError struct {
	op  string
	err error
}

// Error implements the error interface. It returns the string representation of
// the wrapped error.
func (e *operationalError) Error() string {
	return e.err.Error()
}

// Operation returns the operation that caused the error.
func (e *operationalError) Operation() string {
	return e.op
}

var _ error = (*operationalError)(nil)

// AsOperational is a function that checks if an error is an OperationalError.
//
// Parameters:
// - err: The error to check.
//
// Returns:
// - OperationalError: The error as OperationalError if it's of type OperationalError.
// - bool: True if the error is of type OperationalError, false otherwise.
func AsOperational(err error) (OperationalError, bool) {
	// Declare a variable of type OperationalError
	var opErr OperationalError

	// Check if the error is of type OperationalError
	if errors.As(err, &opErr) {
		// If it is, return the error and true
		return opErr, true
	}

	// If the error is not of type OperationalError, check if it is a pointer to a struct
	val := reflect.ValueOf(err)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() == reflect.Struct {
		// Check if the struct has a field named "Op" of type string
		field := val.FieldByName("Op")
		if field.IsValid() && field.Kind() == reflect.String {
			// If it does, extract the value of the field and create a new operationalError
			op := field.String()
			opErr = &operationalError{
				op:  op,
				err: err,
			}
			// Return the new OperationalError and true
			return opErr, true
		}
	}

	// If the error is not of type OperationalError, return nil and false
	return nil, false
}

var _ OperationalError = (*operationalError)(nil)

// Tagged creates a new taggableError from the given error and tags.
//
// The given tags are stored in the returned error, and can be retrieved
// using the Tags() method.
//
// The returned error is of type gtools.TaggableError.
func Tagged(err error, tags ...string) TaggableError {
	taggable := &taggableError{
		tags:  map[string]struct{}{},
		error: err,
	}
	taggable.Tag(tags...)
	return taggable
}

type taggableError struct {
	tags  map[string]struct{}
	error error
}

// Error returns the error message of the underlying error.
//
// This method is part of the error interface.
func (e *taggableError) Error() string {
	return e.error.Error()
}

// Tag sets the given tags on the error.
//
// The tags are stored in the error and can be retrieved
// using the Tags() method.
func (e *taggableError) Tag(tags ...string) {
	for _, tag := range tags {
		e.tags[tag] = struct{}{} // Set the tag in the map
	}
}

// Tags returns all the tags that have been set on this error.
//
// The returned slice is a copy of the internal tags map keys.
func (e *taggableError) Tags() []string {
	tags := make([]string, 0, len(e.tags))
	for tag := range e.tags {
		tags = append(tags, tag)
	}
	return tags
}

var _ error = (*taggableError)(nil)
