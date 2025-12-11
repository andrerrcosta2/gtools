// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testingtools

import (
	"strings"

	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/format/printer"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/gtests"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/internal/logs"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/themes"
)

// LoggersLite creates a new gtests.Loggable from a gtests.FailureLoggableTesting instance.
// It prints the log stack to the console based on the logger strategy.
// The returned instance of LoggableTesting is thread-safe.
func LoggersLite[T gtests.FailureLoggableTesting](testing T, loggerLevel gtests.LoggerStrategy, theme ...themes.Theme) gtests.Loggable {
	testing.Helper()
	if len(theme) == 0 {
		return newLevelLoggableLite(testing, loggerLevel)
	}
	return newTmLevelLoggableLite(testing, loggerLevel, theme[0])
}

func newLevelLoggableLite[T gtests.FailureLoggableTesting](testing T, loggerLevel gtests.LoggerStrategy) *levelLoggableLite[T] {
	tools := &levelLoggableLite[T]{
		FailureLoggableTesting: testing,
		loggerLevel:            loggerLevel,
		logger:                 logs.TimerStack(),
		errors:                 0,
	}
	// Automatically defer PrintLogStack to ensure logs are printed on failure or skip
	testing.Cleanup(func() {
		if shouldLog(tools, tools.loggerLevel) {
			tools.Helper()

			// Check if the test has already failed
			if tools.Failed() {
				tools.Errorf("[Error] Test failed. %d errors were found.: check the logs for details...", tools.errors)
			}

			// Print logs
			tools.logger.Logt(tools.FailureLoggableTesting)
		}
	})
	return tools
}

type levelLoggableLite[T gtests.FailureLoggableTesting] struct {
	gtests.FailureLoggableTesting
	loggerLevel gtests.LoggerStrategy
	logger      logs.LogStack
	errors      int
}

func (t *levelLoggableLite[T]) Error(a ...any) {
	t.Helper()
	t.errors++
	t.FailureLoggableTesting.Error(a...)
}

func (t *levelLoggableLite[T]) Errorf(format string, a ...any) {
	t.Helper()
	t.errors++
	t.FailureLoggableTesting.Errorf(format, a...)
}

func (t *levelLoggableLite[T]) StackError(message string) {
	t.errors++
	t.logger.Addf("[failed]: %s", message)
}

func (t *levelLoggableLite[T]) StackErrorf(format string, a ...any) {
	t.errors++
	t.logger.Addf("[failed]: %s", printer.Sprintf(format, a...))
}

func (t *levelLoggableLite[T]) StackErrorLogf(message string, logs ...string) {
	sb := strings.Builder{}
	for _, log := range logs {
		sb.WriteString(log)
		sb.WriteString("\n")
	}
	t.logger.Addf("[Error Log]: \n%s", printer.Sprintf(message, sb.String()))
	t.Error(message)
}

func (t *levelLoggableLite[T]) StackInfo(message string) {
	t.logger.Addf("[info]: %s", message)
}

func (t *levelLoggableLite[T]) StackInfof(format string, a ...any) {
	t.logger.Addf("[info]: %s", printer.Sprintf(format, a...))
}

func (t *levelLoggableLite[T]) StackInfoLogf(message string, logs ...string) {
	sb := strings.Builder{}
	for _, log := range logs {
		sb.WriteString(log)
		sb.WriteString("\n")
	}
	t.logger.Addf("Info Log: \n%s", printer.Sprintf(message, sb.String()))
}

func (t *levelLoggableLite[T]) StackLn() {
	t.logger.Add("\n")
}

func (t *levelLoggableLite[T]) StackLog(message string) {
	t.logger.Add(message)
}

func (t *levelLoggableLite[T]) StackLogf(format string, a ...any) {
	t.logger.Addf(format, a...)
}

func (t *levelLoggableLite[T]) StackSuccess(message string) {
	t.logger.Addf("[success]: %s", message)
}

func (t *levelLoggableLite[T]) StackSuccessf(format string, a ...any) {
	t.logger.Addf("[success]: %s", printer.Sprintf(format, a...))
}

func (t *levelLoggableLite[T]) StackSuccessLogf(message string, logs ...string) {
	sb := strings.Builder{}
	for _, log := range logs {
		sb.WriteString(log)
		sb.WriteString("\n")
	}
	t.logger.Addf("[Success Log]: \n%s", printer.Sprintf(message, sb.String()))
}

func (t *levelLoggableLite[T]) StackTitle(title string, message ...any) {
	t.logger.Addf("[%s]: %s", strings.ToUpper(title), printer.Sprint(message...))
}

func (t *levelLoggableLite[T]) StackTitlef(title string, format string, a ...any) {
	t.logger.Addf("[%s]: %s", strings.ToUpper(title), printer.Sprintf(format, a...))
}

func (t *levelLoggableLite[T]) StackWarning(message string) {
	t.logger.Addf("[warning]: %s", message)
}

func (t *levelLoggableLite[T]) StackWarningf(format string, a ...any) {
	msg := printer.Sprintf(format, a...)
	t.logger.Addf("[warning]: %s", msg)
}

func (t *levelLoggableLite[T]) PrintLogStack() {
	if shouldLog(t, t.loggerLevel) {
		t.Helper()
		t.logger.Logt(t.FailureLoggableTesting)
	}
}

var _ gtests.FailureLoggableTesting = (*levelLoggableLite[gtests.FailureLoggableTesting])(nil)
var _ gtests.Loggable = (*levelLoggableLite[gtests.FailureLoggableTesting])(nil)

// newTmLevelLoggableLite creates a new gtests.Loggable from a gtests.FailureLoggableTesting instance.
// It prints the log stack to the console based on the logger strategy.
// The returned instance of LoggableTesting is thread-safe.
func newTmLevelLoggableLite[T gtests.FailureLoggableTesting](testing T, loggerLevel gtests.LoggerStrategy, theme themes.Theme) *tmLevelLoggableLite[T] {
	testing.Helper()
	tt := &tmLevelLoggableLite[T]{
		FailureLoggableTesting: testing,
		loggerLevel:            loggerLevel,
		logger:                 logs.TmTimerStack(theme),
		icons:                  theme,
	}

	// Automatically defer PrintLogStack to ensure logs are printed on failure or skip
	testing.Cleanup(func() {
		if shouldLog(tt, tt.loggerLevel) {
			tt.Helper()
			// Check if the test has already failed
			if tt.Failed() {
				tt.Errorf("%s Test failed. %d errors were found.:", tt.icons.Error, tt.errors)
			}

			// Print logs
			tt.logger.Logt(tt.FailureLoggableTesting)
		}
	})

	return tt
}

type tmLevelLoggableLite[T gtests.FailureLoggableTesting] struct {
	gtests.FailureLoggableTesting
	loggerLevel gtests.LoggerStrategy
	logger      logs.LogStack
	errors      int
	icons       themes.Theme
}

func (t *tmLevelLoggableLite[T]) Error(a ...any) {
	t.Helper()
	t.errors++
	t.FailureLoggableTesting.Error(a...)
}

func (t *tmLevelLoggableLite[T]) Errorf(format string, a ...any) {
	t.Helper()
	t.errors++
	t.FailureLoggableTesting.Errorf(format, a...)
}

func (t *tmLevelLoggableLite[T]) StackError(message string) {
	t.Helper()
	t.errors++
	t.logger.Addf("%s failed: %s", t.icons.Error, message)
}

func (t *tmLevelLoggableLite[T]) StackErrorf(format string, a ...any) {
	t.Helper()
	t.errors++
	msg := printer.Sprintf(format, a...)
	t.logger.Addf("%s failed: %s", t.icons.Error, msg)
}

func (t *tmLevelLoggableLite[T]) StackErrorLogf(message string, logs ...string) {
	sb := strings.Builder{}
	for _, log := range logs {
		t.errors++
		sb.WriteString(log)
		sb.WriteString("\n")
	}
	t.logger.Addf("Error Log: \n%s", printer.Sprintf(message, sb.String()))
	t.Error(message)
}

func (t *tmLevelLoggableLite[T]) StackInfo(message string) {
	t.logger.Addf("%s info: %s", t.icons.Info, message)
}

func (t *tmLevelLoggableLite[T]) StackInfof(format string, a ...any) {
	t.logger.Addf("%s info: %s", t.icons.Info, printer.Sprintf(format, a...))
}

func (t *tmLevelLoggableLite[T]) StackInfoLogf(message string, logs ...string) {
	sb := strings.Builder{}
	for _, log := range logs {
		sb.WriteString(log)
		sb.WriteString("\n")
	}
	t.logger.Addf("Info Log: \n%s", printer.Sprintf(message, sb.String()))
}

func (t *tmLevelLoggableLite[T]) StackLn() {
	t.logger.Add("================================================\n")
}

func (t *tmLevelLoggableLite[T]) StackLog(message string) {
	t.logger.Add(message)
}

func (t *tmLevelLoggableLite[T]) StackLogf(format string, a ...any) {
	t.logger.Addf(format, a...)
}

func (t *tmLevelLoggableLite[T]) StackSuccess(message string) {
	t.logger.Addf("%s success: %s", t.icons.Success, message)
}

func (t *tmLevelLoggableLite[T]) StackSuccessf(format string, a ...any) {
	t.logger.Addf("%s success: %s", t.icons.Success, printer.Sprintf(format, a...))
}

func (t *tmLevelLoggableLite[T]) StackSuccessLogf(message string, logs ...string) {
	sb := strings.Builder{}
	for _, log := range logs {
		sb.WriteString(log)
		sb.WriteString("\n")
	}
	t.logger.Addf("Success Log: \n%s", printer.Sprintf(message, sb.String()))
}

func (t *tmLevelLoggableLite[T]) StackTitle(title string, message ...any) {
	t.logger.Addf("%s: %s", fmx.SBoldf("[%s]", strings.ToUpper(title)), printer.Sprint(message...))
}

func (t *tmLevelLoggableLite[T]) StackTitlef(title string, format string, a ...any) {
	t.logger.Addf("%s: %s", fmx.SBoldf("[%s]", strings.ToUpper(title)), printer.Sprintf(format, a...))
}

func (t *tmLevelLoggableLite[T]) StackWarning(message string) {
	t.logger.Addf("%s warning: %s", t.icons.Warning, message)
}

func (t *tmLevelLoggableLite[T]) StackWarningf(format string, a ...any) {
	msg := printer.Sprintf(format, a...)
	t.logger.Addf("%s warning: %s", t.icons.Warning, msg)
}

func (t *tmLevelLoggableLite[T]) PrintLogStack() {
	if shouldLog(t, t.loggerLevel) {
		t.Helper()
		t.logger.Logt(t.FailureLoggableTesting)
	}
}

var _ gtests.FailureLoggableTesting = (*levelLoggableLite[gtests.FailureLoggableTesting])(nil)
var _ gtests.Loggable = (*levelLoggableLite[gtests.FailureLoggableTesting])(nil)
