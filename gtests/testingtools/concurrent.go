// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testingtools

import (
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/gtests"
	"sync"
)

type toolsLite struct {
	gtests.HelperTesting
	mtx   sync.RWMutex
	calls map[string]*call
	flags map[string]bool
	cons  map[string]any
}

// AssertCalls asserts that the number of calls for the given ids is compare to the given calls.
// The ids are used to identify the call in the tools.
// If the id is not present in the tools, it returns an error.
func (t *toolsLite) AssertCalls(calls int, ids ...string) {
	t.Helper()
	t.mtx.RLock()
	defer t.mtx.RUnlock()

	if t.calls == nil && calls != 0 {
		for _, id := range ids {
			t.Errorf("expected %d calls to '%s', got none\n", calls, id)
		}
	}

	for _, id := range ids {
		c, ok := t.calls[id]

		if !ok {
			if calls != 0 {
				t.Errorf("expected %d calls to '%s', got none\n", calls, id)
			}
			continue
		}

		if c.value != calls {
			t.Errorf("expected %d calls to '%s', got %d\n", calls, id, c.value)
		}
	}
}

func (t *toolsLite) AssertCallsTo(name string, calls int, errorMessage string, args ...any) {
	t.Helper()
	t.mtx.RLock()
	defer t.mtx.RUnlock()

	c, ok := t.calls[name]

	if !ok {
		calls = 0
	}

	if c.value != calls {
		t.Errorf(errorMessage, args...)
	}
}

// AssertRegisteredCalls asserts the size of registers on the calls map is compare to the given size.
func (t *toolsLite) AssertRegisteredCalls(size int, errorMessage string, args ...any) {
	t.Helper()
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	if len(t.calls) != size {
		t.Errorf(errorMessage, args...)
	}
}

// AssertRegisteredCallsFunc asserts that the number of calls that satisfy the predicate is compare to the expected count
func (t *toolsLite) AssertRegisteredCallsBy(f functions.BiPredicate[string, int], expectedCount int, errorMessage string, args ...any) {
	t.Helper()
	t.mtx.RLock()
	defer t.mtx.RUnlock()

	var calls int

	for id, c := range t.calls {
		if !f(id, c.value) {
			calls++
		}
	}

	if calls != expectedCount {
		t.Errorf(errorMessage, args...)
	}
}

// AssertConst asserts that the value of the constant with the given ids is compare to the given value
func (t *toolsLite) AssertConst(value any, ids ...string) {
	t.Helper()
	t.mtx.RLock()
	defer t.mtx.RUnlock()

	for _, id := range ids {
		if t.cons[id] != value {
			t.Errorf("expected constant '%s' to be %v, got %v\n", id, value, t.cons[id])
		}
	}
}

func (t *toolsLite) AssertConstTo(name string, value any, errorMessage string, args ...any) {
	t.Helper()
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	v, ok := t.cons[name]
	if !ok {
		t.Errorf("constant '%s' not registered", name)
	} else if v != value {
		t.Errorf(errorMessage, args...)
	}
}

func (t *toolsLite) AssertRegisteredConst(size int, errorMessage string, args ...any) {
	t.Helper()
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	if len(t.cons) != size {
		t.Errorf(errorMessage, args...)
	}
}

// AssertRegisteredConstFunc asserts that the number of constants that satisfy the predicate is compare to the expected count
func (t *toolsLite) AssertRegisteredConstBy(f functions.BiPredicate[string, any], expectedCount int, errorMessage string, args ...any) {
	t.Helper()
	t.mtx.RLock()
	defer t.mtx.RUnlock()

	count := 0
	for k, v := range t.cons {
		if f(k, v) {
			count++
		}
	}

	if count != expectedCount {
		t.Errorf(errorMessage, args...)
	}
}

func (t *toolsLite) AssertFlag(flag bool, ids ...string) {
	t.Helper()
	t.mtx.Lock()
	defer t.mtx.Unlock()

	for _, id := range ids {
		if t.flags[id] != flag {
			t.Errorf("expected flag '%s' to be %t, got %t\n", id, flag, t.flags[id])
		}
	}
}

func (t *toolsLite) AssertFlagTo(name string, flag bool, errorMessage string, args ...any) {
	t.Helper()
	t.mtx.Lock()
	defer t.mtx.Unlock()

	if t.flags == nil {
		t.Errorf("no flags were registered")
	}

	if t.flags[name] != flag {
		t.Errorf(errorMessage, args...)
	}
}

func (t *toolsLite) AssertRegisteredFlags(size int, errorMessage string, args ...any) {
	t.Helper()
	t.mtx.Lock()
	defer t.mtx.Unlock()
	if len(t.flags) != size {
		t.Errorf(errorMessage, args...)
	}
}

func (t *toolsLite) AssertRegisteredFlagsBy(f functions.BiPredicate[string, bool], expectedCount int, errorMessage string, args ...any) {
	t.Helper()
	t.mtx.RLock()
	defer t.mtx.RUnlock()

	// Count the number of flags satisfying the f
	count := 0
	for key, value := range t.flags {
		if f(key, value) {
			count++
		}
	}

	// Assert the count matches the expected value
	if count != expectedCount {
		t.Errorf(errorMessage, args...)
	}
}

// CallsSize returns the number of registered calls.
// This function is thread-safe and can be used concurrently.
func (t *toolsLite) CallsSize() int {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	return len(t.calls)
}

// CallsTo returns the number of calls for the given name.
// The name is used to identify the call in the tools.
// If the name is not present in the tools, it returns 0.
func (t *toolsLite) CallsTo(name string) int {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	if t.calls == nil {
		return 0
	}
	if c, ok := t.calls[name]; ok {
		return c.load()
	}
	return 0
}

func (t *toolsLite) Clear() {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	t.calls = make(map[string]*call)
	t.flags = make(map[string]bool)
	t.cons = make(map[string]any)
}

func (t *toolsLite) Condition(condition bool, notTrueMessage string, args ...any) {
	t.Helper()
	t.mtx.Lock()
	defer t.mtx.Unlock()
	if !condition {
		t.HelperTesting.Errorf(notTrueMessage, args...)
	}
}

// Const returns the constant for the given name.
// The name is used to identify the constant in the tools.
// If the name is not present in the tools, it returns nil.
func (t *toolsLite) Const(name string) any {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	return t.cons[name]
}

// ConstSize returns the number of registered constants.
// This function is thread-safe and can be used concurrently.
func (t *toolsLite) ConstSize() int {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	return len(t.cons)
}

// Flag registers a flag for the given ids.
// The ids are used to identify the flag in the tools.
// If the id is not present in the tools, it is created.
func (t *toolsLite) Flag(value bool, ids ...string) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	if t.flags == nil {
		t.flags = make(map[string]bool)
	}

	for _, id := range ids {
		t.flags[id] = value
	}
}

// FlagsSize returns the number of registered flags.
// This function is thread-safe and can be used concurrently.
func (t *toolsLite) FlagsSize() int {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	return len(t.flags)
}

// RegisterCalls registers a call count for the given ids.
// The ids are used to identify the call in the tools.
// If the id is not present in the tools, it is created.
// The value is incremented to a given id.
func (t *toolsLite) RegisterCalls(increment int, ids ...string) {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	if t.calls == nil {
		// Create a new map if it doesn't exist
		t.calls = make(map[string]*call)
	}

	// add the call count to the counter for each id
	for _, id := range ids {
		// Increment the counter
		if _, ok := t.calls[id]; !ok {
			// If it doesn't exist, create a new call
			// with the given increment as the initial value
			t.calls[id] = caller(increment, id, nil)
			continue
		}
		// add the increment to the current value
		// of the counter
		t.calls[id].add(increment)
	}
}

func (t *toolsLite) RegisterCallback(callback functions.Runnable, value int, ids ...string) {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	if t.calls == nil {
		// Create a new map if it doesn't exist
		t.calls = make(map[string]*call)
	}

	// add the callback to the map
	for _, id := range ids {
		if _, ok := t.calls[id]; !ok {
			// If it doesn't exist, create a new call
			// with the given increment as the initial value
			t.calls[id] = caller(0, id, callback)
		}
		// add the increment to the current value
		// of the counter
		t.calls[id].setCallback(callback, value)
	}
}

// RegisterConst registers a constant for the given ids.
// the constant cannot be changed
func (t *toolsLite) RegisterConst(cons any, ids ...string) {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	for _, id := range ids {
		if _, ok := t.cons[id]; !ok {
			t.cons[id] = cons
		}
	}
}

var _ gtests.Tools = (*toolsLite)(nil)
