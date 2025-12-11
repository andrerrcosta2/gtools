// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package durations

import "time"

// Ns converts nanoseconds to time.Duration
func Ns(d int) time.Duration {
	// Multiply the number of nanoseconds by the time.Nanosecond constant
	// to create a time.Duration of the same length.
	return time.Duration(d) * time.Nanosecond
}

// Ms converts milliseconds to time.Duration
//
// Ms is a std function to convert a number of milliseconds to a time.Duration.
// This is useful when working with APIs that require a time.Duration but the
// duration is specified in milliseconds.
//
// @param d the number of milliseconds to convert
// @return the equivalent time.Duration
func Ms(d int) time.Duration {
	// Multiply the number of milliseconds by the time.Millisecond constant
	// to create a time.Duration of the same length.
	return time.Duration(d) * time.Millisecond
}

// Sec converts seconds to time.Duration
//
// Sec is a std function to convert a number of seconds to a time.Duration.
// This is useful when working with APIs that require a time.Duration but the
// duration is specified in seconds.
//
// @param d the number of seconds to convert
// @return the equivalent time.Duration
func Sec(d int) time.Duration {
	// Multiply the number of seconds by the time.Second constant
	// to create a time.Duration of the same length.
	return time.Duration(d) * time.Second
}

// Min converts minutes to time.Duration
//
// Min is a std function to convert a number of minutes to a time.Duration.
// This is useful when working with APIs that require a time.Duration but the
// duration is specified in minutes.
//
// @param d the number of minutes to convert
// @return the equivalent time.Duration
func Min(d int) time.Duration {
	// Multiply the number of minutes by the time.Minute constant
	// to create a time.Duration of the same length.
	return time.Duration(d) * time.Minute
}
