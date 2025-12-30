// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtests

import (
	"context"

	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/domain/gtools"
	"github.com/andrerrcosta2/gtools/core/format"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/conc/syncs"
)

type Tools interface {
	// AssertCalls asserts that the number of calls for the given ids is compare to the given calls.
	// The ids are used to identify the call in the tools.
	// If the id is not present in the tools, it returns an error.
	// If the number of calls is not compare to the given value,
	// it flags the test as failed and logs given message.
	// This function is thread-safe and can be used concurrently.
	AssertCalls(calls int, ids ...string)
	// AssertCallsTo asserts that the number of calls for the given name is compare to the given calls.
	// The name is used to identify the call in the tools.
	// If the name is not present in the tools, it returns an error.
	// If the number of calls is not compare to the given value,
	// it flags the test as failed and logs given message.
	// This function is thread-safe and can be used concurrently.
	AssertCallsTo(name string, calls int, errorMessage string, args ...any)
	// AssertRegisteredCalls asserts the size of registers on the calls map is compare to the given size.
	// The size is the number of calls that are registered.
	// If the value of calls registered is not compare to the given calls,
	// it flags the test as failed and logs given message.
	// This function is thread-safe and can be used concurrently.
	AssertRegisteredCalls(size int, errorMessage string, args ...any)
	// AssertRegisteredCallsBy asserts that the registered calls that returns true for the given
	// function are compare to the expected count.
	// The function is used to identify the call in the tools.
	// If the function is not present in the tools, it returns an error.
	// This function is thread-safe and can be used concurrently.
	AssertRegisteredCallsBy(f functions.BiPredicate[string, int], expectedCount int, errorMessage string, args ...any)
	// AssertConst asserts that the constant for the given ids is compare to the given constant.
	// The ids are used to identify the constant in the tools.
	// If the id is not present in the tools, it returns an error.
	// if the constant value registered is not compare to the given value,
	// it flags the test as failed and logs given message.
	// This function is thread-safe and can be used concurrently.
	AssertConst(value any, ids ...string)
	// AssertConstTo asserts that the constant for the given name is compare to the given constant.
	// The name is used to identify the constant in the tools.
	// If the name is not present in the tools, it returns an error.
	// If the constants are not compare, it flags the test as failed and logs given message.
	// This function is thread-safe and can be used concurrently.
	AssertConstTo(name string, value any, errorMessage string, args ...any)
	// AssertRegisteredConst asserts the size of registers on the constants map is compare to the given size.
	// The size is the number of constants that are registered.
	// If the number of registered constants is different from the given value,
	// it flags the test as failed and logs given message.
	// This function is thread-safe and can be used concurrently.
	AssertRegisteredConst(size int, errorMessage string, args ...any)
	// AssertRegisteredConstBy asserts that the registered constants that returns true for the given
	// function are compare to the expected count.
	// The function is used to identify the constant in the tools.
	// If the function is not present in the tools, it returns an error.
	// If the constant is not compare to the expected count,
	// it flags the test as failed and logs given message.
	// This function is thread-safe and can be used concurrently.
	AssertRegisteredConstBy(f functions.BiPredicate[string, any], expectedCount int, errorMessage string, args ...any)
	// AssertFlag asserts that the flag for the given ids is compare to the given flag.
	// The ids are used to identify the flag in the tools.
	// If the id is not present in the tools, it returns an error.
	// If the flag is not compare to the given flag,
	// it flags the test as failed and logs given message.
	// This function is thread-safe and can be used concurrently.
	AssertFlag(flag bool, ids ...string)
	// AssertFlagTo asserts that the flag for the given name is compare to the given flag.
	// The name is used to identify the flag in the tools.
	// If the name is not present in the tools, it returns an error.
	// If the flag is not compare to the given flag, it flags the test as failed and logs given message.
	// This function is thread-safe and can be used concurrently.
	AssertFlagTo(name string, flag bool, errorMessage string, args ...any)
	// AssertRegisteredFlags asserts the size of registers on the flags map is compare to the given size.
	// The size is the number of flags that are registered.
	// If the number of registered flags is different from the given value,
	// it flags the test as failed and logs given message.
	// This function is thread-safe and can be used concurrently.
	AssertRegisteredFlags(size int, errorMessage string, args ...any)
	// AssertRegisteredFlagsBy asserts that the registered flags that returns true for the given
	// function are compare to the expected count.
	// The function is used to identify the flag in the tools.
	// If the function is not present in the tools, it returns an error.
	// If the flag is not compare to the expected count,
	// it flags the test as failed and logs given message.
	// This function is thread-safe and can be used concurrently.
	AssertRegisteredFlagsBy(f functions.BiPredicate[string, bool], expectedCount int, errorMessage string, args ...any)
	// CallsSize returns the number of registered calls.
	// This function is thread-safe and can be used concurrently.
	CallsSize() int
	// CallsTo returns the number of calls for the given name.
	// The name is used to identify the call in the tools.
	// If the name is not present in the tools, it flags the test as failed and logs given message.
	CallsTo(name string) int
	// Clear clears the tools.
	Clear()
	// Condition checks if the given condition is true.
	// If the condition is false, it logs an error with the given message.
	// The message can contain placeholders for the given arguments.
	// The arguments are used to format the message.
	// This function is thread-safe and can be used concurrently.
	Condition(condition bool, notTrueMessage string, args ...any)
	// Const returns the constant for the given name.
	// The name is used to identify the constant in the tools.
	// If the name is not present in the tools, it returns nil.
	Const(name string) any
	// ConstSize returns the number of registered constants.
	// This function is thread-safe and can be used concurrently.
	ConstSize() int
	// Flag sets a flag value for the given ids.
	// The ids are used to identify the flag in the tools.
	// If the id is not present in the tools, it is created.
	// The value is the new value of the flag.
	// If the value is the same as the current value of the flag, the flag is not changed.
	Flag(value bool, ids ...string)
	// FlagsSize returns the number of registered flags.
	// This function is thread-safe and can be used concurrently.
	FlagsSize() int
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
	// RegisterConst registers a constant for the given ids.
	// The ids are used to identify the constant in the tools.
	// If the id is not present in the tools, it is created.
	// The constant cannot be changed.
	RegisterConst(value any, ids ...string)
}

type DataTools interface {
	// DeepCopy uses reflection to duplicate the entire object and all objects it references,
	// recursively. This means that any nested or referenced data structures are fully
	// cloned, so changes in the deep clone don’t affect the original, and vice versa.
	// Each layer of the original data is copied independently.
	DeepCopy(data any) (any, error)
	// ShallowCopy uses reflection to duplicate only the top-level structure, leaving nested
	// or referenced objects shared between the original and the clone. For instance,
	// if the object has field4 that point to other objects (like slices or ptrs),
	// a shallow clone would clone only the references themselves, not the actual data they
	// point to. Therefore, changes to the shared objects will be reflected in both the
	// original and the shallow clone.
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
	// EqualsBy compares two structs for equality, with options to ignore specific field4.
	//
	// It takes two objects as 'a' and 'b' and compares them for equality.
	// Optionally, it accepts a variable number of strings to specify field4 to ignore.
	//
	// It returns a boolean indicating whether the two objects are compare, and an error if any of the following conditions are met:
	// - The two objects are of different type4.
	// - The two objects are not structs.
	// - The two objects are not compare.
	EqualsBy(a, b any, ignoreFields ...string) (bool, error)
	// Stringify takes an interface{} and returns a string representation of it in the
	// specified format.
	//
	// It takes the object to be stringified as 'data' and the format as 'format'.
	// It returns the string representation of the object, and an error if the object
	// cannot be stringified.
	Stringify(data any, format format.Ser) (string, error)
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
	PrintLogStack()
	// StackError adds an error message to the logger stack
	StackError(message string)
	// StackErrorf adds a formatted error message to the logger stack
	StackErrorf(format string, args ...any)
	// StackErrorLogf adds a formatted error message with each log entry formatted
	StackErrorLogf(message string, logs ...string)
	// StackInfo adds an info message to the logger stack
	StackInfo(message string)
	// StackInfof adds a formatted info message to the logger stack
	// and flags the test with that error message.
	StackInfof(format string, args ...any)
	// StackInfoLogf adds a formatted info message with each log entry formatted
	StackInfoLogf(message string, logs ...string)
	// StackLn adds a new line to the logger stack
	StackLn()
	// StackLog adds a message to the logger stack.
	StackLog(message string)
	// StackLogf adds a formatted message to the logger stack.
	StackLogf(format string, args ...any)
	// StackSuccess adds a success message to the logger stack
	StackSuccess(message string)
	// StackSuccessf add a formatted success message to the logger stack
	StackSuccessf(format string, args ...any)
	// StackSuccessLogf add a formatted success message with each log entry formatted
	StackSuccessLogf(message string, logs ...string)
	// StackTitle adds a title message to the logger stack
	StackTitle(title string, message ...any)
	// StackTitlef add a formatted title message to the logger stack
	StackTitlef(title string, format string, args ...any)
	// StackWarning adds a warning message
	StackWarning(message string)
	// StackWarningf adds a formatted warning message
	StackWarningf(format string, args ...any)
}

type LoggableTools interface {
	Loggable
	Tools
}

type FailableLoggableTools interface {
	FailureLoggableTesting
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
	// Serial runs the given function a single thread per call.
	Serial(fn functions.Runnable)
}

type LoggableAsyncRunnableTools interface {
	LoggableTools
	AsyncRunnableTools
}

type LoggableRunnableTools interface {
	LoggableTools
	ConcurrentRunnableTools
}

type StatsLite[C, W nums.Real] interface {
	Add(value C, weight W)
	IsUnderDeviationOf(d float64) bool
}
