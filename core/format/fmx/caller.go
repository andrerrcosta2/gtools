// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fmx

import (
	"github.com/andrerrcosta2/gtools/core/domain/execution"
	"github.com/andrerrcosta2/gtools/core/format/printer"
	"runtime"
	"strings"
)

const (
	ansi = "\x1b]8;;%s\x1b\\%s:%d\x1b]8;;\x1b\\: %s"
)

func Call(s ...any) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		printer.Print("⚠️ [?:?]: " + printer.Sprint(s...))
		return
	}

	msg := printer.Sprint(s...)
	displayFile := execution.GetFileDisplayPath(file, "")

	if execution.SupportsTerminalHyperlink {
		fullPath := "file://" + file
		printer.Printf(ansi, fullPath, displayFile, line, msg)
	} else {
		printer.Printf("%s:%d: %s\n", displayFile, line, msg)
	}
}

func Callf(format string, args ...any) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		printer.Printf("⚠️ [?:?]: "+format+"\n", args...)
		return
	}

	msg := printer.Sprintf(format, args...)
	displayFile := execution.GetFileDisplayPath(file, "")

	if execution.SupportsTerminalHyperlink {
		fullPath := "file://" + file
		printer.Printf(ansi, fullPath, displayFile, line, msg)
	} else {
		printer.Printf("%s:%d: %s\n", displayFile, line, msg)
	}
}

func Callln(s ...any) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		printer.Print("⚠️ [?:?]: " + printer.Sprintf("%s\n", s...))
		return
	}

	msg := printer.Sprintln(s...)
	displayFile := execution.GetFileDisplayPath(file, "")

	if execution.SupportsTerminalHyperlink {
		fullPath := "file://" + file
		printer.Printf(ansi, fullPath, displayFile, line, msg)
	} else {
		printer.Printf("%s:%d: %s\n", displayFile, line, msg)
	}
}

func SCall(s ...any) string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return printer.Sprintf("⚠️ [?:?]: %s\n", s...)
	}

	msg := printer.Sprint(s...)
	displayFile := execution.GetFileDisplayPath(file, "")

	if execution.SupportsTerminalHyperlink {
		fullPath := "file://" + file
		return printer.Sprintf(ansi, fullPath, displayFile, line, msg)
	} else {
		return printer.Sprintf("%s:%d: %s\n", displayFile, line, msg)
	}
}

func SCallf(format string, args ...any) string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "⚠️ [?:?]: " + printer.Sprintf(format, args...)
	}

	msg := printer.Sprintf(format, args...)
	displayFile := execution.GetFileDisplayPath(file, "")
	if execution.SupportsTerminalHyperlink {
		fullPath := "file://" + file
		return printer.Sprintf(ansi, fullPath, displayFile, line, msg)
	} else {
		return printer.Sprintf("%s:%d: %s\n", displayFile, line, msg)
	}
}

func SCallln(s ...any) string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "⚠️ [?:?]: " + printer.Sprintln(s...)
	}

	msg := printer.Sprintln(s...)
	displayFile := execution.GetFileDisplayPath(file, "")

	if execution.SupportsTerminalHyperlink {
		fullPath := "file://" + file
		return printer.Sprintf(ansi, fullPath, displayFile, line, msg)
	} else {
		return printer.Sprintf("%s:%d: %s\n", displayFile, line, msg)
	}
}

func Join(sep string, s ...string) string {
	return strings.Join(s, sep)
}
