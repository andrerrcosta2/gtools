// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

//go:build tt

package assertlite

type HelperTesting interface {
	// Helper is a method that is used to write a message to the console.
	// It calls the testing.T.Helper method to mark the message as a helper.
	// It takes a variable number of arguments, and they are passed to the
	// fmt.Println function to print the message to the console.
	Helper()
	// Error methods
	//
	// Error logs a message to the console and marks the test as failed.
	// The message is passed to the fmt.Println function to print it to the console.
	Error(args ...interface{})
	// Errorf logs a message to the console and marks the test as failed.
	// The message is formatted using the fmt.Printf function.
	// The arguments are passed to the fmt.Printf function to format the message.
	Errorf(format string, args ...interface{})

	// FailNow marks the test as failed and stops the test execution immediately.
	// It is typically used in a testing environment to signal a critical failure.
	FailNow()

	Fatal(args ...any)
	Fatalf(format string, args ...any)
}
