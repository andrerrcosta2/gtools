// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtests

import (
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
)

// SkipOnPanic skips the test if the function panics
func SkipOnPanic(t SkippableTesting, f functions.Runnable, message string) {
	defer func() {
		if r := recover(); r != nil {
			skipMessage := fmx.Sprintf("%s: panic message: %v\n", message, r)
			t.Skip(skipMessage)
		}
	}()
	f()
}

func SkipOnPanicf(t SkippableTesting, f functions.Runnable, message string, args ...any) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			msg := fmx.Sprintf(message, args...)
			pnc := fmx.Sprintf("%s: panic message: %v\n", msg, r)
			t.Skip(pnc)
		}
	}()
	f()
}
