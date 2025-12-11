// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gerrors

import (
	"errors"
	"reflect"
	"sync"
)

const (
	NONE Severity = iota
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
	None          = Level{Key: "None", Severity: NONE}
	Not_found_err = Level{Key: "NotFound", Severity: INPUT}
	Operation_err = Level{Key: "Operation", Severity: OPERATION}
	Input_err     = Level{Key: "Input", Severity: INPUT}
	Output_err    = Level{Key: "Output", Severity: OUTPUT}
	Internal_err  = Level{Key: "Internal", Severity: INTERNAL}
	Inherit       = Level{Key: "Inherit", Severity: INHERIT}
	External_err  = Level{Key: "External", Severity: EXTERNAL}
	Unknown_err   = Level{Key: "Unknown", Severity: UNKNOWN}
	Os_err        = Level{Key: "Os", Severity: OS}
	Fatal_err     = Level{Key: "Fatal", Severity: FATAL}
)

// New creates a new defaultErr instance with the given error level and error.
// The function flattens the error using the FlattenError function and returns a pointer to the new defaultErr instance.
//
// Parameters:
//   - lvl: the error level of the new defaultErr instance.
//   - err: the error to be wrapped in the new defaultErr instance.
//
// Returns:
//   - *defaultErr: a pointer to the new defaultErr instance.
func New(lvl Level, err error) error {
	// Flatten the error to get the stack trace and the top error
	stk := stack{FlattenError(err)}

	// Create a new defaultErr instance with the given level, top error, and stack trace
	return &defaultErr{
		lvl: handleLevel(lvl, stk.Peek()),
		Stackable: &stackErr{
			stk: stk,
		},
	}
}

func handleLevel(lvl Level, err error) Level {
	if lvl == Inherit {
		if lvlErr, ok := AsLeveled(err); ok {
			return lvlErr.Level()
		}
		return None
	}
	return lvl
}

type defaultErr struct {
	Stackable
	lvl Level
}

// StackableOf creates a new stackable error from the given error.
// It implements the gtools.Stackable interface.
// It returns a Stackable that can be used to stack errors.
// It isn't thread safe.
func StackableOf(err error) Stackable {
	if err == nil {
		return &stackErr{}
	}
	return &stackErr{
		stk: newStack(FlattenError(err)...),
	}
}

func StackOf(errs ...error) Stackable {
	if len(errs) == 0 {
		return &stackErr{}
	}
	var out []error
	for _, err := range errs {
		flat := FlattenError(err)
		out = append(out, flat...)
	}
	return &stackErr{
		stk: newStack(out...),
	}
}

func Stack() Stackable {
	return &stackErr{}
}

type stackErr struct {
	stk stack
}

func (e *stackErr) Cause() error {
	if e.stk.Len() == 0 {
		return nil
	}
	return e.stk.stk[0]
}

// Error returns the error message of the underlying error.
// It implements the error interface.
func (e *stackErr) Error() string {
	return ReadTrace(e.stk.stk)
}

func (e *stackErr) From(err error) Stackable {
	if err == nil {
		return &stackErr{stk: newStack(e.Unwrap()...)}
	}
	flat := FlattenError(err)
	if e.IsEmpty() {
		return &stackErr{stk: newStack(flat...)}

	}
	return &stackErr{stk: newStack(flat...).Push(e.Unwrap()...)}
}

func (e *stackErr) Is(target error) bool {
	return isStack(e, target)
}

// IsEmpty checks if the error is empty.
// It implements the gtools.Stackable interface.
// It returns true if the error is empty, false otherwise.
func (e *stackErr) IsEmpty() bool {
	return e.stk.IsEmpty()
}

func (e *stackErr) Output() error {
	if e.IsEmpty() {
		return nil
	}
	return e
}

// Len returns the number of errors in the stack trace of the error.
// It implements the gtools.Stackable interface.
// It returns the number of errors in the stack trace of the error.
func (e *stackErr) Len() int {
	return e.stk.Len()
}

// Stack appends the given error to the stack trace of the error and returns the new stackable error.
// It implements the gtools.Stackable interface.
// It takes an error and appends it to the stack trace of the error.
// It returns the new stackable error.
func (e *stackErr) Stack(err error) {
	if err == nil {
		return
	}
	flat := FlattenError(err)
	e.stk = e.stk.Push(flat...)
}

// Unwrap returns the underlying error.
// It implements the gerrors.Wrapper interface.
func (e *stackErr) Unwrap() []error {
	return e.stk.stk
}

var _ error = (*stackErr)(nil)
var _ Stackable = (*stackErr)(nil)
var _ Wrapped = (*stackErr)(nil)

// ConcStackableOf returns a new concurrent stackable error from the given error.
// It implements the gtools.Stackable interface.
// It wraps the given error and provides a thread-safe stack of errors.
func ConcStackableOf(err error) Stackable {
	if err == nil {
		return &concStackErr{}
	}
	return &concStackErr{
		stk: newStack(FlattenError(err)...),
	}
}

func ConcStackOf(errs ...error) Stackable {
	if len(errs) == 0 {
		return &concStackErr{}
	}
	var out []error
	for _, err := range errs {
		flat := FlattenError(err)
		out = append(out, flat...)
	}
	return &concStackErr{
		stk: newStack(out...),
	}
}

// ConcStack returns a new concurrent stackable error.
func ConcStack() Stackable {
	return &concStackErr{}
}

type concStackErr struct {
	mtx sync.RWMutex
	stk stack
}

func (e *concStackErr) Cause() error {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	if e.stk.Len() == 0 {
		return nil
	}
	return e.stk.stk[0]
}

func (e *concStackErr) Error() string {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	return ReadTrace(e.stk.stk)
}

func (e *concStackErr) From(err error) Stackable {
	e.mtx.Lock()
	defer e.mtx.Unlock()
	if err == nil {
		return &concStackErr{stk: newStack(e.unwrap()...)}
	}
	flat := FlattenError(err)
	if e.isEmpty() {
		return &concStackErr{stk: newStack(flat...)}

	}
	return &concStackErr{stk: newStack(flat...).Push(e.unwrap()...)}
}

func (e *concStackErr) Is(target error) bool {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	return isStack(e, target)
}

func (e *concStackErr) IsEmpty() bool {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	return e.isEmpty()
}

func (e *concStackErr) isEmpty() bool {
	return e.len() == 0
}

func (e *concStackErr) Len() int {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	return e.len()
}

func (e *concStackErr) len() int {
	return e.stk.Len()
}

// Output this method is right. it returns itself as an error only if it isn't empty (structs can't be nil)
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
	flat := FlattenError(err)
	e.stk = e.stk.Push(flat...)
}

func (e *concStackErr) Unwrap() []error {
	e.mtx.RLock()
	defer e.mtx.RUnlock()
	return e.unwrap()
}

func (e *concStackErr) unwrap() []error {
	return e.stk.stk
}

var _ error = (*concStackErr)(nil)
var _ Stackable = (*concStackErr)(nil)
var _ Wrapped = (*concStackErr)(nil)

// OperationalOf creates an Operational from the given op and error.
//
// An Operational is an error that is associated with an op.
// It implements the gtools.Operational interface.
func OperationalOf(op string, err error) Operational {
	return &opErr{
		op:  op,
		err: err,
	}
}

// opErr is an implementation of gtools.Operational interface.
//
// It contains the op that caused the error as well as the error itself.
//
// The error returned by the defaultErr method is the same as the error returned by the
// underlying error.
type opErr struct {
	op  string
	err error
}

// defaultErr implements the error interface. It returns the string representation of
// the wrapped error.
func (e *opErr) Error() string {
	return e.err.Error()
}

// Operation returns the op that caused the error.
func (e *opErr) Operation() string {
	return e.op
}

var _ error = (*opErr)(nil)

// AsOperational is a function that checks if an error is an Operational.
//
// Parameters:
// - err: The error to check.
//
// Returns:
// - Operational: The error as Operational if it's of type Operational.
// - bool: True if the error is of type Operational, false otherwise.
func AsOperational(err error) (Operational, bool) {
	// Declare a variable of type Operational
	var opError Operational

	// Check if the error is of type Operational
	if errors.As(err, &opError) {
		// If it is, return the error and true
		return opError, true
	}

	// If the error is not of type Operational, check if it is a pointer to a struct
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
			// Return the new Operational and true
			return opError, true
		}
	}

	// If the error is not of type Operational, return nil and false
	return nil, false
}

var _ Operational = (*opErr)(nil)

// Tagged creates a new taggableError from the given error and tags.
//
// The given tags are stored in the returned error, and can be retrieved
// using the Tags() method.
//
// The returned error is of type gtools.Taggable.
func Tagged(err error, tags ...string) Taggable {
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

// defaultErr returns the error message of the underlying error.
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
// The returned slice is a clone of the internal tags map keys.
func (e *taggableError) Tags() []string {
	tags := make([]string, 0, len(e.tags))
	for tag := range e.tags {
		tags = append(tags, tag)
	}
	return tags
}

var _ error = (*taggableError)(nil)
