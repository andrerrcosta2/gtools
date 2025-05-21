// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testingtools

import (
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/gtests"
)

// shouldLog returns true if the logging level is satisfied, false otherwise.
// It takes a FailureLoggableTesting and a LoggerStrategy as parameters.
// The returned boolean indicates whether the test should log or not.
func shouldLog(t gtests.FailureLoggableTesting, level gtests.LoggerStrategy) bool {
	// Switch based on the logger level
	switch level {
	case testlogs.ShowAfterTests:
		// Log on call
		return true
	case testlogs.OnErrors:
		// Log on errors
		return t.Failed()
	case testlogs.Default, testlogs.OnFailure:
		// Log on failure
		return t.Failed() || t.Skipped()
	default:
		// Log on nothing
		return false
	}
}
