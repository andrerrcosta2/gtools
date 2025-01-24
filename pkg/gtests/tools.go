// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtests

import (
	"context"
	"github.com/andrerrcosta2/gtools/core/format"
	"github.com/andrerrcosta2/gtools/core/gtools"
	"github.com/andrerrcosta2/gtools/core/gtools/functions"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/conc/syncs"
)

type Tools interface {
	// Flag sets a flag value for the given ids.
	// The ids are used to identify the flag in the tools.
	// If the id is not present in the tools, it is created.
	// The value is the new value of the flag.
	// If the value is the same as the current value of the flag, the flag is not changed.
	Flag(value bool, ids ...string)
	// RegisterCallback registers a callback for the given ids.
	// The ids are used to identify the callback in the tools.
	// If the id is not present in the tools, it is created.
	// The callback is called when the
	RegisterCallback(callback functions.Runnable, value int, ids ...string)
	// RegisterCalls registers a call count for the given ids.
	// The ids are used to identify the call in the tools.
	// If the id is not present in the tools, it is created.
	// The value is incremented to a given id.
	RegisterCalls(call int, ids ...string)
	// CallsTo returns the number of calls for the given name.
	// The name is used to identify the call in the tools.
	// If the name is not present in the tools, it returns 0.
	CallsTo(name string) int
	// Condition checks if the given condition is true.
	// If the condition is false, it logs an error with the given message.
	// The message can contain placeholders for the given arguments.
	// The arguments are used to format the message.
	// This function is thread-safe and can be used concurrently.
	Condition(condition bool, notTrueMessage string, args ...any)
	// AssertCalls asserts that the number of calls for the given ids is equal to the given calls.
	// The ids are used to identify the call in the tools.
	// If the id is not present in the tools, it returns an error.
	// If the number of calls is not equal to the given calls, it logs an error with the given message.
	// The message can contain placeholders for the given arguments.
	// The arguments are used to format the message.
	// This function is thread-safe and can be used concurrently.
	AssertCalls(calls int, ids ...string)
	// AssertCallsTo asserts that the number of calls for the given name is equal to the given calls.
	// The name is used to identify the call in the tools.
	// If the name is not present in the tools, it returns an error.
	// If the number of calls is not equal to the given calls, it logs an error with the given message.
	// The message can contain placeholders for the given arguments.
	// The arguments are used to format the message.
	// This function is thread-safe and can be used concurrently.
	AssertCallsTo(name string, calls int, errorMessage string, args ...any)
	// AssertFlag asserts that the flag for the given ids is equal to the given flag.
	// The ids are used to identify the flag in the tools.
	// If the id is not present in the tools, it returns an error.
	// If the flag is not equal to the given flag, it logs an error with the given message.
	// The message can contain placeholders for the given arguments.
	// The arguments are used to format the message.
	// This function is thread-safe and can be used concurrently.
	AssertFlag(flag bool, ids ...string)
	// AssertFlagTo asserts that the flag for the given name is equal to the given flag.
	// The name is used to identify the flag in the tools.
	// If the name is not present in the tools, it returns an error.
	// If the flag is not equal to the given flag, it logs an error with the given message.
	// The message can contain placeholders for the given arguments.
	// The arguments are used to format the message.
	// This function is thread-safe and can be used concurrently.
	AssertFlagTo(name string, flag bool, errorMessage string, args ...any)
}

type DataTools interface {
	// DeepCopy uses reflection to duplicate the entire object and all objects it references,
	// recursively. This means that any nested or referenced data structures are fully
	// cloned, so changes in the deep copy don’t affect the original, and vice versa.
	// Each layer of the original data is copied independently.
	DeepCopy(data any) (any, error)
	// ShallowCopy uses reflection to duplicate only the top-level structure, leaving nested
	// or referenced objects shared between the original and the copy. For instance,
	// if the object has fields that point to other objects (like slices or pointers),
	// a shallow copy would copy only the references themselves, not the actual data they
	// point to. Therefore, changes to the shared objects will be reflected in both the
	// original and the shallow copy.
	ShallowCopy(value any) (any, error)
	// ExtractField extracts a named field from a struct.
	// The obj parameter is the struct to extract the field from.
	// The fieldName parameter is the name of the field to extract.
	// It returns the extracted field and an error if the field is not found.
	// The error is of type *gerrors.StackableError.
	ExtractField(obj any, fieldName string) (any, error)
	// InjectField injects a value into a named field in a struct.
	// The obj parameter is the struct to inject the value into.
	// The fieldName parameter is the name of the field to inject.
	// The value parameter is the value to inject.
	// It returns an error if the field is not found.
	// The error is of type *gerrors.StackableError.
	InjectField(obj any, fieldName string, value any) error
	// EqualsBy compares two structs for equality, with options to ignore specific fields.
	//
	// It takes two objects as 'a' and 'b' and compares them for equality.
	// Optionally, it accepts a variable number of strings to specify fields to ignore.
	//
	// It returns a boolean indicating whether the two objects are equal, and an error if any of the following conditions are met:
	// - The two objects are of different types.
	// - The two objects are not structs.
	// - The two objects are not equal.
	EqualsBy(a, b any, ignoreFields ...string) (bool, error)
	// Stringify takes an interface{} and returns a string representation of it in the
	// specified format.
	//
	// It takes the object to be stringified as 'data' and the format as 'format'.
	// It returns the string representation of the object, and an error if the object
	// cannot be stringified.
	Stringify(data any, format format.Serialization) (string, error)
}

type GToolsSortable interface {
	// Aggregable generates a slice of random AggregableOf objects.
	// It takes the number of objects to generate as 'amount'.
	Aggregable(amount int) []gtools.AggregableOf
	// Comparable generates a slice of random ComparableOf objects.
	// It takes the number of objects to generate as 'amount'.
	Comparable(amount int) []gtools.ComparableOf
	// PersistentComparable generates a slice of random PersistentComparableOf objects.
	// It takes the number of objects to generate as 'amount'.
	PersistentComparable(amount int) []gtools.PersistentComparableOf
	// PersistentSortable generates a slice of random PersistentSortableOf objects.
	// It takes the number of objects to generate as 'amount'.
	PersistentSortable(amount int) []gtools.PersistentSortableOf
	// Sortable generates a slice of random SortableOf objects.
	// It takes the number of objects to generate as 'amount'.
	Sortable(amount int) []gtools.SortableOf
	// Unique generates a slice of random UniqueOf objects.
	// It takes the number of objects to generate as 'amount'.
	Unique(amount int) []gtools.UniqueOf
}

type GToolsLoggable interface {
	GToolsSortable
	Loggable
}

type GtoolsLite interface {
	GToolsLoggable
	DataTools
	Tools
}

type Loggable interface {
	LoggableTesting
	// PrintLogStack prints the log stack to the test output.
	// It's thread-safe and can be used concurrently.
	PrintLogStack()
	// StackLog adds a message to the logger stack.
	// The message is added with the lock held, so it's thread-safe.
	// This function is thread-safe and can be used concurrently.
	StackLog(message string)
	// StackLogf adds a formatted message to the logger stack.
	// The message is added with the lock held, so it's thread-safe.
	// This function is thread-safe and can be used concurrently.
	StackLogf(format string, a ...any)
}

type LoggableTools interface {
	Loggable
	Tools
}

type AsyncRunnableTools interface {
	// RunBlocking executes the given function in a new goroutine and blocks the caller thread
	// until the context is canceled.
	RunBlocking(fn functions.Runnable, ctx context.Context)
	// AsyncBefore executes the given function in a new goroutine and returns immediately holding its context.
	AsyncBefore(fn functions.Runnable, ctx context.Context)
	// RunAfterBlocking executes the provided function inside a new goroutine after the context is
	// canceled blocking the caller thread.
	RunAfterBlocking(fn functions.Runnable, ctx context.Context)
	// AfterAsync executes the provided function inside a new goroutine after the context is canceled.
	AfterAsync(fn functions.Runnable, ctx context.Context)
	// Semaphore returns a semaphore that can be used to limit concurrency.
	Semaphore(maxConcurrent int) *syncs.ChannelSemaphore
}

type ConcurrentRunnableTools interface {
	// After runs the given function in a new goroutine after the context is canceled blocking
	// the caller thread.
	After(fn functions.Runnable, ctx context.Context)
	// Before runs the given function in a new goroutine and blocks the caller thread until the
	// context is canceled.
	Before(fn functions.Runnable, ctx context.Context)
	// Semaphore returns a semaphore that can be used to limit concurrency.
	Semaphore(maxConcurrent int) *syncs.ChannelSemaphore
}

type LoggableAsyncRunnableTools interface {
	LoggableTools
	AsyncRunnableTools
}

type LoggableRunnableTools interface {
	LoggableTools
	ConcurrentRunnableTools
}

type LoggerStrategy int

const (
	AlwaysPrintLog LoggerStrategy = iota
	LogOnErrors
	LogOnFailure
)
