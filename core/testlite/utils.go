// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testlite

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/core/domain/function"
)

// LogOnPanic logs a message if the function panics
func LogOnPanic(t LoggerTesting, f func(), msg string, args ...any) {
	defer func() {
		if r := recover(); r != nil {
			t.Helper()
			t.Fatalf(msg, args...)
		}
	}()
	f()
}

// CloseOrLog closes a resource and logs any close error without failing the test.
// The optional name is used as the resource label in the log line.
func CloseOrLog(t LoggerTesting, c interface{ Close() error }, name ...string) {
	if err := c.Close(); err != nil {
		t.Helper()
		label := "resource"
		if len(name) > 0 {
			label = name[0]
		} else if n, ok := c.(interface{ Name() string }); ok {
			// Priority 2: Does the object have a Name() method? (like *os.File)
			label = n.Name()
		}

		t.Logf("Error closing %s: %v", label, err)
	}
}

// CloseOrFail closes a resource and reports failures through the test handle.
// It supports both Close() error and Close(context.Context) error. If the value
// also implements Done() <-chan struct{}, CloseOrFail waits up to 5 seconds for
// shutdown completion.
func CloseOrFail(t HelperTesting, c any, name ...string) {
	t.Helper()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	label := "resource"
	if len(name) > 0 {
		label = name[0]
	} else if n, ok := c.(interface{ Name() string }); ok {
		label = n.Name()
	}

	// Close
	var err error
	switch v := c.(type) {
	case interface{ Close(context.Context) error }:
		err = v.Close(ctx)
	case interface{ Close() error }:
		err = v.Close()
	default:
		t.Fatalf("type %T has no Close method", c)
	}

	if err != nil {
		t.Errorf("close failed for %s: %v", label, err)
		return // If close failed, waiting for Done might hang forever
	}

	// Wait for Shutdown signal
	if w, ok := c.(interface{ Done() <-chan struct{} }); ok {
		select {
		case <-w.Done():
			// Success: Resource fully shut down
		case <-ctx.Done():
			t.Errorf("timeout waiting for %s to shutdown completely", label)
		}
	}
}

// TrackClose registers a cleanup callback that calls CloseOrFail on test teardown.
func TrackClose(t *testing.T, c interface{ Close() error }, name ...string) {
	t.Cleanup(func() {
		CloseOrFail(t, c, name...)
	})
}

// SkipOnPanic skips the test if the function panics
func SkipOnPanic(t SkippableTesting, f function.Runnable, message string) {
	defer func() {
		if r := recover(); r != nil {
			skipMessage := fmt.Sprintf("%s: panic message: %v\n", message, r)
			t.Skip(skipMessage)
		}
	}()
	f()
}

func SkipOnPanicf(t SkippableTesting, f function.Runnable, message string, args ...any) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			msg := fmt.Sprintf(message, args...)
			pnc := fmt.Sprintf("%s: panic message: %v\n", msg, r)
			t.Skip(pnc)
		}
	}()
	f()
}
