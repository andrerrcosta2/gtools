// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testlite

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
