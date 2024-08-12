// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

// Package opt: this package is a wrapper around the Option struct for some convenience methods.
//
// 1. These methods don't handle "zero-values".
// 2. It doesn't perform Thread-Safe operations.
// 3. This is useful in avoiding nil pointer panics.
package opt

import (
	"github.com/andrerrcosta2/gtools/pkg/functions"
	"reflect"
)

// Option is a wrapper for nullable variables.
// This isn't Thread-Safe.
type Option[T any] struct {
	value *T
	isSet bool
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

// Set sets the value of the Option and marks it as set.
// It returns a pointer to the modified Option.
//
// Parameters:
// - value: The new value to set for the Option.
//
// Returns:
// - *Option[T]: a pointer to the modified Option.
func (o *Option[T]) Set(value T) *Option[T] {
	// Set the value of the Option
	o.value = &value
	// Mark the Option as set
	o.isSet = true
	// Return a pointer to the modified Option
	return o
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

// Unset sets the IsSet field of the Option struct to false and returns a pointer to the modified Option struct.
// This function is used to unset the value of an Option.
//
// Returns:
// - A pointer to the modified Option struct.
func (o *Option[T]) Unset() *Option[T] {
	o.isSet = false // Set the IsSet field to false
	return o        // Return a pointer to the modified Option struct
}

func (o *Option[T]) OrAssert(value *T) *Option[T] {
	// Check if the previous value exists
	if o.isSet {
		return o
	}
	// Check if the pointer is nil
	if value == nil {
		// Panic if the pointer is nil
		panic("Or assert will panic if the provided pointer is nil")
	}
	// Check if the value is nil
	// Reflection pays the price for generalizations.
	// If you don't agree, you may create different optionals for each type then handle it natively
	if isDeepReflectedNullable(value) {
		// Panic if the pointer is nil
		panic("Or assert will panic if the provided pointer is nil")
	}
	return Of(*value)
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
	// If you don't agree, you may create different optionals for each type then handle it natively
	if isDeepReflectedNullable(value) {
		// Return a None Option
		return None[T]()
	}

	// Return a new Option with the value of the given pointer
	return Of(*value)
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

	// Use reflection to check for nil slices, maps, or pointers
	val := reflect.ValueOf(value).Elem()
	return (val.Kind() == reflect.Slice || val.Kind() == reflect.Map || val.Kind() == reflect.Ptr) && val.IsNil()
}

// EmptyString returns the non-empty string between `a` and `b`.
//
// If `a` is empty, it returns `b`. Otherwise, it returns `a`.
//
// Parameters:
// - a: The first string to compare.
// - b: The second string to compare.
//
// Returns:
// - string: The non-empty string between `a` and `b`.
func EmptyString(a, b string) string {
	if a == "" {
		return b
	}
	return a
}
