// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package queues

import (
	"testing"

	"github.com/andrerrcosta2/gtools/core/data/comparator"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
)

// ---------------------------------------------------------------------
// fixtures
// ---------------------------------------------------------------------

// task is a domain type that does NOT implement constraints.Ordered.
// It is the intended payload for Priority(cmp), where the caller must
// supply the ordering externally.
type task struct {
	prio    int
	payload string
}

func tk(prio int, payload string) task {
	return task{prio: prio, payload: payload}
}

// taskComparer orders tasks by priority only. This is the comparer the
// caller passes to Priority[task].
//
// We deliberately make Compare return values OTHER than -1/0/1 (-7 and
// +7) so that any implementation that tests `== -1` instead of `< 0`
// fails loudly here. See TestPriority_ComparatorContract.
func taskComparer() comparator.Comparer[task] {
	return comparator.Build[task, any]().
		Compare(func(a, b task) int {
			switch {
			case a.prio < b.prio:
				return -7
			case a.prio > b.prio:
				return 7
			default:
				return 0
			}
		}).
		Build()
}

// rank is a domain type that DOES implement constraints.Ordered.
// It is the intended payload for PriorityOf[rank].
type rank int

func (r rank) Less(other rank) bool { return r < other }

func payloadsT(tasks []task) []string {
	out := make([]string, 0, len(tasks))
	for _, t := range tasks {
		out = append(out, t.payload)
	}
	return out
}

// TestPriority_Construction (cmp-driven constructor)
//
// Builds the queue with the caller-supplied comparator.
//
// We should assert:
//  1. Priority(cmp) returns a non-nil, empty queue
//  2. The queue is immediately usable (Push does not panic)
//  3. The stored values are the caller's type, not a wrapped one
func TestPriority_Construction(t *testing.T) {
	t.Run("returns an empty queue", func(t *testing.T) {
		q := Priority[task](taskComparer())
		assertlite.NotNil(t, q)
		assertlite.Equals(t, 0, q.Len())
	})

	t.Run("accepts pushes of the caller's type", func(t *testing.T) {
		q := Priority[task](taskComparer())
		q.Push(tk(1, "a"))
		assertlite.Equals(t, 1, q.Len())
	})
}

// TestPriority_ComparatorContract (the == -1 vs < 0 bug)
//
// The queue must accept any comparator that follows the conventional
// "<0, 0, >0" contract. It must NOT assume the specific value -1 for
// "less than".
//
// We should assert:
//  1. A comparator returning -7 for "less" still produces ascending order
//  2. A comparator returning +7 for "greater" still produces ascending order
//  3. Insertion order is irrelevant
func TestPriority_ComparatorContract(t *testing.T) {
	t.Run("non-trivial negative values still order correctly", func(t *testing.T) {
		q := Priority[task](taskComparer())

		q.Push(tk(3, "c"))
		q.Push(tk(1, "a"))
		q.Push(tk(2, "b"))

		var got []string
		for q.Len() > 0 {
			v, _ := q.Pop()
			got = append(got, v.payload)
		}
		assertlite.EqualSlices(t, []string{"a", "b", "c"}, got)
	})
}

// TestPriorityOf_Construction (Less-driven constructor)
//
// Builds the queue with the synthesized Ordered comparator.
//
// We should assert:
//  1. PriorityOf[T]() returns a non-nil, empty queue
//  2. The queue is immediately usable
//  3. The synthesized comparator orders by Less
func TestPriorityOf_Construction(t *testing.T) {
	t.Run("returns an empty queue", func(t *testing.T) {
		q := PriorityOf[rank]()
		assertlite.NotNil(t, q)
		assertlite.Equals(t, 0, q.Len())
	})

	t.Run("accepts pushes", func(t *testing.T) {
		q := PriorityOf[rank]()
		q.Push(1)
		assertlite.Equals(t, 1, q.Len())
	})
}

// TestPriorityOf_Ordering
//
// The synthesized Ordered comparator yields ascending order by Less.
//
// We should assert:
//  1. Values come out in ascending order
//  2. Push order does not matter
//  3. Duplicates are preserved
//  4. Negative and zero values order correctly
func TestPriorityOf_Ordering(t *testing.T) {
	t.Run("ascending order", func(t *testing.T) {
		q := PriorityOf[rank]()
		q.Push(3)
		q.Push(1)
		q.Push(2)

		var got []rank
		for q.Len() > 0 {
			v, _ := q.Pop()
			got = append(got, v)
		}
		assertlite.EqualSlices(t, []rank{1, 2, 3}, got)
	})

	t.Run("push order does not matter", func(t *testing.T) {
		asc := PriorityOf[rank]()
		for _, v := range []rank{1, 2, 3, 4, 5} {
			asc.Push(v)
		}
		desc := PriorityOf[rank]()
		for _, v := range []rank{5, 4, 3, 2, 1} {
			desc.Push(v)
		}

		var a, d []rank
		for asc.Len() > 0 {
			v, _ := asc.Pop()
			a = append(a, v)
		}
		for desc.Len() > 0 {
			v, _ := desc.Pop()
			d = append(d, v)
		}
		assertlite.EqualSlices(t, a, d)
	})

	t.Run("duplicates are preserved", func(t *testing.T) {
		q := PriorityOf[rank]()
		q.Push(2)
		q.Push(1)
		q.Push(2)
		q.Push(1)

		var got []rank
		for q.Len() > 0 {
			v, _ := q.Pop()
			got = append(got, v)
		}
		assertlite.EqualSlices(t, []rank{1, 1, 2, 2}, got)
	})

	t.Run("negative and zero order correctly", func(t *testing.T) {
		q := PriorityOf[rank]()
		q.Push(0)
		q.Push(-1)
		q.Push(1)
		q.Push(-2)

		var got []rank
		for q.Len() > 0 {
			v, _ := q.Pop()
			got = append(got, v)
		}
		assertlite.EqualSlices(t, []rank{-2, -1, 0, 1}, got)
	})
}

// TestPriority_Peek
//
// Peek returns the minimum without removing it, for both constructors.
//
// We should assert:
//  1. Peek returns the smallest element
//  2. Peek does not change Len
//  3. Peek on an empty queue returns (zero, false)
//  4. Repeated Peek returns the same value
func TestPriority_Peek(t *testing.T) {
	t.Run("PriorityOf: returns the smallest", func(t *testing.T) {
		q := PriorityOf[rank]()
		q.Push(3)
		q.Push(1)
		q.Push(2)

		v, ok := q.Peek()
		assertlite.True(t, ok)
		assertlite.Equals(t, rank(1), v)
	})

	t.Run("Priority: returns the smallest by cmp", func(t *testing.T) {
		q := Priority[task](taskComparer())
		q.Push(tk(3, "c"))
		q.Push(tk(1, "a"))

		v, ok := q.Peek()
		assertlite.True(t, ok)
		assertlite.Equals(t, "a", v.payload)
	})

	t.Run("Peek does not change Len", func(t *testing.T) {
		q := PriorityOf[rank]()
		q.Push(3)
		q.Push(1)

		_, _ = q.Peek()
		assertlite.Equals(t, 2, q.Len())
	})

	t.Run("empty queue returns ok=false", func(t *testing.T) {
		q := PriorityOf[rank]()
		v, ok := q.Peek()
		assertlite.False(t, ok)
		assertlite.Equals(t, rank(0), v)
	})

	t.Run("repeated Peek returns the same value", func(t *testing.T) {
		q := PriorityOf[rank]()
		q.Push(5)
		q.Push(2)

		a, _ := q.Peek()
		b, _ := q.Peek()
		assertlite.Equals(t, a, b)
		assertlite.Equals(t, rank(2), a)
	})
}

// TestPriority_Pop
//
// Pop removes and returns the minimum element.
//
// We should assert:
//  1. Pop returns the minimum
//  2. Pop decrements Len
//  3. Popping the last element still returns its value
//  4. Pop on an empty queue returns (zero, false) without panicking
func TestPriority_Pop(t *testing.T) {
	t.Run("returns the minimum", func(t *testing.T) {
		q := PriorityOf[rank]()
		q.Push(3)
		q.Push(1)
		q.Push(2)

		v, _ := q.Pop()
		assertlite.Equals(t, rank(1), v)
	})

	t.Run("decrements Len", func(t *testing.T) {
		q := PriorityOf[rank]()
		q.Push(3)
		q.Push(1)

		_, _ = q.Pop()
		assertlite.Equals(t, 1, q.Len())
	})

	t.Run("popping the last element returns its value", func(t *testing.T) {
		q := PriorityOf[rank]()
		q.Push(42)

		v, _ := q.Pop()
		assertlite.Equals(t, rank(42), v)
		assertlite.Equals(t, 0, q.Len())
	})

	t.Run("empty queue returns ok=false without panicking", func(t *testing.T) {
		q := PriorityOf[rank]()

		var (
			v  rank
			ok bool
		)
		assertlite.NoPanics(t, func() {
			v, ok = q.Pop()
		})
		assertlite.False(t, ok)
		assertlite.Equals(t, rank(0), v)
	})
}

// TestPriority_Clear
//
// Clear empties the queue without invalidating it.
//
// We should assert:
//  1. Len becomes zero
//  2. Previously pushed values are gone
//  3. The queue is reusable after Clear
func TestPriority_Clear(t *testing.T) {
	t.Run("Len becomes zero", func(t *testing.T) {
		q := PriorityOf[rank]()
		q.Push(1)
		q.Push(2)
		q.Clear()
		assertlite.Equals(t, 0, q.Len())
	})

	t.Run("values are gone", func(t *testing.T) {
		q := PriorityOf[rank]()
		q.Push(1)
		q.Clear()

		_, ok := q.Peek()
		assertlite.False(t, ok)
	})

	t.Run("queue is reusable after Clear", func(t *testing.T) {
		q := PriorityOf[rank]()
		q.Push(1)
		q.Clear()
		q.Push(2)

		v, ok := q.Peek()
		assertlite.True(t, ok)
		assertlite.Equals(t, rank(2), v)
	})
}

// TestPriority_WheelKeyContract
//
// The bucket wheel depends on exactly this subset of behavior. Locking
// it in here means any regression in the queue fails at this test
// before it can confuse the wheel suite.
//
// We should assert:
//  1. Push -> Peek -> Pop -> Peek yields the next minimum (the wheel's
//     popIfDue pattern)
//  2. Len tracks mixed pushes and pops
//  3. Clear followed by Push keeps the queue usable (the wheel calls
//     Clear on shutdown)
func TestPriority_WheelKeyContract(t *testing.T) {
	t.Run("peek-pop-peek yields the next minimum", func(t *testing.T) {
		q := PriorityOf[rank]()
		q.Push(3)
		q.Push(1)
		q.Push(2)

		first, _ := q.Peek()
		popped, _ := q.Pop()
		assertlite.Equals(t, first, popped)

		next, _ := q.Peek()
		assertlite.Equals(t, rank(2), next)
	})

	t.Run("Len tracks mixed pushes and pops", func(t *testing.T) {
		q := PriorityOf[rank]()
		q.Push(1)
		q.Push(2)
		q.Push(3)
		_, _ = q.Pop()
		_, _ = q.Pop()
		q.Push(0)
		assertlite.Equals(t, 2, q.Len())
	})

	t.Run("Clear then Push keeps the queue usable", func(t *testing.T) {
		q := PriorityOf[rank]()
		q.Push(1)
		q.Clear()
		q.Push(2)

		v, ok := q.Peek()
		assertlite.True(t, ok)
		assertlite.Equals(t, rank(2), v)
	})
}
