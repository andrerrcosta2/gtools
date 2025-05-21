// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package logs

import (
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/gtests"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/themes"
	"sort"
	"strings"
	"sync"
	"time"
)

type LogStack interface {
	Add(message string)
	Addf(format string, a ...any)
	Logt(t gtests.LoggableTesting)
	Print()
	Sprint() string
}

// Entry returns a LogEntry with the elapsed time from the start time
// and the formatted message.
//
// The start time is used to calculate the elapsed time.
// The format string is used to format the message with the arguments.
// The arguments are passed after the format string.
func Entry(start time.Time, format string, a ...any) LogEntry {
	return LogEntry{
		Message: fmx.Sprintf(format, a...),
		Elapsed: time.Since(start),
	}
}

// LogEntry represents a log message and its elapsed time
type LogEntry struct {
	Message string
	Elapsed time.Duration
}

// TimerStack creates a new timerLogStack and captures the start time
//
// It creates a new timerLogStack, captures the current time as the start time,
// and returns the timerLogStack.
func TimerStack() LogStack {
	return &timerLogStack{
		entries: []LogEntry{},
		start:   time.Now(), // Capture the start time when the timerLogStack is created
	}
}

// timerLogStack holds all log entries and records the start time
type timerLogStack struct {
	mtx     sync.RWMutex
	entries []LogEntry
	start   time.Time
}

// Add adds a new log entry with a message and automatically captures the elapsed time
func (ls *timerLogStack) Add(message string) {
	ls.mtx.Lock()
	defer ls.mtx.Unlock()
	ls.entries = append(ls.entries, Entry(ls.start, message))
}

// Addf adds a new log entry with a formatted message and automatically captures the elapsed time
func (ls *timerLogStack) Addf(format string, a ...any) {
	ls.mtx.Lock()
	defer ls.mtx.Unlock()
	ls.entries = append(ls.entries, Entry(ls.start, format, a...))
}

// Print prints all log entries sorted by their elapsed time in ascending order
func (ls *timerLogStack) Print() {
	ls.mtx.RLock()
	defer ls.mtx.RUnlock()
	// Sort the entries by the elapsed time in ascending order
	sort.Slice(ls.entries, func(i, j int) bool {
		return ls.entries[i].Elapsed < ls.entries[j].Elapsed
	})

	// Print each log entry
	for _, entry := range ls.entries {
		fmx.Printf("[%s] %s\n", entry.Elapsed, entry.Message)
	}
}

// Sprint returns a string representation of the log stack sorted by their elapsed time in ascending order
func (ls *timerLogStack) Sprint() string {
	ls.mtx.RLock()
	defer ls.mtx.RUnlock()
	// Sort the entries by the elapsed time in ascending order
	//sort.Slice(ls.entries, func(i, j int) bool {
	//	return ls.entries[i].Elapsed < ls.entries[j].Elapsed
	//})

	// Print each log entry
	var sb strings.Builder
	for _, entry := range ls.entries {
		sb.WriteString(fmx.Sprintf("[%s] %s\n", entry.Elapsed, entry.Message))
	}
	return sb.String()
}

// Logt logs the log stack to the test output.
func (ls *timerLogStack) Logt(t gtests.LoggableTesting) {
	t.Helper()
	// Print the log stack to the test output
	t.Logf("\n%s\n[Log size: %d] ⏱ elapsed time: %s\n", ls.Sprint(), len(ls.entries), time.Since(ls.start))
}

// TmTimerStack creates a new timerLogStack and captures the start time
//
// It creates a new timerLogStack, captures the current time as the start time,
// and returns the timerLogStack.
func TmTimerStack(theme themes.Theme) LogStack {
	return &tmTimerLogStack{
		entries: []LogEntry{},
		start:   time.Now(), // Capture the start time when the timerLogStack is created
		theme:   theme,
	}
}

// timerLogStack holds all log entries and records the start time
type tmTimerLogStack struct {
	mtx     sync.RWMutex
	entries []LogEntry
	start   time.Time
	theme   themes.Theme
}

// Add adds a new log entry with a message and automatically captures the elapsed time
func (ls *tmTimerLogStack) Add(message string) {
	ls.mtx.Lock()
	defer ls.mtx.Unlock()
	ls.entries = append(ls.entries, Entry(ls.start, message))
}

// Addf adds a new log entry with a formatted message and automatically captures the elapsed time
func (ls *tmTimerLogStack) Addf(format string, a ...any) {
	ls.mtx.Lock()
	defer ls.mtx.Unlock()
	ls.entries = append(ls.entries, Entry(ls.start, format, a...))
}

// Print prints all log entries sorted by their elapsed time in ascending order
func (ls *tmTimerLogStack) Print() {
	ls.mtx.RLock()
	defer ls.mtx.RUnlock()
	// Sort the entries by the elapsed time in ascending order
	sort.Slice(ls.entries, func(i, j int) bool {
		return ls.entries[i].Elapsed < ls.entries[j].Elapsed
	})

	// Print each log entry
	for _, entry := range ls.entries {
		fmx.Printf("%s %s\n", fmx.SBoldf("[%s]", entry.Elapsed), entry.Message)
	}
}

// Sprint returns a string representation of the log stack sorted by their elapsed time in ascending order
func (ls *tmTimerLogStack) Sprint() string {
	ls.mtx.RLock()
	defer ls.mtx.RUnlock()

	// Print each log entry
	var sb strings.Builder
	for _, entry := range ls.entries {
		sb.WriteString(fmx.Sprintf("[%s] %s\n", entry.Elapsed, entry.Message))
	}
	return sb.String()
}

// Logt logs the log stack to the test output.
func (ls *tmTimerLogStack) Logt(t gtests.LoggableTesting) {
	t.Helper()
	// Print the log stack to the test output
	t.Logf(fmx.Sprintf("\n%s\n[Log size: %d] %s elapsed time: %s\n", ls.Sprint(),
		len(ls.entries), ls.theme.Clock, time.Since(ls.start)))
}
