// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package subscr

import (
	"sync"
	"testing"

	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gflux/core/lifecycle/gates"
	"github.com/andrerrcosta2/gtools/gflux/obs/internal/observers"
)

// TestBind tests the Bind subscription wrapper.
//
// We aim to assert:
//  1. Unsub closes the provided gate
//  2. Unsub propagates to the inner subscription
//  3. Unsub is idempotent (inner called only once)
//  4. Closed reflects gate states
//  5. Works correctly when inner is nil (gate-only behavior)
func TestBind(t *testing.T) {
	t.Run("unsub closes gate and propagates to inner", func(t *testing.T) {
		gate := gates.Atomic()
		innerCalled := 0
		inner := Func(func() {
			innerCalled++
		})

		sub := Bind(inner, gate)
		sub.Unsub()

		assertlite.True(t, gate.Closed())
		assertlite.True(t, sub.Closed())
		assertlite.True(t, 1 == innerCalled, "expected inner to be called once but got '%d'", innerCalled)
	})

	t.Run("unsub is idempotent", func(t *testing.T) {
		gate := gates.Atomic()

		innerCalled := 0
		inner := Func(func() {
			innerCalled++
		})

		sub := Bind(inner, gate)

		sub.Unsub()
		sub.Unsub()

		assertlite.True(t, gate.Closed())
		assertlite.True(t, 1 == innerCalled, "expected inner to be called once but got %d", innerCalled)
	})

	t.Run("closed reflects gate states", func(t *testing.T) {
		gate := gates.Atomic()

		inner := Func(func() {})
		sub := Bind(inner, gate)

		assertlite.False(t, sub.Closed())

		gate.Close()

		assertlite.True(t, sub.Closed())
	})

	t.Run("nil inner behaves as gate-only subscription", func(t *testing.T) {
		gate := gates.Atomic()

		sub := Bind(nil, gate)

		assertlite.False(t, sub.Closed())

		sub.Unsub()

		assertlite.True(t, gate.Closed())
		assertlite.True(t, sub.Closed())
	})
}

// TestSlotSub tests the SlotSub subscription implementation.
//
// We aim to assert:
//  1. Unsub closes the associated gate
//  2. Closed reflects gate states correctly
//  3. Unsub is idempotent
//  4. Multiple calls to Unsub do not panic
//  5. Gate external close is reflected in Closed
func TestSlotSub(t *testing.T) {

	t.Run("unsub closes gate", func(t *testing.T) {
		gate := gates.Atomic()
		observer := observers.Noop[int]()

		sub := NewSlotSub(observer, gate)

		assertlite.False(t, gate.Closed())
		assertlite.False(t, sub.Closed())

		sub.Unsub()

		assertlite.True(t, gate.Closed())
		assertlite.True(t, sub.Closed())
	})

	t.Run("closed reflects external gate close", func(t *testing.T) {
		gate := gates.Atomic()
		observer := observers.Noop[int]()

		sub := NewSlotSub(observer, gate)

		gate.Close()

		assertlite.True(t, sub.Closed())
	})

	t.Run("unsub is idempotent", func(t *testing.T) {
		gate := gates.Atomic()
		observer := observers.Noop[int]()

		sub := NewSlotSub(observer, gate)

		sub.Unsub()
		sub.Unsub()

		assertlite.True(t, gate.Closed())
		assertlite.True(t, sub.Closed())
	})

	t.Run("multiple unsub does not panic", func(t *testing.T) {
		gate := gates.Atomic()
		observer := observers.Noop[int]()

		sub := NewSlotSub(observer, gate)

		assertlite.NoPanics(t, func() {
			sub.Unsub()
			sub.Unsub()
		})
	})
}

// TestDiSubscr tests the DiSubscr implementation.
//
// We aim to assert:
//  1. Unsub calls the provided remove function
//  2. Remove function is called only once (idempotent)
//  3. Closed reflects unsubscribe states
//  4. Does not call remove before Unsub
//  5. Nil remove function is handled safely
func TestDiSubscr(t *testing.T) {
	t.Run("unsub calls remove function", func(t *testing.T) {
		called := 0

		remove := UnsubscribableFunc(func(s Subscription) {
			called++
		})

		sub := NewDiSub[any](remove)

		assertlite.False(t, sub.Closed())
		assertlite.True(t, 0 == called)

		sub.Unsub()

		assertlite.True(t, sub.Closed())
		assertlite.True(t, 1 == called)
	})

	t.Run("unsub is idempotent", func(t *testing.T) {
		called := 0
		remove := UnsubscribableFunc(func(s Subscription) {
			called++
		})

		sub := NewDiSub[any](remove)

		sub.Unsub()
		sub.Unsub()
		sub.Unsub()

		assertlite.True(t, sub.Closed())
		assertlite.True(t, 1 == called, "expected unsub to be called 1 times, but got '%d'", called)
	})

	t.Run("does not call remove before unsub", func(t *testing.T) {
		called := 0
		remove := UnsubscribableFunc(func(s Subscription) {
			called++
		})

		sub := NewDiSub[any](remove)

		assertlite.True(t, 0 == called)
		assertlite.False(t, sub.Closed())
	})

	t.Run("nil remove is safe", func(t *testing.T) {
		sub := NewDiSub[any](nil)

		assertlite.False(t, sub.Closed())

		sub.Unsub()
		sub.Unsub()

		assertlite.True(t, sub.Closed())
	})
}

// TestNoop tests the Noop subscription implementation.
//
// We aim to assert:
//  1. Unsub does nothing
//  2. Closed always returns false
//  3. Unsub is idempotent
//  4. Multiple Unsub calls do not panic
func TestNoop(t *testing.T) {
	t.Run("unsub does nothing", func(t *testing.T) {
		sub := NewNoop()
		prev := sub.Closed()

		assertlite.False(t, sub.Closed())
		sub.Unsub()
		late := sub.Closed()

		assertlite.False(t, sub.Closed(), "expected always return false for closed")
		assertlite.True(t, prev == late, "expected unsub to be the same but got previous '%t'"+
			" and latest '%t'", prev, late)
	})

	t.Run("unsub is idempotent", func(t *testing.T) {
		sub := NewNoop()
		prev := sub.Closed()

		sub.Unsub()
		sub.Unsub()
		sub.Unsub()
		late := sub.Closed()

		assertlite.False(t, sub.Closed(), "expected always return false for closed")
		assertlite.True(t, prev == late, "expected unsub to be the same but got previous '%t'"+
			" and latest '%t'", prev, late)
	})
}

// TestSwitching tests the Switching subscription implementation.
//
// We aim to assert:
//  1. Switching replaces the active subscription
//  2. Previous subscription is unsubscribed when switching
//  3. Unsub closes the current active subscription
//  4. Unsub is idempotent
//  5. Switching after Unsub immediately unsubscribes the new subscription
//  6. Concurrent switching keeps only the last active subscription
func TestSwitching(t *testing.T) {
	t.Run("switch replaces active subscription", func(t *testing.T) {
		calledA := 0
		calledB := 0

		a := Func(func() { calledA++ })
		b := Func(func() { calledB++ })

		sw := NewSwitching()

		sw.Switch(a)
		sw.Switch(b)

		// A should have been unsubscribed when B replaced it
		assertlite.True(t, 1 == calledA)
		assertlite.True(t, 0 == calledB)

		sw.Unsub()

		assertlite.True(t, 1 == calledB)
		assertlite.True(t, sw.Closed())
	})

	t.Run("unsub closes current active subscription", func(t *testing.T) {
		called := 0
		inner := Func(func() { called++ })

		sw := NewSwitching()
		sw.Switch(inner)

		sw.Unsub()

		assertlite.True(t, 1 == called)
		assertlite.True(t, sw.Closed())
	})

	t.Run("unsub is idempotent", func(t *testing.T) {
		called := 0
		inner := Func(func() { called++ })

		sw := NewSwitching()
		sw.Switch(inner)

		sw.Unsub()
		sw.Unsub()
		sw.Unsub()

		assertlite.True(t, 1 == called)
		assertlite.True(t, sw.Closed())
	})

	t.Run("switch after unsub immediately unsubscribes new sub", func(t *testing.T) {
		called := 0
		inner := Func(func() { called++ })

		sw := NewSwitching()
		sw.Unsub()

		sw.Switch(inner)

		// Since switching is already closed, inner must be unsubscribed immediately
		assertlite.True(t, 1 == called)
		assertlite.True(t, sw.Closed())
	})

	t.Run("concurrent switching keeps only last active", func(t *testing.T) {
		sw := NewSwitching()

		const n = 50
		calls := make([]int, n)

		var wg sync.WaitGroup
		wg.Add(n)

		for i := 0; i < n; i++ {
			go func(call int) {
				defer wg.Done()
				sw.Switch(Func(func() {
					calls[call]++
				}))
			}(i)
		}

		wg.Wait()
		sw.Unsub()
		total := 0

		for _, c := range calls {
			total += c
		}

		assertlite.True(t, total == n, "expected just a single call after unsub, but got '%d'", total)
		assertlite.True(t, sw.Closed())
	})
}

// TestSet tests the Set composite subscription implementation.
//
// We aim to assert:
//  1. Unsub unsubscribes all initial subscriptions
//  2. Add registers additional subscriptions
//  3. Unsub unsubscribes both initial and added subscriptions
//  4. Unsub is idempotent (subscriptions unsubscribed only once)
//  5. Add after Unsub immediately unsubscribes the subscription
//  6. Add(nil) is ignored safely
//  7. Closed reflects lifecycle states
func TestSet(t *testing.T) {
	t.Run("unsub unsubscribes all initial subscriptions", func(t *testing.T) {
		calls := 0

		s := SetOf(
			Func(func() { calls++ }),
			Func(func() { calls++ }),
			Func(func() { calls++ }),
		)

		s.Unsub()

		assertlite.True(t, calls == 3)
		assertlite.True(t, s.Closed())
	})

	t.Run("add registers additional subscriptions", func(t *testing.T) {
		calls := 0

		s := SetOf()
		s.Add(Func(func() { calls++ }))
		s.Add(Func(func() { calls++ }))

		s.Unsub()

		assertlite.True(t, calls == 2)
		assertlite.True(t, s.Closed())
	})

	t.Run("unsub unsubscribes initial and added subscriptions", func(t *testing.T) {
		calls := 0

		s := SetOf(Func(func() { calls++ }))
		s.Add(Func(func() { calls++ }))

		s.Unsub()

		assertlite.True(t, calls == 2)
		assertlite.True(t, s.Closed())
	})

	t.Run("unsub is idempotent", func(t *testing.T) {
		calls := 0

		s := SetOf(
			Func(func() { calls++ }),
			Func(func() { calls++ }),
		)

		s.Unsub()
		s.Unsub()
		s.Unsub()

		assertlite.True(t, calls == 2)
		assertlite.True(t, s.Closed())
	})

	t.Run("add after unsub immediately unsubscribes subscription", func(t *testing.T) {
		calls := 0

		s := SetOf(Func(func() { calls++ }))
		s.Unsub()

		s.Add(Func(func() { calls++ }))
		s.Add(Func(func() { calls++ }))

		// 1 from initial + 2 from late additions
		assertlite.True(t, calls == 3)
		assertlite.True(t, s.Closed())
	})

	t.Run("add nil is ignored safely", func(t *testing.T) {
		calls := 0

		s := SetOf(Func(func() { calls++ }))
		s.Add(nil)

		s.Unsub()

		assertlite.True(t, calls == 1)
		assertlite.True(t, s.Closed())
	})

	t.Run("closed reflects lifecycle states", func(t *testing.T) {
		s := SetOf()

		assertlite.True(t, !s.Closed())

		s.Unsub()

		assertlite.True(t, s.Closed())
	})
}
