// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package assertlite

func fail(t HelperTesting, defaultMsg string, msgAndArgs ...any) {
	t.Helper()
	// If msgAndArgs contains a format string (first argument), pass it to Fatalf along with additional args
	if len(msgAndArgs) > 0 {
		// Ensure the first argument is a string format
		if format, ok := msgAndArgs[0].(string); ok {
			t.Fatalf(format, msgAndArgs[1:]...)
		} else {
			// If no format string is found, fall back to the default error message
			t.Fatal(defaultMsg)
		}
	} else {
		// No arguments, use the default message
		t.Fatal(defaultMsg)
	}
}
