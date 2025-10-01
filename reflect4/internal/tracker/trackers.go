// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tracker

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"reflect"
	"unsafe"
)

func newCmpKey(v1, v2 reflect.Value) cmpKey {
	addr1, addr2 := v1.UnsafePointer(), v2.UnsafePointer()
	if uintptr(addr1) > uintptr(addr2) {
		addr1, addr2 = addr2, addr1
	}
	return cmpKey{
		addr1: addr1,
		addr2: addr2,
		typ:   v1.Type(),
	}
}

type cmpKey struct {
	addr1 unsafe.Pointer
	addr2 unsafe.Pointer
	typ   reflect.Type
}

func newRefKey(v reflect.Value) refKey {
	return refKey{
		addr: v.UnsafePointer(),
		typ:  v.Type(),
	}
}

type refKey struct {
	addr unsafe.Pointer
	typ  reflect.Type
}

func Eq() *EqTracker {
	return &EqTracker{
		nodes: make(map[cmpKey]bool),
	}
}

type EqTracker struct {
	nodes map[cmpKey]bool
}

func (t *EqTracker) Check(a, b reflect.Value) bool {
	_, ok := t.nodes[newCmpKey(a, b)]
	return ok
}

func (t *EqTracker) Mark(a, b reflect.Value) {
	t.nodes[newCmpKey(a, b)] = true
}

// Reference creates a new reference tracker.
func Reference() *RefTracker {
	return &RefTracker{cache: make(map[refKey]reflect.Value)}
}

type RefTracker struct {
	cache map[refKey]reflect.Value
}

func (t *RefTracker) Get(v reflect.Value) (reflect.Value, bool) {
	key := newRefKey(v)
	cache, ok := t.cache[key]
	return cache, ok
}

func (t *RefTracker) Mark(inst, cache reflect.Value) reflect.Value {
	key := newRefKey(inst)
	t.cache[key] = cache
	return cache
}

func Sprint() *SprintTracker {
	return &SprintTracker{
		cache: make(map[refKey]string),
	}
}

type SprintTracker struct {
	cache map[refKey]string
}

// Get checks if an instance has already been visited
func (t *SprintTracker) Get(inst reflect.Value) (string, bool) {
	key := newRefKey(inst)
	sprint, ok := t.cache[key]
	return sprint, ok
}

// Mark marks an instance as visited and stores its sprint for further use
func (t *SprintTracker) Mark(inst reflect.Value, sprint string) string {
	key := newRefKey(inst)
	t.cache[key] = sprint
	return sprint
}

func Diff() *DiffTracker {
	return &DiffTracker{
		cache: make(map[cmpKey]differs.Difference),
	}
}

type DiffTracker struct {
	cache map[cmpKey]differs.Difference // Tracks cache of objects by their memory address
}

// Get checks if the pair (a, b) has already been compared
func (t *DiffTracker) Get(a, b reflect.Value) (diff differs.Difference, ok bool) {
	key := newCmpKey(a, b)
	//fmx.Redf("[DiffTracker] Comparing: %v\n", key)
	if a.Kind() != b.Kind() {
		return differs.Difference{
			Message: differs.TypesMismatch(indent.Zero(), a.Kind().String(), b.Kind().String()),
			Equals:  false,
		}, true
	}
	if cache, ok := t.cache[key]; ok {
		return cache, true
	}
	return // Not found

}

// Mark marks the pair as compared and stores the result
func (t *DiffTracker) Mark(a, b reflect.Value, result differs.Difference) differs.Difference {
	key := newCmpKey(a, b)
	t.cache[key] = result
	return result
}
