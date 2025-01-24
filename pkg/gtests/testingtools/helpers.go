// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testingtools

import (
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/gtools/functions"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"github.com/andrerrcosta2/gtools/gtests"
)

// shouldLog returns true if the logging level is satisfied, false otherwise.
// It takes a FailureLoggableTesting and a LoggerStrategy as parameters.
// The returned boolean indicates whether the test should log or not.
func shouldLog(t gtests.FailureLoggableTesting, level gtests.LoggerStrategy) bool {
	// Switch based on the logger level
	switch level {
	case gtests.AlwaysPrintLog:
		// Log on call
		return true
	case gtests.LogOnErrors:
		// Log on errors
		return t.Failed()
	case gtests.LogOnFailure:
		// Log on failure
		return t.Failed() || t.Skipped()
	default:
		// Log on nothing
		return false
	}
}

func randUniqueBytesOf[T any](amount int, seeder functions.Function[[]byte, T]) []T {
	var out = make([]T, amount)
	rnd, err := random.UniqueByteSlices(amount, 3, 1, 100)
	if data.IsTaggable[string](err) {
		panic(err)
	}
	rnd.EachN(func(i int, e []byte) {
		out[i] = seeder(e)
	})
	return out
}
