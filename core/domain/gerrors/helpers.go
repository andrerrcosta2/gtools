// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gerrors

import (
	"errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/andrerrcosta2/gtools/core/domain/execution"
)

func newStackFromError(origin Stackable, from error) stack {
	if from == nil {
		return newStack(origin.Unwrap()...)
	}
	flat := FlattenError(from)
	if origin.IsEmpty() {
		return newStack(flat...)

	}
	return newStack(flat...).Push(origin.Unwrap()...)
}

func isStack(e Stackable, target error) bool {
	if target == nil {
		return e.IsEmpty()
	}
	for _, err := range e.Unwrap() {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}

func printStackTrace() {
	for _, entry := range getStackTrace(1) {
		filePath := "file://" + entry // Ensure clickable format
		fmt.Printf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\\n", filePath, entry)
	}
}

func getStackTrace(skip int) []string {
	pc := make([]uintptr, 10)
	n := runtime.Callers(skip+2, pc)
	frames := runtime.CallersFrames(pc[:n])

	var trace []string
	for {
		frame, more := frames.Next()
		if frame.Function == "" || shouldSkipFunction(frame.Function) {
			continue
		}
		file := execution.GetFileDisplayPath(frame.File, "")
		trace = append(trace, fmt.Sprintf("%s:%d → %s()", file, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	return trace
}

func shouldSkipFunction(funcName string) bool {
	skippedPrefixes := []string{
		"runtime.", "testing.", "log.", "mylogger.", // Ignore Go internals & logging
	}
	for _, prefix := range skippedPrefixes {
		if strings.HasPrefix(funcName, prefix) {
			return true
		}
	}
	return false
}
