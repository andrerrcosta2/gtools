// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gerrors

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/domain/execution"
	"runtime"
	"strings"
)

func printStackTrace() {
	stack := getStackTrace(1)
	for _, entry := range stack {
		filePath := "file://" + entry // Ensure clickable format
		fmt.Printf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\\n", filePath, entry)
	}
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
