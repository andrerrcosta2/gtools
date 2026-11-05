// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package chrono

import "time"

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

func Now() Time {
	return Time(time.Now())
}

// Ns converts nanoseconds to a Duration.
func Ns(d int) Duration {
	return Duration(d) * Nanosecond
}

// Ms converts milliseconds to a Duration.
//
// Ms is a std function to convert a number of milliseconds to a Duration.
// This is useful when working with APIs that require a Duration but the
// duration is specified in milliseconds.
//
// @param d the number of milliseconds to convert
// @return the equivalent Duration
func Ms(d int) Duration {
	return Duration(d) * Millisecond
}

// Sec converts seconds to a Duration.
//
// Sec is a std function to convert a number of seconds to a Duration.
// This is useful when working with APIs that require a Duration but the
// duration is specified in seconds.
//
// @param d the number of seconds to convert
// @return the equivalent Duration
func Sec(d int) Duration {
	return Duration(d) * Second
}

// Min converts minutes to a Duration.
//
// Min is a std function to convert a number of minutes to a Duration.
// This is useful when working with APIs that require a Duration but the
// duration is specified in minutes.
//
// @param d the number of minutes to convert
// @return the equivalent Duration
func Min(d int) Duration {
	return Duration(d) * Minute
}

// Hours converts hours to a Duration.
//
// Hours is intended for coarse-grained scheduling intervals where expressing
// a value in whole hours improves readability.
func Hours(d int) Duration {
	return Duration(d) * Hour
}

// Days converts days to a Duration.
//
// A day is normalized to exactly 24 hours, which is usually appropriate for
// scheduler intervals and retry windows.
func Days(d int) Duration {
	return Hours(d * 24)
}

// Weeks converts weeks to a Duration.
//
// A week is normalized to exactly 7 days.
func Weeks(d int) Duration {
	return Days(d * 7)
}

// Zero is a zero-length duration.
var Zero = Duration(0)
