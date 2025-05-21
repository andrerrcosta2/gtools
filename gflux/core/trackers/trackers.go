// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package trackers

import (
	"github.com/andrerrcosta2/gtools/core/data/comparators"
	"sort"
	"strings"
	"unsafe"
)

func Instance() *InstanceTracker {
	return &InstanceTracker{
		instances: make(map[unsafe.Pointer]bool),
	}
}

type InstanceTracker struct {
	instances map[unsafe.Pointer]bool
}

// Has checks if an instance has already been visited
func (t *InstanceTracker) Has(inst any) bool {
	if inst == nil {
		return true
	}

	ptr := extractPointer(inst)
	_, ok := t.instances[ptr]
	return ok
}

// Mark marks an instance as visited
func (t *InstanceTracker) Mark(inst any) {
	if inst != nil {
		ptr := extractPointer(inst)
		t.instances[ptr] = true
	}
}

func Operation[D any, R any](comparator comparators.KeyTyped[D, string]) *OpTracker[D, R] {
	return &OpTracker[D, R]{
		cmp: comparator,
		ops: make(map[string]R),
	}
}

type OpTracker[D any, R any] struct {
	cmp comparators.KeyTyped[D, string]
	ops map[string]R // Tracks pairs of objects by their memory addresses
}

// Has checks if the pair (a, b) has already been compared
func (t *OpTracker[D, R]) Has(a ...D) (cached R, ok bool) {
	key := t.computeKey(a...)
	if op, exists := t.ops[key]; exists {
		return op, true
	}
	return cached, false
}

// Mark marks the pair (a, b) as compared and stores the result
func (t *OpTracker[D, R]) Mark(result R, a ...D) R {
	key := t.computeKey(a...)
	t.ops[key] = result
	return result
}

func (t *OpTracker[D, R]) computeKey(a ...D) string {
	if len(a) == 0 {
		return "<empty>"
	}

	hashes := make([]string, 0, len(a))
	for _, x := range a {
		hash := t.cmp.Hash(x)
		hashes = append(hashes, hash)
	}

	sort.Strings(hashes)
	return strings.Join(hashes, ",")
}

// extractPointer returns the memory address of the value, so it won’t panic.
// But for values, it won’t track correctly (won’t prevent redundant comparisons)
// Since Go doesn’t allow value recursion, this is harmless.
func extractPointer(x any) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&x))
}
