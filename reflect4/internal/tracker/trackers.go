// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package tracker

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/reflect4/internal/pointers"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"reflect"
)

// Reference creates a new reference tracker.
func Reference() *RefTracker {
	return &RefTracker{pointers: make(map[uintptr]*ref)}
}

type ref struct {
	Value reflect.Value
	Key   reflect.Value
}
type RefTracker struct {
	pointers map[uintptr]*ref
}

func (t *RefTracker) Get(v reflect.Value) (reflect.Value, bool, error) {
	addr, err := pointers.AddrOf(v)
	if err != nil {
		// If it isn't a stable pointer, it can't produce cyclic references directly
		return values.Empty, false, err
	}
	str := v.String()
	fmx.Printf("[get]: <%d>%s\n", addr, str)
	sprint, ok := t.pointers[addr]
	if ok {
		cstr := sprint.Value.String()
		ckey := sprint.Key.String()
		fmx.Printf("[get] returning: %s%s\n", ckey, cstr)
		return sprint.Value, true, nil
	}
	return values.Empty, false, nil
}

func (t *RefTracker) Mark(inst, cache reflect.Value) (reflect.Value, error) {
	addr, err := pointers.AddrOf(inst)
	if err != nil {
		// If it isn't a stable pointer, it can't produce cyclic references directly
		// so it shouldn't shortcut its object traversal.
		return values.Empty, err
	}
	str := inst.String()
	fmx.Printf("[mark] inst: <%d>%s\n", addr, str)
	cstr := cache.String()
	fmx.Printf("[mark] cache: %s\n", cstr)
	t.pointers[addr] = &ref{Value: cache, Key: inst}
	return cache, nil
}

func Sprint() *SprintTracker {
	return &SprintTracker{
		pointers: make(map[uintptr]string),
		values:   make(map[uintptr]string),
	}
}

type SprintTracker struct {
	pointers map[uintptr]string
	values   map[uintptr]string
}

// Get checks if an instance has already been visited
func (t *SprintTracker) Get(inst reflect.Value) (string, bool, error) {
	addr, err := pointers.AddrOf(inst)
	if err != nil {
		return sprints.Errorf(indent.Zero(), "%s", err), true, err
	}

	if inst.Kind() == reflect.Ptr {
		sprint, ok := t.pointers[addr]
		return sprint, ok, nil
	}

	sprint, ok := t.values[addr]
	return sprint, ok, nil
}

// Mark marks an instance as visited and stores its sprint for further use
func (t *SprintTracker) Mark(inst reflect.Value, sprint string) error {
	addr, err := pointers.AddrOf(inst)
	if err != nil {
		return err
	}

	if inst.Kind() == reflect.Ptr {
		t.pointers[addr] = sprint
	} else {
		t.values[addr] = sprint
	}
	return nil
}

func Diff() *DiffTracker {
	return &DiffTracker{
		pointers: make(map[uintptr]map[uintptr]differs.Difference),
		values:   make(map[uintptr]map[uintptr]differs.Difference),
	}
}

type DiffTracker struct {
	pointers map[uintptr]map[uintptr]differs.Difference // Tracks pointers of objects by their memory addresses
	values   map[uintptr]map[uintptr]differs.Difference // Tracks values of objects by their memory addresses
}

// Get checks if the pair (a, b) has already been compared
func (t *DiffTracker) Get(a, b reflect.Value) (diff differs.Difference, ok bool) {
	var rec, exp string
	addrA, err := pointers.AddrOf(a)
	if err != nil {
		rec = sprints.Errorf(indent.Zero(), "%s", err)
	}
	addrB, err := pointers.AddrOf(b)
	if err != nil {
		exp = sprints.Errorf(indent.Zero(), "%s", err)
	}
	if rec != "" || exp != "" {
		return differs.Difference{
			Message: "invalid pointer",
			Diff:    differs.Diff(indent.Zero(), rec, exp),
			Equals:  true,
			Err:     err,
		}, true // true means it should return this value
		// avoiding recursion over an invalid pointer
	}

	if a.Kind() != b.Kind() {
		return t.diffKnd(a, b, addrA, addrB)
	}

	if a.Kind() == reflect.Ptr {
		if inner, ok := t.pointers[addrA]; ok {
			if result, exists := inner[addrB]; exists {
				return result, true // Found cached result
			}
		}
		return // Not found
	}

	if innerMap, ok := t.values[addrA]; ok {
		if result, exists := innerMap[addrB]; exists {
			return result, true // Found cached result
		}
	}
	return // Not found
}

func (t *DiffTracker) diffKnd(a, b reflect.Value, addrA, addrB uintptr) (differs.Difference, bool) {
	if res, ok := t.hasPair(addrA, addrB); ok {
		return res, true
	}
	if inner, ok := t.values[addrA]; ok {
		inner[addrB] = differs.Difference{
			Message: differs.TypesMismatch(indent.Zero(), a.Kind().String(), b.Kind().String()),
			Equals:  false,
		}
		return differs.Difference{
			Message: differs.TypesMismatch(indent.Zero(), a.Kind().String(), b.Kind().String()),
			Equals:  false,
		}, true
	}
	res := differs.Difference{
		Message: differs.TypesMismatch(indent.Zero(), a.Kind().String(), b.Kind().String()),
		Equals:  false,
	}
	inner := make(map[uintptr]differs.Difference)
	inner[addrB] = res
	t.values[addrA] = inner
	return res, true // true means it should return this value
}

func (t *DiffTracker) hasPair(a, b uintptr) (differs.Difference, bool) {
	if inner, ok := t.values[a]; ok {
		if result, exists := inner[b]; exists {
			return result, true // Found cached result
		}
	} else if inner, ok := t.pointers[a]; ok {
		if result, exists := inner[b]; exists {
			return result, true // Found cached result
		}
	}
	return differs.Difference{}, false
}

// Mark marks the pair (a, b) as compared and stores the result
func (t *DiffTracker) Mark(a, b reflect.Value, result differs.Difference) differs.Difference {
	var rec, exp string
	addrA, err := pointers.AddrOf(a)
	if err != nil {
		rec = sprints.Errorf(indent.Zero(), "%s", err)
	}
	addrB, err := pointers.AddrOf(b)
	if err != nil {
		exp = sprints.Errorf(indent.Zero(), "%s", err)
	}
	if rec != "" || exp != "" {
		return differs.Difference{
			Message: "invalid pointer",
			Diff:    differs.Diff(indent.Zero(), rec, exp),
			Equals:  true,
			Err:     err,
		}
	}

	if a.Kind() != b.Kind() {
		return differs.Difference{
			Message: differs.TypesMismatch(indent.Zero(), a.Kind().String(), b.Kind().String()),
			Equals:  false,
		}
	}
	if a.Kind() != b.Kind() {
		if diff, ok := t.diffKnd(a, b, addrA, addrB); ok {
			return diff
		}
	}

	if a.Kind() == reflect.Ptr {
		if _, ok := t.pointers[addrA]; !ok {
			t.pointers[addrA] = make(map[uintptr]differs.Difference)
		}
		t.pointers[addrA][addrB] = result
		return result
	}

	if _, ok := t.values[addrA]; !ok {
		t.values[addrA] = make(map[uintptr]differs.Difference)
	}
	t.values[addrA][addrB] = result
	return result
}
