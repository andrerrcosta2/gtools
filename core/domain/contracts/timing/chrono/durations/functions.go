// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package durations

import "github.com/andrerrcosta2/gtools/core/domain/contracts/chrono"

// Ns converts nanoseconds to a chrono.Duration.
func Ns(d int) chrono.Duration {
	// Multiply the number of nanoseconds by the engine.Nanosecond constant
	// to create a engine.Duration of the same length.
	return chrono.Duration(d) * chrono.Nanosecond
}

// Ms converts milliseconds to a chrono.Duration.
//
// Ms is a std function to convert a number of milliseconds to a chrono.Duration.
// This is useful when working with APIs that require a chrono.Duration but the
// duration is specified in milliseconds.
//
// @param d the number of milliseconds to convert
// @return the equivalent chrono.Duration
func Ms(d int) chrono.Duration {
	// Multiply the number of milliseconds by the engine.Millisecond constant
	// to create a engine.Duration of the same length.
	return chrono.Duration(d) * chrono.Millisecond
}

// Sec converts seconds to a chrono.Duration.
//
// Sec is a std function to convert a number of seconds to a chrono.Duration.
// This is useful when working with APIs that require a chrono.Duration but the
// duration is specified in seconds.
//
// @param d the number of seconds to convert
// @return the equivalent chrono.Duration
func Sec(d int) chrono.Duration {
	// Multiply the number of seconds by the engine.Second constant
	// to create a engine.Duration of the same length.
	return chrono.Duration(d) * chrono.Second
}

// Min converts minutes to a chrono.Duration.
//
// Min is a std function to convert a number of minutes to a chrono.Duration.
// This is useful when working with APIs that require a chrono.Duration but the
// duration is specified in minutes.
//
// @param d the number of minutes to convert
// @return the equivalent chrono.Duration
func Min(d int) chrono.Duration {
	// Multiply the number of minutes by the engine.Minute constant
	// to create a engine.Duration of the same length.
	return chrono.Duration(d) * chrono.Minute
}

// Hours converts hours to a chrono.Duration.
//
// Hours is intended for coarse-grained scheduling intervals where expressing
// a value in whole hours improves readability.
func Hours(d int) chrono.Duration {
	return chrono.Duration(d) * chrono.Hour
}

// Days converts days to a chrono.Duration.
//
// A day is normalized to exactly 24 hours, which is usually appropriate for
// scheduler intervals and retry windows.
func Days(d int) chrono.Duration {
	return chrono.Duration(d) * 24 * chrono.Hour
}

// Weeks converts weeks to a chrono.Duration.
//
// A week is normalized to exactly 7 days.
func Weeks(d int) chrono.Duration {
	return Days(d * 7)
}

// Zero is a zero-length duration.
var Zero = chrono.Duration(0)
