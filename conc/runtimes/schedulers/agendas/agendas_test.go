// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package agendas

import (
	"testing"
	"time"

	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/chrono"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/schedulable"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
)

// ref is the reference instant for all agenda tests.
//
// We deliberately anchor to the Unix epoch (1970-01-01 UTC) rather than
// time.Time{} (year 1) because the bucket wheel keys on UnixNano(), and
// the zero time's UnixNano() is undefined per the stdlib docs.
func ref() chrono.Time {
	return chrono.FromStdTime(time.Unix(0, 0).UTC())
}

// after returns the instant ref()+n milliseconds.
func after(n int64) chrono.Time {
	return ref().Add(ms(n))
}

// ms builds a chrono.Duration of n milliseconds.
//
// chrono.Duration is a named type over time.Duration, so a plain
// integer conversion would be nanoseconds, not milliseconds. Going
// through time.Millisecond keeps the intent explicit.
func ms(n int64) chrono.Duration {
	return chrono.FromStdDuration(time.Duration(n) * time.Millisecond)
}

type event struct {
	at      chrono.Time
	payload string
}

func (e event) At() chrono.Time { return e.at }

func ev(at chrono.Time, payload string) event {
	return event{at: at, payload: payload}
}

func payloads(evs []schedulable.Event[chrono.Time]) []string {
	out := make([]string, 0, len(evs))
	for _, e := range evs {
		out = append(out, e.(event).payload)
	}
	return out
}

// TestOfTasks_Construction
//
// Constructs the agenda with valid bucket geometry.
//
// We should assert:
//  1. The returned agenda is non-nil
//  2. An empty agenda drains nothing and completes iteration cleanly
//  3. The agenda is usable immediately without any prior tick
func TestOfTasks_Construction(t *testing.T) {
	t.Run("returns a usable agenda", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)
		assertlite.NotNil(t, a)
	})

	t.Run("empty agenda drains nothing", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		got := a.Due(after(1000))
		assertlite.Equals(t, 0, len(got))

		ok := a.RangeDue(after(1000), func(schedulable.Event[chrono.Time]) bool {
			return true
		})
		assertlite.True(t, ok)

		assertlite.Equals(t, 0, a.Drain(func(schedulable.Event[chrono.Time]) bool {
			return true
		}))
	})
}

// TestOfTasks_ExactTick ("sync" analogue)
//
// The scheduler calls Due() exactly once per bucket span, at the
// boundary. This is the ideal, drift-free regime and mirrors the
// synchronous case in emission-driven tests: the caller is perfectly
// aligned with the agenda's internal clock.
//
// We should assert:
//  1. Due() returns exactly the tasks whose bucket interval has elapsed
//  2. Tasks scheduled for future buckets are NOT returned early
//  3. Calling Due() again without advancing time yields nothing
//  4. Each bucket boundary releases only its own tasks
func TestOfTasks_ExactTick(t *testing.T) {
	t.Run("drains tasks due in the current bucket", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		a.Push(ev(after(10), "a"))
		a.Push(ev(after(20), "b"))

		got := a.Due(after(100))
		assertlite.EqualSlices(t, []string{"a", "b"}, payloads(got))
	})

	t.Run("does not release future buckets early", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		a.Push(ev(after(50), "now"))
		a.Push(ev(after(150), "later"))

		got := a.Due(after(100))
		assertlite.EqualSlices(t, []string{"now"}, payloads(got))
	})

	t.Run("second Due at the same instant is empty", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		a.Push(ev(after(10), "a"))
		_ = a.Due(after(100))

		got := a.Due(after(100))
		assertlite.Equals(t, 0, len(got))
	})

	t.Run("each boundary releases only its own bucket", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		a.Push(ev(after(10), "b0"))
		a.Push(ev(after(110), "b1"))
		a.Push(ev(after(210), "b2"))

		assertlite.EqualSlices(t, []string{"b0"}, payloads(a.Due(after(100))))
		assertlite.EqualSlices(t, []string{"b1"}, payloads(a.Due(after(200))))
		assertlite.EqualSlices(t, []string{"b2"}, payloads(a.Due(after(300))))
	})
}

// TestOfTasks_JitterRecovery ("async" analogue)
//
// The scheduler calls Due() late or irregularly, as happens when the
// driving tick is larger than bucketSpan or when the runtime stalls.
// This is the asynchronous regime: the caller is not aligned with the
// agenda's internal clock and the agenda must recover missed periods
// without shifting the schedule permanently.
//
// We should assert:
//  1. Skipped buckets are drained on the next Due() call
//  2. Recovered tasks are emitted exactly once
//  3. A single late call can release more than one bucket
//  4. Recovery does not permanently shift subsequent boundaries
func TestOfTasks_JitterRecovery(t *testing.T) {
	t.Run("recovers a skipped bucket on the next call", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		a.Push(ev(after(50), "first"))
		a.Push(ev(after(150), "second"))

		// First call jumps to 150ms: only [0,100) is entirely elapsed.
		got1 := a.Due(after(150))
		assertlite.EqualSlices(t, []string{"first"}, payloads(got1))

		// Second call recovers [100,200).
		got2 := a.Due(after(300))
		assertlite.EqualSlices(t, []string{"second"}, payloads(got2))
	})

	t.Run("single late call releases multiple buckets", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		a.Push(ev(after(10), "b0"))
		a.Push(ev(after(110), "b1"))
		a.Push(ev(after(210), "b2"))

		got := a.Due(after(300))
		assertlite.EqualSlices(t, []string{"b0", "b1", "b2"}, payloads(got))
	})

	t.Run("recovered tasks are emitted exactly once", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)
		a.Push(ev(after(10), "x"))

		seen := map[string]int{}
		for _, e := range a.Due(after(150)) {
			seen[e.(event).payload]++
		}
		for _, e := range a.Due(after(400)) {
			seen[e.(event).payload]++
		}

		assertlite.Equals(t, 1, seen["x"])
	})

	t.Run("recovery does not shift subsequent boundaries", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		a.Push(ev(after(10), "b0"))
		a.Push(ev(after(110), "b1"))
		a.Push(ev(after(210), "b2"))

		// Late first tick recovers b0 only.
		_ = a.Due(after(150))
		// Next tick at exactly 300ms recovers b1 and b2 — proving the
		// schedule was not shifted by the jitter.
		got := a.Due(after(300))
		assertlite.EqualSlices(t, []string{"b1", "b2"}, payloads(got))
	})
}

// ---------------------------------------------------------------------
// OfTasks_RangeDue
// ---------------------------------------------------------------------

// TestOfTasks_RangeDue
//
// RangeDue is the pull-based iterator variant of Due. The caller supplies
// a predicate and can stop iteration early; the agenda guarantees that
// undrained periods remain pending for the next call.
//
// We should assert:
//  1. All due tasks are visited in order when the predicate always passes
//  2. Returning false stops iteration and returns false
//  3. Stopping early leaves undrained periods pending for the next call
//  4. Calling RangeDue with no due tasks returns true without invoking fn
func TestOfTasks_RangeDue(t *testing.T) {
	t.Run("visits all due tasks in order", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		a.Push(ev(after(10), "b0-a"))
		a.Push(ev(after(20), "b0-b"))
		a.Push(ev(after(110), "b1-a"))

		var got []string
		ok := a.RangeDue(after(300), func(e schedulable.Event[chrono.Time]) bool {
			got = append(got, e.(event).payload)
			return true
		})

		assertlite.True(t, ok)
		assertlite.EqualSlices(t, []string{"b0-a", "b0-b", "b1-a"}, got)
	})

	t.Run("returns false when stopped early", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)
		a.Push(ev(after(10), "a"))

		ok := a.RangeDue(after(300), func(schedulable.Event[chrono.Time]) bool {
			return false
		})
		assertlite.False(t, ok)
	})

	t.Run("stopped period remains pending", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		a.Push(ev(after(10), "b0"))
		a.Push(ev(after(110), "b1"))

		// Stop on the very first callback.
		_ = a.RangeDue(after(300), func(schedulable.Event[chrono.Time]) bool {
			return false
		})

		// Next call must still deliver the un-drained work.
		got := a.Due(after(400))
		assertlite.True(t, len(got) >= 1)
	})

	t.Run("no due tasks returns true without invoking fn", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		called := false
		ok := a.RangeDue(after(50), func(schedulable.Event[chrono.Time]) bool {
			called = true
			return true
		})

		assertlite.True(t, ok)
		assertlite.False(t, called)
	})
}

// ---------------------------------------------------------------------
// OfTasks_Drain
// ---------------------------------------------------------------------

// TestOfTasks_Drain
//
// Drain empties the wheel unconditionally, ignoring bucket boundaries.
// It is the "shutdown" path: every remaining task must be delivered
// regardless of its scheduled time.
//
// We should assert:
//  1. Drain returns the number of tasks delivered
//  2. Every remaining task is visited exactly once
//  3. A second Drain returns 0
//  4. Drain visits tasks regardless of their scheduled bucket
func TestOfTasks_Drain(t *testing.T) {
	t.Run("returns the number of tasks drained", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		a.Push(ev(after(10), "a"))
		a.Push(ev(after(200), "b"))
		a.Push(ev(after(500), "c"))

		n := a.Drain(func(schedulable.Event[chrono.Time]) bool { return true })
		assertlite.Equals(t, 3, n)
	})

	t.Run("visits every remaining task exactly once", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		a.Push(ev(after(10), "a"))
		a.Push(ev(after(200), "b"))
		a.Push(ev(after(500), "c"))

		seen := map[string]int{}
		a.Drain(func(e schedulable.Event[chrono.Time]) bool {
			seen[e.(event).payload]++
			return true
		})

		assertlite.Equals(t, 1, seen["a"])
		assertlite.Equals(t, 1, seen["b"])
		assertlite.Equals(t, 1, seen["c"])
	})

	t.Run("second drain is empty", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)
		a.Push(ev(after(10), "a"))

		a.Drain(func(schedulable.Event[chrono.Time]) bool { return true })

		n := a.Drain(func(schedulable.Event[chrono.Time]) bool { return true })
		assertlite.Equals(t, 0, n)
	})

	t.Run("drains tasks scheduled far in the future", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)
		a.Push(ev(after(10_000), "far"))

		var got []string
		a.Drain(func(e schedulable.Event[chrono.Time]) bool {
			got = append(got, e.(event).payload)
			return true
		})

		assertlite.EqualSlices(t, []string{"far"}, got)
	})
}

// ---------------------------------------------------------------------
// OfTasks_BucketOrdering
// ---------------------------------------------------------------------

// TestOfTasks_BucketOrdering
//
// Inside a bucket, tasks are ordered by their At() key. Push order must
// not affect emission order.
//
// We should assert:
//  1. Tasks in the same bucket come out sorted by At()
//  2. Out-of-order pushes are corrected by the internal priority queue
//  3. Ordering holds across bucket boundaries
func TestOfTasks_BucketOrdering(t *testing.T) {
	t.Run("sorts tasks within a bucket by At()", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		a.Push(ev(after(30), "c"))
		a.Push(ev(after(10), "a"))
		a.Push(ev(after(20), "b"))

		got := a.Due(after(100))
		assertlite.EqualSlices(t, []string{"a", "b", "c"}, payloads(got))
	})

	t.Run("ordering holds across buckets", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		a.Push(ev(after(210), "b2"))
		a.Push(ev(after(10), "b0"))
		a.Push(ev(after(110), "b1"))

		got := a.Due(after(300))
		assertlite.EqualSlices(t, []string{"b0", "b1", "b2"}, payloads(got))
	})
}

// ---------------------------------------------------------------------
// OfTasks_LapWrapping
// ---------------------------------------------------------------------

// TestOfTasks_LapWrapping
//
// With 4 buckets of 100ms the wheel covers 400ms per lap. Tasks beyond
// that horizon land on a later lap and must only be released when the
// wheel actually reaches that lap.
//
// We should assert:
//  1. Lap-1 tasks are not released during lap 0
//  2. Lap-1 tasks are released once the wheel wraps
//  3. No task is lost across a lap boundary
func TestOfTasks_LapWrapping(t *testing.T) {
	t.Run("lap-1 tasks are not released during lap 0", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		a.Push(ev(after(10), "lap0"))
		a.Push(ev(after(510), "lap1"))

		got := a.Due(after(100))
		assertlite.EqualSlices(t, []string{"lap0"}, payloads(got))
	})

	t.Run("lap-1 tasks are released after wrap", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		a.Push(ev(after(10), "lap0"))
		a.Push(ev(after(510), "lap1"))

		_ = a.Due(after(100))    // drain lap 0
		got := a.Due(after(600)) // reach lap 1's bucket 1
		assertlite.EqualSlices(t, []string{"lap1"}, payloads(got))
	})

	t.Run("no task is lost across a lap boundary", func(t *testing.T) {
		a := OfTasks(ref(), ms(100), 4)

		for i := int64(0); i < 10; i++ {
			a.Push(ev(after(i*100+10), "t"))
		}

		total := 0
		for step := int64(1); step <= 10; step++ {
			total += len(a.Due(after(step * 100)))
		}
		assertlite.Equals(t, 10, total)
	})
}
