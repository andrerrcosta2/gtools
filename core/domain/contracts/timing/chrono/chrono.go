// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package chrono

import (
	"time"
)

// Clock abstracts time for deterministic,
type Clock interface {
	// NewTicker returns a new Ticker containing a channel that will send the
	// current time on the channel after each tick.
	// The period of the ticks is specified by the duration argument.
	// The ticker will adjust the time interval or drop ticks to make
	// up for slow receivers. The duration d must be greater than zero;
	// if not, NewTicker will panic.
	NewTicker(d Duration) Ticker
	// Now returns the current local time.
	Now() Time
	// Since returns the duration since the given time and now
	Since(t Time) Duration
	// Until returns the duration between now and the given time
	Until(t Time) Duration
	// PeriodsSince returns how manu periods have elapsed since the given time
	PeriodsSince(since Time, period Duration) int
	// PeriodsUntil returns how many periods from now exists until the given time
	PeriodsUntil(until Time, period Duration) int
}

type AdvanceableClock interface {
	Clock
	// Advance advances the clock "d" duration from now
	Advance(d Duration)
	// AdvanceTo advances the clock to the exact time.
	// if the time is before the current clock it doesn't perform any operation
	AdvanceTo(t Time)
}

type Ticker interface {
	C() <-chan Time
	Stop()
}

type Timer interface {
	Reset(d Duration)
	Stop()
}

type TimerFactory interface {
	NewTimer(func()) Timer
}

func Sleep(d Duration) {
	time.Sleep(d.ToStd())
}

// FromStdMonth converts a time.Month to chrono.Month
func FromStdMonth(m time.Month) Month {
	return Month(m)
}

// A Month specifies a month of the year (January = 1, ...).
type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

// String returns the English name of the month ("January", "February", ...).
func (m Month) String() string {
	return time.Month(m).String()
}

// ToStd converts the chrono.Month to time.Month
func (m Month) ToStd() time.Month {
	return time.Month(m)
}

// FromStdWeekday converts a time.Weekday to chrono.Weekday
func FromStdWeekday(m time.Weekday) Weekday {
	return Weekday(m)
}

// A Weekday specifies a day of the week (Sunday = 0, ...).
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (w Weekday) String() string {
	return time.Weekday(w).String()
}

// ToStd converts the chrono.Weekday to time.Weekday
func (w Weekday) ToStd() time.Weekday {
	return time.Weekday(w)
}

// Date returns the Time corresponding to
//
//	yyyy-mm-dd hh:mm:ss + nsec nanoseconds
//
// in the appropriate zone for that time in the given location.
//
// The month, day, hour, min, sec, and nsec values may be outside
// their usual ranges and will be normalized during the conversion.
// For example, October 32 converts to November 1.
//
// A daylight savings time transition skips or repeats times.
// For example, in the United States, March 13, 2011 2:15am never occurred,
// while November 6, 2011 1:15am occurred twice. In such cases, the
// choice of time zone, and therefore the time, is not well-defined.
// Date returns a time that is correct in one of the two zones involved
// in the transition, but it does not guarantee which.
//
// Date panics if loc is nil.
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	return FromStdTime(time.Date(year, time.Month(month), day, hour, min, sec, nsec, loc))
}

// FromStdTime converts a standard time.Time to chrono.Time
func FromStdTime(t time.Time) Time {
	return Time(t)
}

type Time time.Time

// Add returns the time t+d
func (t Time) Add(d Duration) Time {
	return Time(t.ToStd().Add(d.ToStd()))
}

// After reports whether the time instant t is after u
func (t Time) After(o Time) bool {
	return time.Time(t).After(o.ToStd())
}

// Before reports whether the time instant t is before u.
func (t Time) Before(o Time) bool {
	return time.Time(t).Before(o.ToStd())
}

// Clamp clamps the times between min and max
func (t Time) Clamp(min, max Time) Time {
	if t.Before(min) {
		return min
	}
	if t.After(max) {
		return max
	}
	return t
}

// Compare returns -1 if the current time is less than the parameter time
// returns 1 if it is higher and 0 if they are exactly the same time
func (t Time) Compare(o Time) int {
	if t.Before(o) {
		return -1
	}
	if t.After(o) {
		return 1
	}
	return 0
}

// Equal reports whether t and u represent the same time instant.
// Two times can be equal even if they are in different locations.
// For example, 6:00 +0200 and 4:00 UTC are Equal.
// See the documentation on the Time type for the pitfalls of using == with
// Time values; most code should use Equal instead.
func (t Time) Equal(o Time) bool {
	return time.Time(t).Equal(o.ToStd())
}

// IsZero reports whether t represents the zero time instant,
// January 1, year 1, 00:00:00 UTC.
func (t Time) IsZero() bool {
	return time.Time(t).IsZero()
}

// Less is a domain semantics wrapper for time.Time Before
// that considers before a smaller time for ordering purposes
func (t Time) Less(o Time) bool {
	return time.Time(t).Before(o.ToStd())
}

// NextAfter returns the first occurrence of the periodic sequence
// starting at t with the given interval that is strictly after now.
//
// If t is already after now, t is returned unchanged.
//
// Non-positive intervals are normalized to Nanosecond.
//
// Example:
//
//	start = 10:00
//	interval = 5m
//
//	now = 10:12 -> 10:15
//	now = 10:15 -> 10:20
func (t Time) NextAfter(now Time, interval Duration) Time {
	if interval <= 0 {
		interval = Nanosecond
	}

	// If still in the future, keep original
	if t.After(now) {
		return t
	}

	delta := now.Sub(t)
	steps := delta.DivDuration(interval)

	return t.Add(interval.Mul(steps + 1))
}

// NextAfterOrNow returns the first occurrence of the periodic sequence
// starting at t with the given interval that is at or after now.
//
// If now falls exactly on an occurrence, now is returned.
// If t is already after now, t is returned unchanged.
//
// Non-positive intervals are normalized to Nanosecond.
//
// Example:
//
//	start = 10:00
//	interval = 5m
//
//	now = 10:12 -> 10:15
//	now = 10:15 -> 10:15
func (t Time) NextAfterOrNow(now Time, interval Duration) Time {
	if interval <= 0 {
		interval = Nanosecond
	}

	// Future → keep
	if t.After(now) {
		return t
	}

	delta := now.Sub(t)
	steps := delta.DivDuration(interval)

	next := t.Add(interval.Mul(steps))

	// if we are exactly on or behind, execute now
	if !next.After(now) {
		return now
	}

	return next
}

// PeriodsSince returns how manu periods have elapsed since the given time
func (t Time) PeriodsSince(since Time, period Duration) int {
	interval := t.Since(since)
	return int(interval / period)
}

// PeriodsUntil returns how many periods from now exists until the given time
func (t Time) PeriodsUntil(until Time, period Duration) int {
	interval := t.Until(until)
	return int(interval / period)
}

// Round returns the result of rounding t to the nearest multiple of d (since the zero time).
// The rounding behavior for halfway values is to round up.
// If d <= 0, Round returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Round operates on the time as an absolute duration since the
// zero time; it does not operate on the presentation form of the
// time. Thus, Round(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
func (t Time) Round(d Duration) Time {
	return Time(t.ToStd().Round(d.ToStd()))
}

// Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
// value that can be stored in a [Duration], the maximum (or minimum) duration
// will be returned.
// To compute t-d for a duration d, use t.Add(-d).
func (t Time) Sub(o Time) Duration {
	return Duration(t.ToStd().Sub(o.ToStd()))
}

func (t Time) SubDuration(d Duration) Time {
	return Time(t.ToStd().Add(-d.ToStd()))
}

// Since returns the time elapsed between it and the given time.
// it is an alias for t.Sub(o)
func (t Time) Since(o Time) Duration {
	return Duration(t.ToStd().Sub(o.ToStd()))
}

func (t Time) Until(o Time) Duration {
	return Duration(o.ToStd().Sub(t.ToStd()))
}

// ToStd converts the chrono.Time to the standard time.Time
func (t Time) ToStd() time.Time {
	return time.Time(t)
}

// Truncate returns the result of rounding t down to a multiple of d (since the zero time).
// If d <= 0, Truncate returns t stripped of any monotonic clock reading but otherwise unchanged.
//
// Truncate operates on the time as an absolute duration since the
// zero time; it does not operate on the presentation form of the
// time. Thus, Truncate(Hour) may return a time with a non-zero
// minute, depending on the time's Location.
func (t Time) Truncate(d Duration) Time {
	return Time(t.ToStd().Truncate(d.ToStd()))
}

// Unix returns t as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC. The result does not depend on the
// location associated with t.
// Unix-like operating systems often record time as a 32-bit
// count of seconds, but since the method here returns a 64-bit
// value it is valid for billions of years into the past or future.
func (t Time) Unix() int64 {
	return t.ToStd().Unix()
}

// UnixMilli returns t as a Unix time, the number of milliseconds elapsed since
// January 1, 1970 UTC. The result is undefined if the Unix time in
// milliseconds cannot be represented by an int64 (a date more than 292 million
// years before or after 1970). The result does not depend on the
// location associated with t.
func (t Time) UnixMilli() int64 {
	return t.ToStd().UnixMilli()
}

// UnixMicro returns t as a Unix time, the number of microseconds elapsed since
// January 1, 1970 UTC. The result is undefined if the Unix time in
// microseconds cannot be represented by an int64 (a date before year -290307 or
// after year 294246). The result does not depend on the location associated
// with t.
func (t Time) UnixMicro() int64 {
	return t.ToStd().UnixMicro()
}

// AppendBinary implements the [encoding.BinaryAppender] interface.
func (t Time) AppendBinary(b []byte) ([]byte, error) {
	return t.ToStd().AppendBinary(b)
}

// AppendText implements the [encoding.TextAppender] interface.
// The time is formatted in RFC 3339 format with sub-second precision.
// If the timestamp cannot be represented as valid RFC 3339
// (e.g., the year is out of range), then an error is returned.
func (t Time) AppendText(b []byte) ([]byte, error) {
	return t.ToStd().AppendText(b)
}

// Date returns the year, month, and day in which t occurs.
func (t Time) Date() (year int, month Month, day int) {
	y, m, d := t.ToStd().Date()
	return y, Month(m), d
}

// Year returns the year in which t occurs.
func (t Time) Year() int {
	return t.ToStd().Year()
}

// Month returns the month of the year specified by t.
func (t Time) Month() Month {
	return Month(t.ToStd().Month())
}

// Day returns the day of the month specified by t.
func (t Time) Day() int {
	return t.ToStd().Day()
}

// Weekday returns the day of the week specified by t.
func (t Time) Weekday() Weekday {
	return Weekday(t.ToStd().Weekday())
}

// ISOWeek returns the ISO 8601 year and week number in which t occurs.
// Week ranges from 1 to 53. Jan 01 to Jan 03 of year n might belong to
// week 52 or 53 of year n-1, and Dec 29 to Dec 31 might belong to week 1
// of year n+1.
func (t Time) ISOWeek() (year int, week int) {
	return t.ToStd().ISOWeek()
}

// Clock returns the hour, minute, and second within the day specified by t.
func (t Time) Clock() (hour, min, sec int) {
	return t.ToStd().Clock()
}

// Hour returns the hour within the day specified by t, in the range [0, 23].
func (t Time) Hour() int {
	return t.ToStd().Hour()
}

// Minute returns the minute offset within the hour specified by t, in the range [0, 59].
func (t Time) Minute() int {
	return t.ToStd().Minute()
}

// Second returns the second offset within the minute specified by t, in the range [0, 59].
func (t Time) Second() int {
	return t.ToStd().Second()
}

// Nanosecond returns the nanosecond offset within the second specified by t,
// in the range [0, 999999999].
func (t Time) Nanosecond() int {
	return t.ToStd().Nanosecond()
}

// YearDay returns the day of the year specified by t, in the range [1,365] for non-leap years,
// and [1,366] in leap years.
func (t Time) YearDay() int {
	return t.ToStd().YearDay()
}

// UnixNano returns t as a Unix time, the number of nanoseconds elapsed since
// January 1, 1970 UTC. The result is undefined if the Unix time in nanoseconds
// cannot be represented by an int64 (a date before the year 1678 or after 2262).
// Note that this means the result of calling UnixNano on the zero Time is undefined.
// The result does not depend on the location associated with t.
func (t Time) UnixNano() int64 {
	return time.Time(t).UnixNano()
}

// UTC returns t with the location set to UTC.
func (t Time) UTC() Time {
	return Time(t.ToStd().UTC())
}

// Local returns t with the location set to local time.
func (t Time) Local() Time {
	return Time(t.ToStd().Local())
}

// In returns a copy of t representing the same time instant, but
// with the copy's location information set to loc for display
// purposes.
//
// In panics if loc is nil.
func (t Time) In(loc *time.Location) Time {
	return Time(t.ToStd().In(loc))
}

// Location returns the time zone information associated with t.
func (t Time) Location() *time.Location {
	return t.ToStd().Location()
}

// Zone computes the time zone in effect at time t, returning the abbreviated
// name of the zone (such as "CET") and its offset in seconds east of UTC.
func (t Time) Zone() (name string, offset int) {
	return t.ToStd().Zone()
}

// ZoneBounds returns the bounds of the time zone in effect at time t.
// The zone begins at start and the next zone begins at end.
// If the zone begins at the beginning of time, start will be returned as a zero Time.
// If the zone goes on forever, end will be returned as a zero Time.
// The Location of the returned times will be the same as t.
func (t Time) ZoneBounds() (start, end Time) {
	return t.ZoneBounds()
}

// FromStdDuration converts a time.Duration to chrono.Duration
func FromStdDuration(d time.Duration) Duration {
	return Duration(d)
}

type Duration time.Duration

// Abs returns the absolute value of d.
// As a special case, Duration([math.MinInt64]) is converted to Duration([math.MaxInt64]),
// reducing its magnitude by 1 nanosecond.
func (d Duration) Abs() Duration {
	return Duration(time.Duration(d).Abs())
}

// Add returns the duration d+d.
func (d Duration) Add(o Duration) Duration {
	return Duration(time.Duration(d) + time.Duration(o))
}

// Div returns the duration d divided by an int.
func (d Duration) Div(n int) Duration {
	if n == 0 {
		panic("engine: division by zero")
	}
	return Duration(time.Duration(d) / time.Duration(n))
}

// DivDuration returns the duration d/d.
func (d Duration) DivDuration(o Duration) int {
	if o == 0 {
		panic("engine: division by zero duration")
	}
	return int(time.Duration(d) / time.Duration(o))
}

// IsNegative reports whether the duration is a negative value
func (d Duration) IsNegative() bool {
	return d < 0
}

// Mul returns the duration d multiiplied by an int.
func (d Duration) Mul(n int) Duration {
	return Duration(time.Duration(d) * time.Duration(n))
}

// Sub returns the duration d-d
func (d Duration) Sub(o Duration) Duration {
	return Duration(time.Duration(d) - time.Duration(o))
}

// ToStd converts the Duration to the standard time.Duration
func (d Duration) ToStd() time.Duration {
	return time.Duration(d)
}

// String returns a string representing the duration in the form "72h3m0.5s".
// Leading zero units are omitted. As a special case, durations less than one
// second format use a smaller unit (milli-, micro-, or nanoseconds) to ensure
// that the leading digit is non-zero. The zero duration formats as 0s.
func (d Duration) String() string {
	return time.Duration(d).String()
}

// Nanoseconds returns the duration as an integer nanosecond count.
func (d Duration) Nanoseconds() int64 { return int64(d) }

// Microseconds returns the duration as an integer microsecond count.
func (d Duration) Microseconds() int64 { return int64(d) / 1e3 }

// Milliseconds returns the duration as an integer millisecond count.
func (d Duration) Milliseconds() int64 { return int64(d) / 1e6 }

func (d Duration) Seconds() float64 {
	sec := d / Second
	nsec := d % Second
	return float64(sec) + float64(nsec)/1e9
}

// Minutes returns the duration as a floating point number of minutes.
func (d Duration) Minutes() float64 {
	min := d / Minute
	nsec := d % Minute
	return float64(min) + float64(nsec)/(60*1e9)
}

// Hours returns the duration as a floating point number of hours.
func (d Duration) Hours() float64 {
	hour := d / Hour
	nsec := d % Hour
	return float64(hour) + float64(nsec)/(60*60*1e9)
}

// Truncate returns the result of rounding d toward zero to a multiple of m.
// If m <= 0, Truncate returns d unchanged.
func (d Duration) Truncate(m Duration) Duration {
	if m <= 0 {
		return d
	}
	return d - d%m
}

// Round returns the result of rounding d to the nearest multiple of m.
// The rounding behavior for halfway values is to round away from zero.
// If the result exceeds the maximum (or minimum)
// value that can be stored in a [Duration],
// Round returns the maximum (or minimum) duration.
// If m <= 0, Round returns d unchanged.
func (d Duration) Round(m Duration) Duration {
	return Duration(time.Duration(d).Round(time.Duration(m)))
}
