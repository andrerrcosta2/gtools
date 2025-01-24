// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testingtools

import (
	"github.com/andrerrcosta2/gtools/core/gtools/functions"
	"github.com/andrerrcosta2/gtools/gtests"
	"sync"
)

type toolsLite struct {
	gtests.HelperTesting
	mtx   sync.RWMutex
	calls map[string]*call
	flags map[string]bool
}

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

func (t *toolsLite) Condition(condition bool, notTrueMessage string, args ...any) {
	t.Helper()
	t.mtx.Lock()
	defer t.mtx.Unlock()
	if !condition {
		t.HelperTesting.Errorf(notTrueMessage, args...)
	}
}

func (t *toolsLite) AssertCalls(calls int, ids ...string) {
	t.Helper()
	t.mtx.Lock()
	defer t.mtx.Unlock()

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
	t.mtx.Lock()
	defer t.mtx.Unlock()

	c, ok := t.calls[name]

	if !ok {
		calls = 0
	}

	if c.value != calls {
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

var _ gtests.Tools = (*toolsLite)(nil)
