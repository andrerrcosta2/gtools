// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package agendas

import (
	"testing"

	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
)

type item struct {
	key     int64
	payload string
}

func it(key int64, payload string) item {
	return item{key: key, payload: payload}
}

// cmpInt64 builds the CmpKeyer the wheel expects: it orders by key and
// extracts the key from the item.
func cmpInt64() comparator.CmpKeyer[item, int64] {
	return comparator.Build[item, int64]().
		Compare(func(a, b item) int {
			switch {
			case a.key < b.key:
				return -1
			case a.key > b.key:
				return 1
			default:
				return 0
			}
		}).
		Key(func(i item) int64 { return i.key }).
		Build()
}

func itemPayloads(items []item) []string {
	out := make([]string, 0, len(items))
	for _, i := range items {
		out = append(out, i.payload)
	}
	return out
}

// TestBucketWheel_Construction
//
// Constructs the wheel with valid and invalid geometry.
//
// We should assert:
//  1. A valid wheel is returned and starts empty
//  2. Non-power-of-two bucket counts panic
//  3. Non-positive bucket counts panic
//  4. The smallest valid wheel (one bucket) is accepted
func TestBucketWheel_Construction(t *testing.T) {
	t.Run("returns an empty wheel", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		assertlite.NotNil(t, w)
		assertlite.Equals(t, 0, w.Len())
	})

	t.Run("single bucket is valid", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 1)
		assertlite.NotNil(t, w)
		assertlite.Equals(t, 0, w.Len())
	})

	t.Run("panics on non-power-of-two", func(t *testing.T) {
		assertlite.Panics(t, func() {
			_ = newBucketWheel[item, int64](cmpInt64(), 0, 10, 3)
		})
	})

	t.Run("panics on zero or negative buckets", func(t *testing.T) {
		assertlite.Panics(t, func() {
			_ = newBucketWheel[item, int64](cmpInt64(), 0, 10, 0)
		})
		assertlite.Panics(t, func() {
			_ = newBucketWheel[item, int64](cmpInt64(), 0, 10, -4)
		})
	})
}

// TestBucketWheel_PushLen
//
// Tracks the wheel's element count across pushes, pops and clears.
//
// We should assert:
//  1. Push increments Len
//  2. Pop decrements Len
//  3. PopBucket subtracts the number of drained elements
//  4. Clear resets Len to zero
func TestBucketWheel_PushLen(t *testing.T) {
	t.Run("push increments Len", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(0, "a"))
		w.Push(it(5, "b"))
		assertlite.Equals(t, 2, w.Len())
	})

	t.Run("Pop decrements Len", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(0, "a"))
		_, ok := w.Pop()
		assertlite.True(t, ok)
		assertlite.Equals(t, 0, w.Len(), "expected a len of '0' but got '%d'", w.Len())
	})

	t.Run("PopBucket subtracts drained count", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(0, "a"))
		w.Push(it(1, "b"))
		w.Push(it(2, "c"))

		assertlite.Equals(t, 3, w.Len())
		got, ok := w.PopBucket()
		assertlite.True(t, ok)
		assertlite.EqualSlices(t, []item{
			it(0, "a"),
			it(1, "b"),
			it(2, "c"),
		}, got)
		assertlite.Equals(t, 3, len(got))
		assertlite.Equals(t, 0, w.Len())
	})

	t.Run("Clear resets Len", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(0, "a"))
		w.Push(it(5, "b"))

		w.Clear()
		assertlite.Equals(t, 0, w.Len())
	})
}

// TestBucketWheel_Placement
//
// Each push lands in the bucket identified by (key - origin) / span,
// modulo the wheel size. This is the core scheduling math.
//
// We should assert:
//  1. Items in the same span share a bucket
//  2. Items in different spans are separated
//  3. The origin is bucket 0
//  4. A key exactly on a span boundary starts a new bucket
func TestBucketWheel_Placement(t *testing.T) {
	t.Run("same span shares a bucket", func(t *testing.T) {
		// span=10, buckets=8: keys 0..9 -> bucket 0
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(0, "a"))
		w.Push(it(9, "b"))

		got, ok := w.PopBucket()
		assertlite.True(t, ok)
		assertlite.EqualSlices(t, []string{"a", "b"}, itemPayloads(got))
	})

	t.Run("different spans are separated", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(0, "b0"))
		w.Push(it(10, "b1"))

		first, ok := w.PopBucket()
		assertlite.True(t, ok)
		assertlite.EqualSlices(t, []string{"b0"}, itemPayloads(first))

		second, ok := w.PopBucket()
		assertlite.True(t, ok)
		assertlite.EqualSlices(t, []string{"b1"}, itemPayloads(second))
	})

	t.Run("boundary key starts a new bucket", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(9, "end-of-b0"))
		w.Push(it(10, "start-of-b1"))

		first, _ := w.PopBucket()
		assertlite.EqualSlices(t, []string{"end-of-b0"}, itemPayloads(first))

		second, _ := w.PopBucket()
		assertlite.EqualSlices(t, []string{"start-of-b1"}, itemPayloads(second))
	})

	t.Run("origin is bucket zero", func(t *testing.T) {
		// origin=100, span=10: key 100 -> bucket 0
		w := newBucketWheel[item, int64](cmpInt64(), 100, 10, 8)
		w.Push(it(100, "at-origin"))

		got, ok := w.PopBucket()
		assertlite.True(t, ok)
		assertlite.EqualSlices(t, []string{"at-origin"}, itemPayloads(got))
	})
}

// TestBucketWheel_Ordering
//
// Items within a bucket are emitted in key order regardless of push order.
//
// We should assert:
//  1. Push order does not affect emission order
//  2. Interleaved pushes still sort
//  3. Ordering is local to a bucket, not global
func TestBucketWheel_Ordering(t *testing.T) {
	t.Run("push order does not affect emission order", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 100, 8)

		w.Push(it(30, "c"))
		w.Push(it(10, "a"))
		w.Push(it(20, "b"))

		got, _ := w.PopBucket()
		assertlite.EqualSlices(t, []string{"a", "b", "c"}, itemPayloads(got))
	})

	t.Run("ordering is local to a bucket", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)

		// bucket 0
		w.Push(it(5, "b0-high"))
		w.Push(it(1, "b0-low"))
		// bucket 1
		w.Push(it(15, "b1-high"))
		w.Push(it(11, "b1-low"))

		b0, _ := w.PopBucket()
		assertlite.EqualSlices(t, []string{"b0-low", "b0-high"}, itemPayloads(b0))

		b1, _ := w.PopBucket()
		assertlite.EqualSlices(t, []string{"b1-low", "b1-high"}, itemPayloads(b1))
	})
}

// TestBucketWheel_Peek
//
// Peek inspects the current bucket without consuming it.
//
// We should assert:
//  1. Peek returns the smallest item in the current bucket
//  2. Peek does not change Len
//  3. Peek on an empty wheel returns ok=false
//  4. PeekBucket returns all due items without consuming
//  5. PeekRange visits due items without consuming
func TestBucketWheel_Peek(t *testing.T) {
	t.Run("Peek returns the smallest item", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(5, "b"))
		w.Push(it(1, "a"))

		v, ok := w.Peek()
		assertlite.True(t, ok)
		assertlite.Equals(t, "a", v.payload)
	})

	t.Run("Peek does not change Len", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(1, "a"))
		w.Push(it(2, "b"))

		_, _ = w.Peek()
		assertlite.Equals(t, 2, w.Len())
	})

	t.Run("Peek on empty wheel returns ok=false", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		_, ok := w.Peek()
		assertlite.False(t, ok)
	})

	t.Run("PeekBucket does not consume", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(1, "a"))
		w.Push(it(2, "b"))

		got, ok := w.PeekBucket()
		assertlite.True(t, ok)
		assertlite.EqualSlices(t, []string{"a", "b"}, itemPayloads(got))
		assertlite.Equals(t, 2, w.Len())
	})

	t.Run("PeekRange does not consume", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(1, "a"))
		w.Push(it(2, "b"))

		var got []string
		w.PeekRange(func(i item) bool {
			got = append(got, i.payload)
			return true
		})
		assertlite.EqualSlices(t, []string{"a", "b"}, got)
		assertlite.Equals(t, 2, w.Len())
	})
}

// ---------------------------------------------------------------------
// BucketWheel_PopBucket
// ---------------------------------------------------------------------

// TestBucketWheel_PopBucket
//
// PopBucket drains the current bucket and advances the wheel by one.
//
// We should assert:
//  1. Drains all items in the current bucket, in key order
//  2. Advances curr, so a second call drains the next bucket
//  3. Returns ok=false on an empty bucket
//  4. Does not cross bucket boundaries in a single call
func TestBucketWheel_PopBucket(t *testing.T) {
	t.Run("drains current bucket in order", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(2, "b"))
		w.Push(it(1, "a"))

		got, ok := w.PopBucket()
		assertlite.True(t, ok)
		assertlite.EqualSlices(t, []string{"a", "b"}, itemPayloads(got))
	})

	t.Run("advances to the next bucket", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(0, "b0"))
		w.Push(it(10, "b1"))

		first, _ := w.PopBucket()
		assertlite.EqualSlices(t, []string{"b0"}, itemPayloads(first))

		second, _ := w.PopBucket()
		assertlite.EqualSlices(t, []string{"b1"}, itemPayloads(second))
	})

	t.Run("empty bucket returns ok=false but still advances", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(10, "b1")) // bucket 1

		// bucket 0 is empty
		_, ok := w.PopBucket()
		assertlite.False(t, ok)

		// wheel advanced to bucket 1
		got, ok := w.PopBucket()
		assertlite.True(t, ok)
		assertlite.EqualSlices(t, []string{"b1"}, itemPayloads(got))
	})

	t.Run("does not cross bucket boundary", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(0, "b0"))
		w.Push(it(10, "b1"))

		got, _ := w.PopBucket()
		assertlite.EqualSlices(t, []string{"b0"}, itemPayloads(got))
		assertlite.Equals(t, 1, w.Len()) // b1 still in the wheel
	})
}

// ---------------------------------------------------------------------
// BucketWheel_PopRange
// ---------------------------------------------------------------------

// TestBucketWheel_PopRange
//
// PopRange drains the current bucket through a predicate and only
// advances the wheel if the bucket was fully drained. Stopping early
// leaves the remainder in place for the next call.
//
// We should assert:
//  1. Draining fully advances the wheel and returns true
//  2. Stopping early returns false and does not advance
//  3. Stopped items are recoverable on the next call
//  4. An empty bucket advances and returns true
func TestBucketWheel_PopRange(t *testing.T) {
	t.Run("full drain returns true and advances", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(1, "a"))
		w.Push(it(2, "b"))

		var got []string
		ok := w.PopRange(func(i item) bool {
			got = append(got, i.payload)
			return true
		})

		assertlite.True(t, ok)
		assertlite.EqualSlices(t, []string{"a", "b"}, got)
		assertlite.Equals(t, 0, w.Len())
	})

	t.Run("early stop returns false and does not advance", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(1, "a"))
		w.Push(it(2, "b"))
		w.Push(it(3, "c"))

		ok := w.PopRange(func(i item) bool {
			return i.payload == "a" // stop after the first
		})
		assertlite.False(t, ok)
	})

	t.Run("stopped items are recoverable", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(1, "a"))
		w.Push(it(2, "b"))
		w.Push(it(3, "c"))

		// Drain only "a".
		w.PopRange(func(i item) bool { return i.payload != "a" })

		// The rest of the bucket must still be there.
		var rest []string
		w.PopRange(func(i item) bool {
			rest = append(rest, i.payload)
			return true
		})
		assertlite.EqualSlices(t, []string{"b", "c"}, rest)
	})

	t.Run("empty bucket advances and returns true", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(10, "b1"))

		ok := w.PopRange(func(item) bool { return true })
		assertlite.True(t, ok)

		var got []string
		w.PopRange(func(i item) bool {
			got = append(got, i.payload)
			return true
		})
		assertlite.EqualSlices(t, []string{"b1"}, got)
	})
}

// ---------------------------------------------------------------------
// BucketWheel_Lap
// ---------------------------------------------------------------------

// TestBucketWheel_Lap
//
// When the key space exceeds buckets*span, pushes land on later laps.
// popIfDue only releases an item when its lap matches the wheel's lap.
//
// We should assert:
//  1. A far-future item is not drained on the current lap
//  2. It is drained once the wheel reaches its lap
//  3. It is not lost in between
func TestBucketWheel_Lap(t *testing.T) {
	t.Run("far item is not drained on the current lap", func(t *testing.T) {
		// span=10, buckets=8 -> 80 units per lap
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(0, "lap0"))  // bucket 0, lap 0
		w.Push(it(90, "lap1")) // 90/10=9 -> bucket 1, lap 1

		got, _ := w.PopBucket()
		assertlite.EqualSlices(t, []string{"lap0"}, itemPayloads(got))
	})

	t.Run("far item is drained when its lap is reached", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(90, "lap1")) // bucket 1, lap 1

		// Lap 0: bucket 0 is empty, bucket 1 holds "lap1" but wrong lap.
		_, _ = w.PopBucket() // bucket 0 -> advance
		got, ok := w.PopBucket()
		// The wheel is now at lap 0, bucket 1, but "lap1" belongs to lap 1,
		// so it must not be released yet.
		assertlite.False(t, ok)
		assertlite.Equals(t, 0, len(got))
		assertlite.Equals(t, 1, w.Len()) // still held

		// Advance the rest of lap 0.
		for i := 0; i < 6; i++ {
			_, _ = w.PopBucket()
		}
		// One more advance brings us to bucket 0 of lap 1.
		_, _ = w.PopBucket()

		// Now bucket 1 of lap 1 should release the item.
		got, ok = w.PopBucket()
		assertlite.True(t, ok)
		assertlite.EqualSlices(t, []string{"lap1"}, itemPayloads(got))
	})

	t.Run("multiple laps keep items distinct", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(5, "lap0"))   // bucket 0, lap 0
		w.Push(it(85, "lap1"))  // 85/10=8 -> bucket 0, lap 1
		w.Push(it(165, "lap2")) // 165/10=16 -> bucket 0, lap 2

		first, ok := w.PopBucket()
		assertlite.True(t, ok)
		assertlite.EqualSlices(t, []string{"lap0"}, itemPayloads(first))
		assertlite.Equals(t, 2, w.Len())
	})
}

// ---------------------------------------------------------------------
// BucketWheel_Wrap
// ---------------------------------------------------------------------

// TestBucketWheel_Wrap
//
// A full lap returns curr to bucket 0 and increments lap.
//
// We should assert:
//  1. Advancing through all buckets wraps to bucket 0
//  2. After a full wrap, items of the next lap become due
//  3. The wheel never indexes out of range
func TestBucketWheel_Wrap(t *testing.T) {
	t.Run("advancing through all buckets wraps cleanly", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 4)

		for i := 0; i < 4; i++ {
			_, _ = w.PopBucket()
		}
		// Wheel should be back at bucket 0. Pushing something in bucket 0
		// of the next lap and draining should work without panic.
		w.Push(it(40, "next-lap"))
		got, _ := w.PopBucket()
		assertlite.EqualSlices(t, []string{"next-lap"}, itemPayloads(got))
	})
}

// ---------------------------------------------------------------------
// BucketWheel_Clear
// ---------------------------------------------------------------------

// TestBucketWheel_Clear
//
// Clear resets the wheel.
//
// We should assert:
//  1. Len becomes zero
//  2. Subsequent pop does not return a previously-pushed item
//  3. The wheel is still usable after Clear (push works)
func TestBucketWheel_Clear(t *testing.T) {
	t.Run("Len becomes zero", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(0, "a"))
		w.Push(it(5, "b"))

		w.Clear()
		assertlite.Equals(t, 0, w.Len())
	})

	t.Run("cleared items are not returned", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(0, "a"))

		w.Clear()
		got, ok := w.PopBucket()
		assertlite.False(t, ok)
		assertlite.Equals(t, 0, len(got))
	})

	t.Run("wheel is usable after Clear", func(t *testing.T) {
		w := newBucketWheel[item, int64](cmpInt64(), 0, 10, 8)
		w.Push(it(0, "old"))
		w.Clear()
		w.Push(it(0, "new"))

		got, ok := w.PopBucket()
		assertlite.True(t, ok)
		assertlite.EqualSlices(t, []string{"new"}, itemPayloads(got))
	})
}
