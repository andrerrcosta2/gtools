// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

// Package opt: this package is a wrapper around the Option struct for some convenience methods.
//
// 1. These methods don't handle "zero-values".
// 2. It doesn't perform Thread-Safe ops.
// 3. This is useful in avoiding nil pointer panics.
package opt

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/generics"
	"reflect"
)

// Option is a wrapper for nullable variables.
// This isn't Thread-Safe.
type Option[T any] struct {
	value *T
	isSet bool
}

// Get returns the value stored in the Option struct.
//
// Returns:
// - The value stored in the Option struct.
func (o *Option[T]) Get() T {
	if o.value == nil {
		panic("Optional is empty")
	}
	return *o.value
}

// IsPresent checks if the Option is present.
//
// Returns:
// - bool: True if the Option is present, false otherwise.
func (o *Option[T]) IsPresent() bool {
	// The IsSet field indicates if the Option is present.
	return o.isSet
}

// IfPresent checks if the Option is present and, if so, calls the provided function with the value.
//
// Parameters:
// fn: a function which takes a value of type T and returns nothing.
func (o *Option[T]) IfPresent(consumer functions.Consumer[T]) {
	// Check if the Option is present
	if o.IsPresent() {
		// If it is, call the provided function with the value
		consumer(*o.value)
	}
}

func isDeepReflectedNullable[T any](value *T) bool {
	if value == nil {
		return true
	}

	// Use reflection to check for nil slices, maps, or ptrs
	val := reflect.ValueOf(value).Elem()
	return (val.Kind() == reflect.Slice || val.Kind() == reflect.Map || val.Kind() == reflect.Ptr) && val.IsNil()
}

// Filter returns the Option if the predicate is true
func (o *Option[T]) Filter(predicate functions.Predicate[T]) *Option[T] {
	if o.isSet && predicate(*o.value) {
		return o
	}
	return None[T]()
}

// FlatMap returns the result of the provided function if the option is present
func (o *Option[T]) FlatMap(fn functions.Function[T, *Option[T]]) *Option[T] {
	if o.isSet {
		return fn(*o.value)
	}
	return o
}

// Match returns the result of the provided function if the option is present and the none function otherwise
func (o *Option[T]) Match(some functions.Function[T, any], none functions.Supplier[any]) any {
	if o.isSet {
		return some(*o.value)
	}
	return none()
}

// Map returns the result of the provided function if the option is present
func (o *Option[T]) Map(fn functions.Function[T, T]) *Option[T] {
	if o.isSet {
		return Of(fn(*o.value))
	}
	return o
}

// None returns a new Option with no value set.
//
// The Option returned by None is considered to be "None" or "unset".
// It is used to represent the absence of a value.
//
// Parameters:
// - None doesn't take any parameters.
//
// Returns:
// - A pointer to a None Option.
func None[T any]() *Option[T] {
	return &Option[T]{
		isSet: false,
	}
}

// Of returns a new Option with the given value.
//
// Parameters:
// - value: the value to be wrapped in the Option.
//
// Returns:
// - A pointer to the Option that contains the given value.
func Of[T any](value T) *Option[T] {
	return &Option[T]{
		value: &value,
		isSet: true,
	}
}

// OfNullable returns an Option that contains the value of the given pointer if it isn't nil.
// If the pointer is nil, it returns a None Option.
// It is used to represent the absence of a value.
// It doesn't handle "zero-values".
//
// Parameters:
// - value: a pointer to the value to be wrapped in an Option.
//
// Returns:
// - A pointer to the Option that contains the value of the given pointer if it isn't nil.
// - A pointer to a None Option if the given pointer is nil.
func OfNullable[T any](value *T) *Option[T] {
	// Check if the pointer is nil
	if value == nil {
		// Return a None Option
		return None[T]()
	}
	// Check if the value is nil
	// Reflection pays the price for generalizations.
	// If you disagree, you may create different optionals for each type then handle it natively
	if isDeepReflectedNullable(value) {
		// Return a None Option
		return None[T]()
	}

	// Return a new Option with the value of the given pointer
	return Of(*value)
}

func (o *Option[T]) OrAssert(value *T) *Option[T] {
	// Check if the previous value exists
	if o.isSet {
		return o
	}
	// Check if the pointer is nil
	if value == nil {
		// Panic if the pointer is nil
		panic("OrAssert panics if the provided pointer is nil")
	}
	// Check if the value is nil
	// Reflection pays the price for generalizations.
	// If you don't agree, you may create different optionals for each type then handle it natively
	if isDeepReflectedNullable(value) {
		// Panic if the pointer is nil
		panic("OrAssert panics if the provided pointer is nil")
	}
	return Of(*value)
}

func (o *Option[T]) OrElse(value T) *Option[T] {
	if o.isSet {
		return o
	}
	return Of(value)
}

// OrElseGet returns the value of the Option if it is set, otherwise it returns the provided default value.
//
// Parameters:
// - value: The default value to return if the Option isn't set.
//
// Returns:
// - The value of the Option if it is set, otherwise the provided default value.
func (o *Option[T]) OrElseGet(value T) T {
	// Check if the Option is set
	if o.isSet {
		// Return the value of the Option
		return *o.value
	}
	// Return the provided default value
	return value
}

// OrElseGetPtr returns the value of the Option if it is set, otherwise it returns the provided default value.
//
// Parameters:
// - value: The default value to return if the Option isn't set.
//
// Returns:
// - The value of the Option if it is set, otherwise the provided default value.
func (o *Option[T]) OrElseGetPtr(value *T) *T {
	// Check if the Option is set
	if o.isSet {
		return o.value
	}
	// Return the provided default value
	return value
}

// OrElseFunc returns the value of the Option if it is set, otherwise it returns the result of the provided function.
func (o *Option[T]) OrElseFunc(supplier functions.Supplier[T]) T {
	if o.isSet {
		return *o.value
	}
	return supplier()
}

// OrElseZero returns the value of the Option if it is set, otherwise it returns the zero value of the type T.
func (o *Option[T]) OrElseZero() T {
	if o.isSet {
		return *o.value
	}
	var zero T
	return zero
}

// OrPanic returns the value of the Option if it is set, otherwise it panics with the provided message.
func (o *Option[T]) OrPanic(msg string) T {
	if o.isSet {
		return *o.value
	}
	panic(msg)
}

// Peek calls the provided function with the value if the option is present
func (o *Option[T]) Peek(consumer functions.Consumer[T]) *Option[T] {
	if o.isSet {
		consumer(*o.value)
	}
	return o
}

// Set sets the value of the Option and marks it as set.
// It returns a pointer to the modified Option.
//
// Parameters:
// - value: The new value to set for the Option.
//
// Returns:
// - *Option[T]: a pointer to the modified Option.
func (o *Option[T]) Set(value T) *Option[T] {
	// ToSet the value of the Option
	o.value = &value
	// Mark the Option as set
	o.isSet = true
	// Return a pointer to the modified Option
	return o
}

func (o *Option[T]) ToPtr() *T {
	if o.isSet {
		return o.value
	}
	return nil
}

func (o *Option[T]) ToSlice() []T {
	if o.isSet {
		return []T{*o.value}
	}
	return nil
}

// Unset sets the IsSet field of the Option struct to false and returns a pointer to the modified Option struct.
// This function is used to unset the value of an Option.
//
// Returns:
// - A pointer to the modified Option struct.
func (o *Option[T]) Unset() *Option[T] {
	o.isSet = false // ToSet the IsSet field to false
	return o        // Return a pointer to the modified Option struct
}

// Variadic returns an Option that contains the value of the first element of the given variadic list
// if it isn't nil.
// If the variadic list is empty, it returns a None Option.
// It is used to represent the absence of a value.
// It doesn't handle "zero-values".
//
// Parameters:
// - values: a variadic list of values to be wrapped in an Option.
//
// Returns:
// - A pointer to the Option that contains the value of the first element of the given variadic list if it isn't nil.
// - A pointer to a None Option if the given variadic list is empty.
func Variadic[T any](values ...T) *Option[T] {
	if len(values) == 0 {
		return None[T]()
	}
	return Of(values[0])
}

// Zero returns an Option that contains the value if it isn't zero.
// If the value is zero, it returns a None Option.
// It is used to represent the absence of a value from a point of view of zero-values.
// That means it considers as nil a zero-value, which might be a false-positive depending
// on the requirements of your application.
//
// Parameters:
// - value: the value to be wrapped in an Option.
//
// Returns:
// - A pointer to the Option that contains the value if it isn't zero.
// - A pointer to a None Option if the value is zero.
func Zero[T any](value T) *Option[T] {
	if constraints.IsZeroOf[T](value) {
		return None[T]()
	}
	return Of(value)
}

// Zip combines two Optionals into a Tuple2.
//
// Parameters:
// - optA: The first Optional to combine.
// - optB: The second Optional to combine.
// - t2: The Tuple2 to combine the Optionals into.
//
// Returns:
// - *Option[Tuple2]: A pointer to the combined Tuple2.
func Zip[A, B any, Z generics.Tuple2[A, B]](optA *Option[A], optB *Option[B], t2 Z) *Option[Z] {
	if optA.isSet && optB.isSet {
		t2.SetFirst(*optA.value)
		t2.SetSecond(*optB.value)
		return Of(t2)
	}
	return None[Z]()
}
