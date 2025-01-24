// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package gtests

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
}

type LoggableTesting interface {
	HelperTesting
	// Log methods
	//
	// Log logs a message to the console.
	// It takes a variable number of arguments, and they are passed to the
	// fmt.Println function to print the message to the console.
	Log(args ...interface{})
	// Logf logs a message to the console.
	// The message is formatted using the fmt.Printf function.
	// The arguments are passed to the fmt.Printf function to format the message.
	Logf(format string, args ...interface{})
	// Fatal methods
	//
	// Fatal logs a message to the console and marks the test as failed.
	// The message is passed to the fmt.Println function to print it to the console.
	// After the message is printed, the test is stopped and the program is terminated.
	Fatal(args ...interface{})
	// Fatalf logs a message to the console and marks the test as failed.
	// The message is formatted using the fmt.Printf function.
	// The arguments are passed to the fmt.Printf function to format the message.
	// After the message is printed, the test is stopped and the program is terminated.
	Fatalf(format string, args ...interface{})
}

type FailureLoggableTesting interface {
	LoggableTesting
	// Fail marks the test as failed but continues executing the test.
	//
	// It takes no arguments.
	Fail()
	// FailNow marks the test as failed and stops the test execution.
	//
	// It takes no arguments.
	// The test is stopped and the program is terminated.
	FailNow()
	// Failed returns true if the test has failed, false otherwise.
	//
	// It takes no arguments.
	Failed() bool
	// Skipped returns true if the test was skipped, false otherwise.
	//
	// It takes no arguments.
	Skipped() bool
}
