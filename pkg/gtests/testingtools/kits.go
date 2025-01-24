// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testingtools

import (
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/gtests/internal/logs"
)

// GtoolsLite returns a new tests.GtoolsLite instance configured with the provided
// testing and loggerLevel parameters.
func GtoolsLite[T gtests.FailureLoggableTesting](testing T, loggerLevel gtests.LoggerStrategy) gtests.GtoolsLite {
	return gtoolsLiteKit[T](testing, loggerLevel)
}

// gtoolsLiteKit builds a new gtoolsLite instance with a tests.FailureLoggableTesting,
// a tests.HelperTesting, a toolsLite and a levelLoggableLite implementations.
func gtoolsLiteKit[T gtests.FailureLoggableTesting](testing T, loggerLevel gtests.LoggerStrategy) *gtoolsLite[T] {
	return &gtoolsLite[T]{
		gtoolsRandSortable:      gtoolsRandSortable{},
		reflectionDataToolsLite: reflectionDataToolsLite{},
		HelperTesting:           testing,
		toolsLite: toolsLite{
			HelperTesting: testing,
		},
		levelLoggableLite: levelLoggableLite[T]{
			FailureLoggableTesting: testing,
			loggerLevel:            loggerLevel,
			logger:                 logs.TimerStack(),
		},
	}
}

type gtoolsLite[T gtests.FailureLoggableTesting] struct {
	gtoolsRandSortable
	reflectionDataToolsLite
	levelLoggableLite[T]
	gtests.HelperTesting
	toolsLite
}

func LoggableToolsLite[T gtests.FailureLoggableTesting](testing T, loggerLevel gtests.LoggerStrategy) gtests.LoggableTools {
	return loggableToolsLiteKit[T](testing, loggerLevel)
}

func loggableToolsLiteKit[T gtests.FailureLoggableTesting](testing T, loggerLevel gtests.LoggerStrategy) *loggableToolsLite[T] {
	return &loggableToolsLite[T]{
		HelperTesting: testing,
		toolsLite: toolsLite{
			HelperTesting: testing,
		},
		levelLoggableLite: levelLoggableLite[T]{
			FailureLoggableTesting: testing,
			loggerLevel:            loggerLevel,
			logger:                 logs.TimerStack(),
		},
	}
}

type loggableToolsLite[T gtests.FailureLoggableTesting] struct {
	gtests.HelperTesting
	toolsLite
	levelLoggableLite[T]
}

// Errorf logs a message to the console and marks the test as failed.
// The message is formatted using the fmt.Printf function.
// The arguments are passed to the fmt.Printf function to format the message.
func (t *loggableToolsLite[T]) Errorf(format string, args ...any) {
	t.Helper()
	t.HelperTesting.Errorf(format, args...)
}

// Error logs a message to the console and marks the test as failed.
// The message is passed to the fmt.Println function to print it to the console.
func (t *loggableToolsLite[T]) Error(args ...any) {
	t.Helper()
	t.HelperTesting.Error(args...)
}

// ConcLite creates a new LoggableRunnableTools instance from a FailureLoggableTesting instance and a loggerLevel.
// It returns a new concLite instance with the FailureLoggableTesting and loggerLevel parameters.
// The returned instance of LoggableRunnableTools is thread-safe.
func ConcLite[T gtests.FailureLoggableTesting](testing T, loggerLevel gtests.LoggerStrategy) gtests.LoggableRunnableTools {
	return concLiteKit(testing, loggerLevel)
}

func concLiteKit[T gtests.FailureLoggableTesting](testing T, loggerLevel gtests.LoggerStrategy) *concLite[T] {
	return &concLite[T]{
		toolsLite: toolsLite{
			HelperTesting: testing,
		},
		levelLoggableLite: levelLoggableLite[T]{
			FailureLoggableTesting: testing,
			loggerLevel:            loggerLevel,
			logger:                 logs.TimerStack(),
		},
	}
}

type concLite[T gtests.FailureLoggableTesting] struct {
	gtests.HelperTesting
	toolsLite
	concurrentContextToolsLite
	levelLoggableLite[T]
}

// Errorf logs a message to the console and marks the test as failed.
// The message is formatted using the fmt.Printf function.
// The arguments are passed to the fmt.Printf function to format the message.
func (t *concLite[T]) Errorf(format string, args ...any) {
	t.Helper()
	t.HelperTesting.Errorf(format, args...)
}

// Error logs a message to the console and marks the test as failed.
// The message is passed to the fmt.Println function to print it to the console.
func (t *concLite[T]) Error(args ...any) {
	t.Helper()
	t.HelperTesting.Error(args...)
}

var _ gtests.LoggableRunnableTools = (*concLite[gtests.FailureLoggableTesting])(nil)

// AsyncLite returns a new tests.LoggableAsyncRunnableTools instance that can be used to
// test functions with logging and runnable contexts. The loggerLevel parameter
// determines the minimum log level for which messages are
// logged.
// The returned instance of tests.LoggableAsyncRunnableTools is thread-safe.
func AsyncLite[T gtests.FailureLoggableTesting](testing T, loggerLevel gtests.LoggerStrategy) gtests.LoggableAsyncRunnableTools {
	return asyncLiteKit(testing, loggerLevel)
}

// asyncLiteKit builds a new asyncLite instance with a toolsLite, levelLoggableLite
// and asyncContextToolsLite implementations.
// The returned instance of tests.LoggableAsyncRunnableTools is thread-safe.
func asyncLiteKit[T gtests.FailureLoggableTesting](testing T, loggerLevel gtests.LoggerStrategy) *asyncLite[T] {
	return &asyncLite[T]{
		toolsLite: toolsLite{
			HelperTesting: testing,
		},
		levelLoggableLite: levelLoggableLite[T]{
			FailureLoggableTesting: testing,
			loggerLevel:            loggerLevel,
			logger:                 logs.TimerStack(),
		},
		asyncContextToolsLite: asyncContextToolsLite{},
	}
}

// asyncLite is a collection of tools selected for use on common scenarios for async tests.
type asyncLite[T gtests.FailureLoggableTesting] struct {
	gtests.HelperTesting
	toolsLite
	levelLoggableLite[T]
	asyncContextToolsLite
}

// Errorf logs a message to the console and marks the test as failed.
// The message is formatted using the fmt.Printf function.
// The arguments are passed to the fmt.Printf function to format the message.
func (t *asyncLite[T]) Errorf(format string, args ...any) {
	t.Helper()
	t.HelperTesting.Errorf(format, args...)
}

// Error logs a message to the console and marks the test as failed.
// The message is passed to the fmt.Println function to print it to the console.
func (t *asyncLite[T]) Error(args ...any) {
	t.Helper()
	t.HelperTesting.Error(args...)
}

var _ gtests.LoggableAsyncRunnableTools = (*asyncLite[gtests.FailureLoggableTesting])(nil)
