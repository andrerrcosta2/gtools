// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package agendas

import (
	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/agenda"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/chrono"
	"github.com/andrerrcosta2/gtools/core/domain/contracts/timing/schedulable"
	"github.com/andrerrcosta2/gtools/core/domain/function"
)

// OfTasks creates a scheduler agenda backed by a bucket wheel.
//
// bucketSpan defines the time interval represented by each bucket.
//
// The agenda is resilient to scheduler jitter. Calls to Due(now) do not need
// to occur exactly at bucket boundaries. Instead, the agenda tracks the last
// processed instant and drains every bucket whose interval has completely
// elapsed since then.
//
// For example, with:
//
//	bucketSpan = 100ms
//
// successive calls:
//
//	Due(150ms) -> drains [0ms,100ms)
//	Due(300ms) -> drains [100ms,200ms) and [200ms,300ms)
//
// The 50ms delay observed on the first call is automatically recovered on the
// second one. Temporary drift therefore accumulates only until the next Due()
// invocation and does not permanently shift the schedule.
//
// In practice, bucketSpan determines the scheduler resolution. Tasks may be
// delayed by at most one bucket interval, depending on when Due() is called.
//
// For best accuracy, the scheduler driving Due() should use a tick no larger
// than bucketSpan. Larger ticks remain correct but increase the maximum
// scheduling latency.
func OfTasks(start chrono.Time, bucketSpan chrono.Duration, buckets int) agenda.Chrono[schedulable.Event[chrono.Time]] {
	cmp := comparator.Build[schedulable.Event[chrono.Time], int64]().
		Compare(func(a, b schedulable.Event[chrono.Time]) int {
			if a.At().Less(b.At()) {
				return -1
			}
			if b.At().Less(a.At()) {
				return 1
			}
			return 0
		}).
		Key(func(t schedulable.Event[chrono.Time]) int64 {
			return t.At().UnixNano()
		}).Build()
	wheel := newBucketWheel(cmp, start.UnixNano(), bucketSpan.Nanoseconds(), buckets)
	return &task{
		wheel:     wheel,
		span:      bucketSpan,
		processed: start,
	}
}

type task struct {
	wheel     *bucketWheel[schedulable.Event[chrono.Time], int64]
	processed chrono.Time
	span      chrono.Duration
}

func (a *task) Push(t schedulable.Event[chrono.Time]) {
	a.wheel.Push(t)
}

func (a *task) Due(now chrono.Time) []schedulable.Event[chrono.Time] {
	out := make([]schedulable.Event[chrono.Time], 0)
	periods := now.PeriodsSince(a.processed, a.span)
	a.processed = now
	for range periods {
		tasks, ok := a.wheel.PopBucket()
		if !ok {
			continue
		}
		out = append(out, tasks...)
	}
	return out
}

// RangeDue invokes fn for every task due since the last call, stopping early
// if fn returns false. Returns false if iteration was stopped early — any
// undrained periods remain pending and will be retried on the next call.
func (a *task) RangeDue(now chrono.Time, fn function.Predicate[schedulable.Event[chrono.Time]]) bool {
	for range now.PeriodsSince(a.processed, a.span) {
		if !a.wheel.PopRange(fn) {
			return false
		}
		a.processed = a.processed.Add(a.span)
	}
	return true
}

func (a *task) Drain(fn function.Predicate[schedulable.Event[chrono.Time]]) int {
	count := 0
	for a.wheel.Len() > 0 {
		a.wheel.PopRange(func(t schedulable.Event[chrono.Time]) bool {
			count++
			return fn(t)
		})
		a.processed = a.processed.Add(a.span)
	}
	return count
}
