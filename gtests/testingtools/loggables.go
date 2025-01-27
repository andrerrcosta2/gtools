// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testingtools

import (
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/gtests/internal/logs"
)

// LoggersLite creates a new tests.Loggable from a tests.FailureLoggableTesting instance.
// It prints the log stack to the console based on the logger strategy.
// The returned instance of LoggableTesting is thread-safe.
func LoggersLite[T gtests.FailureLoggableTesting](testing T, loggerLevel gtests.LoggerStrategy) gtests.Loggable {
	return &levelLoggableLite[T]{
		FailureLoggableTesting: testing,
		loggerLevel:            loggerLevel,
		logger:                 logs.TimerStack(),
	}
}

type levelLoggableLite[T gtests.FailureLoggableTesting] struct {
	gtests.FailureLoggableTesting
	loggerLevel gtests.LoggerStrategy
	logger      logs.LogStack
}

func (t *levelLoggableLite[T]) StackLog(message string) {
	t.logger.Add(message)
}

func (t *levelLoggableLite[T]) StackLogf(format string, a ...any) {
	t.logger.Addf(format, a...)
}

func (t *levelLoggableLite[T]) PrintLogStack() {
	if shouldLog(t, t.loggerLevel) {
		t.Helper()
		t.logger.Logt(t.FailureLoggableTesting)
	}
}

var _ gtests.FailureLoggableTesting = (*levelLoggableLite[gtests.FailureLoggableTesting])(nil)
var _ gtests.Loggable = (*levelLoggableLite[gtests.FailureLoggableTesting])(nil)
