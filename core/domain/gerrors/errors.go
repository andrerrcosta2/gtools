// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gerrors

import (
	"errors"
	"reflect"
	"sync"
)

const (
	NONE ErrorSeverity = iota
	FATAL
	ERROR
	OPERATION
	INPUT
	OUTPUT
	OS
	INTERNAL
	INHERIT
	EXTERNAL
	UNKNOWN
)

var (
	None          = ErrorLevel{Key: "None", Severity: NONE}
	Not_found_err = ErrorLevel{Key: "NotFound", Severity: INPUT}
	Operation_err = ErrorLevel{Key: "Operation", Severity: OPERATION}
	Input_err     = ErrorLevel{Key: "Input", Severity: INPUT}
	Output_err    = ErrorLevel{Key: "Output", Severity: OUTPUT}
	Internal_err  = ErrorLevel{Key: "Internal", Severity: INTERNAL}
	Inherit       = ErrorLevel{Key: "Inherit", Severity: INHERIT}
	External_err  = ErrorLevel{Key: "External", Severity: EXTERNAL}
	Unknown_err   = ErrorLevel{Key: "Unknown", Severity: UNKNOWN}
	Os_err        = ErrorLevel{Key: "Os", Severity: OS}
	Fatal_err     = ErrorLevel{Key: "Fatal", Severity: FATAL}
)

// NewError creates a new Error instance with the given error level and error.
// The function flattens the error using the FlattenError function and returns a pointer to the new Error instance.
//
// Parameters:
//   - lvl: the error level of the new Error instance.
//   - err: the error to be wrapped in the new Error instance.
//
// Returns:
//   - *Error: a pointer to the new Error instance.
func NewError(lvl ErrorLevel, err error) Error {
	// Flatten the error to get the stack trace and the top error
	stk, top, isw := FlattenError(err)

	if !isw {
		// Create a new Error instance with the given level, top error, and stack trace
		return Error{
			lvl: handleLevel(lvl, err),
			StackableError: &stackErr{
				err: err,
			},
		}
	}

	// Create a new Error instance with the given level, top error, and stack trace
	return Error{
		lvl: handleLevel(lvl, top),
		StackableError: &stackErr{
			err: top,
			stk: newStack(stk...),
		},
	}
}

func handleLevel(lvl ErrorLevel, err error) ErrorLevel {
	if lvl == Inherit {
		if lvlErr, ok := AsLeveled(err); ok {
			return lvlErr.Level()
		}
		return Unknown_err
	}
	return lvl
}

type Error struct {
	StackableError
	lvl ErrorLevel
}

// Stackable creates a new stackable error from the given error.
// It implements the gtools.StackableError interface.
// It returns a StackableError that can be used to stack errors.
// It isn't thread safe.
func Stackable(err error) StackableError {
	if err == nil {
		return &stackErr{}
	}
	return &stackErr{
		err: err,
		stk: newStack(err),
	}
}

func StackOf(err ...error) StackableError {
	if len(err) == 0 {
		return &stackErr{}
	}
	return &stackErr{
		err: err[0],
		stk: newStack(err...),
	}
}

func Stack() StackableError {
	return &stackErr{}
}

type stackErr struct {
	err error
	stk stack
}

func (e *stackErr) Cause() error {
	if e.IsEmpty() {
		return nil
	}
	return e.stk.Cause()
}

// Error returns the error message of the underlying error.
// It implements the error interface.
func (e *stackErr) Error() string {
	return e.err.Error()
}

func (e *stackErr) From(err error) StackableError {
	if err == nil {
		return e
	}

	if e.err == nil {
		if stk, peek, isw := FlattenError(err); isw {
			return &stackErr{
				err: peek,
				stk: newStack(stk...),
			}
		}
		return &stackErr{err: err}
	}

	if stk, _, isw := FlattenError(err); isw {
		return &stackErr{
			err: e.err,
			stk: newStack(stk...).Push(e.stk.stk...),
		}
	}

	return &stackErr{err: e.err, stk: newStack(err).Push(e.stk.stk...)}
}

func (e *stackErr) Is(target error) bool {
	// Check if the current error matches the target
	if errors.Is(e.err, target) {
		return true
	}

	// Check if any error in the stack matches the target
	for _, err := range e.stk.stk {
		if errors.Is(err, target) {
			return true
		}
	}

	return false
}

// IsEmpty checks if the error is empty.
// It implements the gtools.StackableError interface.
// It returns true if the error is empty, false otherwise.
func (e *stackErr) IsEmpty() bool {
	return e.err == nil && e.stk.IsEmpty()
}

func (e *stackErr) Output() error {
	if e.IsEmpty() {
		return nil
	}
	return e
}

// Stack appends the given error to the stack trace of the error and returns the new stackable error.
// It implements the gtools.StackableError interface.
// It takes an error and appends it to the stack trace of the error.
// It returns the new stackable error.
func (e *stackErr) Stack(err error) {
	if err == nil {
		return
	}
	// Flatten the given error to extract its stack trace
	if flat, fer, isw := FlattenError(e); isw {
		// Append the stack trace to the stack trace of the current error
		e.stk = e.stk.Push(flat...)
		e.err = fer
		return
	}

	// ToSet the underlying error to the flattened error
	e.stk = e.stk.Push(err)
	e.err = err
}

func (e *stackErr) String() string {
	return ReadTrace(e.stk.stk)
}

// Trace returns a string representing the stack trace of the error.
// It implements the gtools.StackableError interface.
// It returns a string that represents the stack trace of the error.
func (e *stackErr) Trace() string {
	// ReadTrace takes a slice of errors and returns a formatted string representing the stack trace of these errors.
	// It takes the stack trace of the error and returns a string representing it.
	return ReadTrace(e.stk.stk)
}

// Unwrap returns the underlying error.
// It implements the gerrors.Wrapper interface.
func (e *stackErr) Unwrap() error {
	if e.stk.IsEmpty() {
		return nil
	}

	// Create a copy of the stack without modifying the original
	stk, _ := e.stk.Pop()

	// Safely peek at the top element of the new stack
	peek := stk.Peek()
	if peek == nil {
		return nil
	}

	// Return a new stackErr with the updated stack
	return &stackErr{err: peek, stk: stk}
}

// Len returns the number of errors in the stack trace of the error.
// It implements the gtools.StackableError interface.
// It returns the number of errors in the stack trace of the error.
func (e *stackErr) Len() int {
	return e.stk.Len()
}

func (e *stackErr) Unstack() []error {
	return e.stk.stk
}

var _ error = (*stackErr)(nil)
var _ StackableError = (*stackErr)(nil)
var _ WrappedError = (*stackErr)(nil)

// ConcStackable returns a new concurrent stackable error from the given error.
// It implements the gtools.StackableError interface.
// It wraps the given error and provides a thread-safe stack of errors.
func ConcStackable(err error) StackableError {
	if err == nil {
		return &concStackErr{}
	}
	if flat, peek, isw := FlattenError(err); isw {
		return &concStackErr{
			stk: newStack(flat...),
			err: peek,
		}
	}
	return &concStackErr{err: err, stk: newStack(err)}
}

// ConcStack returns a new concurrent stackable error.
func ConcStack() StackableError {
	return &concStackErr{}
}

type concStackErr struct {
	mtx sync.RWMutex
	err error
	stk stack
}

func (e *concStackErr) Cause() error {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	if e.IsEmpty() {
		return nil
	}
	return e.stk.Cause()
}

func (e *concStackErr) Error() string {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	return e.err.Error()
}

func (e *concStackErr) From(err error) StackableError {
	e.mtx.Lock()
	defer e.mtx.Unlock()
	if err == nil {
		return e
	}

	if e.IsEmpty() {
		if stk, peek, isw := FlattenError(err); isw {
			return &concStackErr{
				err: peek,
				stk: newStack(stk...),
			}
		}
		return &concStackErr{err: err}
	}

	if stk, _, isw := FlattenError(err); isw {
		return &concStackErr{
			err: e.err,
			stk: newStack(stk...).Push(e.stk.stk...),
		}
	}

	return &concStackErr{err: e.err, stk: newStack(err).Push(e.stk.stk...)}
}

func (e *concStackErr) Is(target error) bool {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	// Check if the current error matches the target
	if errors.Is(e.err, target) {
		return true
	}

	// Check if any error in the stack matches the target
	for _, err := range e.stk.stk {
		if errors.Is(err, target) {
			return true
		}
	}

	return false
}

func (e *concStackErr) IsEmpty() bool {
	e.mtx.RLock()
	e.mtx.RUnlock()
	return e.Len() == 0 && e.err == nil
}

func (e *concStackErr) Len() int {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	return e.stk.Len()
}

func (e *concStackErr) Output() error {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	if e.IsEmpty() {
		return nil
	}
	return e
}

func (e *concStackErr) Stack(err error) {
	e.mtx.Lock()
	defer e.mtx.Unlock()
	if err == nil {
		return
	}
	// Flatten the given error to extract its stack trace
	if flat, peek, isw := FlattenError(err); isw {
		// Append the stack trace to the stack trace of the current error
		e.stk = e.stk.Push(flat...)
		e.err = peek
		return
	}

	// ToSet the underlying error to the flattened error
	e.stk = e.stk.Push(err)
	e.err = err
}

func (e *concStackErr) String() string {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	return ReadTrace(e.stk.stk)
}

func (e *concStackErr) Trace() string {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	return ReadTrace(e.stk.stk)
}

func (e *concStackErr) Unstack() []error {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	return e.stk.stk
}

func (e *concStackErr) Unwrap() error {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	return e.err // this is correct. the stackable error never store wrapped errors.
}

var _ error = (*concStackErr)(nil)
var _ StackableError = (*concStackErr)(nil)
var _ WrappedError = (*concStackErr)(nil)

// Operational creates an OperationalError from the given operation and error.
//
// An OperationalError is an error that is associated with an operation.
// It implements the gtools.OperationalError interface.
func Operational(op string, err error) OperationalError {
	return &opErr{
		op:  op,
		err: err,
	}
}

// opErr is an implementation of gtools.OperationalError interface.
//
// It contains the operation that caused the error as well as the error itself.
//
// The error returned by the Error method is the same as the error returned by the
// underlying error.
type opErr struct {
	op  string
	err error
}

// Error implements the error interface. It returns the string representation of
// the wrapped error.
func (e *opErr) Error() string {
	return e.err.Error()
}

// Operation returns the operation that caused the error.
func (e *opErr) Operation() string {
	return e.op
}

var _ error = (*opErr)(nil)

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
	var opError OperationalError

	// Check if the error is of type OperationalError
	if errors.As(err, &opError) {
		// If it is, return the error and true
		return opError, true
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
			// If it does, extract the value of the field and create a new opErr
			op := field.String()
			opError = &opErr{
				op:  op,
				err: err,
			}
			// Return the new OperationalError and true
			return opError, true
		}
	}

	// If the error is not of type OperationalError, return nil and false
	return nil, false
}

var _ OperationalError = (*opErr)(nil)

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
		e.tags[tag] = struct{}{} // ToSet the tag in the map
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
