// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sortables

import (
	"fmt"
	"reflect"

	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

type Semantics interface {
	Equal(a, b any) (bool, bool)
	Less(a, b any) (bool, bool)
	Fingerprint(v any) (string, bool)
}

type Engine struct {
	resolvers []Semantics
}

func NewEngine(resolvers ...Semantics) *Engine {
	return &Engine{resolvers: resolvers}
}

func (e *Engine) Equal(a, b any) bool {
	for _, r := range e.resolvers {
		if result, ok := r.Equal(a, b); ok {
			return result
		}
	}
	return false
}

func (e *Engine) Less(a, b any) bool {
	for _, r := range e.resolvers {
		if result, ok := r.Less(a, b); ok {
			return result
		}
	}
	return false
}

func (e *Engine) Compare(a, b any) int {
	if e.Equal(a, b) {
		return 0
	}
	if e.Less(a, b) {
		return -1
	}
	return 1
}

func (e *Engine) Fingerprint(v any) string {
	for _, r := range e.resolvers {
		if result, ok := r.Fingerprint(v); ok {
			return result
		}
	}
	return fallbackFingerprint(v)
}

type comparableSemantics struct{}

func (comparableSemantics) Equal(a, b any) (bool, bool) {
	if x, ok := a.(interface{ Equal(any) bool }); ok {
		return x.Equal(b), true
	}
	if y, ok := b.(interface{ Equal(any) bool }); ok {
		return y.Equal(a), true
	}
	return false, false
}

func (comparableSemantics) Less(a, b any) (bool, bool) {
	if x, ok := a.(interface{ Less(any) bool }); ok {
		return x.Less(b), true
	}
	if y, ok := b.(interface{ Less(any) bool }); ok {
		return y.Less(a), true
	}
	return false, false
}

func (comparableSemantics) Fingerprint(v any) (string, bool) {
	return "", false
}

type primitiveSemantics struct{}

func (primitiveSemantics) Equal(a, b any) (bool, bool) {
	if eq, err := prim.TryCompare(a, b); err == nil {
		return eq, true
	}
	return false, false
}

func (primitiveSemantics) Less(a, b any) (bool, bool) {
	if less, err := prim.TryLess(a, b); err == nil {
		return less, true
	}
	return false, false
}

func (primitiveSemantics) Fingerprint(v any) (string, bool) {
	return primitiveFingerprint(v)
}

type identitySemantics struct{}

func (identitySemantics) Equal(a, b any) (bool, bool) {
	ia, okA := a.(interface{ ID() string })
	ib, okB := b.(interface{ ID() string })

	if okA && okB {
		return ia.ID() == ib.ID(), true
	}
	return false, false
}

func (identitySemantics) Less(a, b any) (bool, bool) {
	return false, false
}

func (identitySemantics) Fingerprint(v any) (string, bool) {
	if id, ok := v.(interface{ ID() string }); ok {
		return id.ID(), true
	}
	return "", false
}

type pointerSemantics struct{}

func (pointerSemantics) Equal(a, b any) (bool, bool) {
	return a == b, true
}

func (pointerSemantics) Less(a, b any) (bool, bool) {
	return false, false
}

func (pointerSemantics) Fingerprint(v any) (string, bool) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr && !rv.IsNil() {
		return fmt.Sprintf("0x%x", rv.Pointer()), true
	}
	return "", false
}

type reflectSemantics struct{}

func (reflectSemantics) Equal(a, b any) (bool, bool) {
	// DeepEqual handles nested structures, dicts, and slices.
	return reflect.DeepEqual(a, b), true
}

func (reflectSemantics) Less(a, b any) (bool, bool) {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	// 1. Compare by Type/Kind first to ensure a stable sort
	if va.Kind() != vb.Kind() {
		return va.Kind() < vb.Kind(), true
	}

	// 2. Handle specific kinds that have a logical "Compare"
	switch va.Kind() {
	case reflect.String:
		return va.String() < vb.String(), true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return va.Int() < vb.Int(), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return va.Uint() < vb.Uint(), true
	case reflect.Float32, reflect.Float64:
		return va.Float() < vb.Float(), true
	case reflect.Slice, reflect.Array, reflect.Map:
		// Compare by length as a fallback
		if va.Len() != vb.Len() {
			return va.Len() < vb.Len(), true
		}
		// Deep comparison for slices/array could be added here if needed,
		// but it's computationally expensive.
		return false, true
	}

	return false, false
}

func (reflectSemantics) Fingerprint(v any) (string, bool) {
	return fmt.Sprintf("%T(%v)", v, v), true
}
