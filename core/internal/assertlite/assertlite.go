// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

//go:build tt

package assertlite

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"reflect"
)

func True(t HelperTesting, b bool, msgAndArgs ...any) {
	t.Helper()
	if !b {
		fail(t, "expected true, but got false\n", msgAndArgs...)
	}
}

func False(t HelperTesting, b bool, msgAndArgs ...any) {
	t.Helper()
	if b {
		fail(t, "expected false, but got true\n", msgAndArgs...)
	}
}

func AllTrue[T any](t HelperTesting, data []T, fn functions.Function[T, bool], msgAndArgs ...any) {
	t.Helper()
	for _, d := range data {
		if !fn(d) {
			fail(t, fmt.Sprintf("expected all elements to be true, but got false for '%v'\n", d), msgAndArgs...)
		}
	}
}

func Panic(t HelperTesting, f func(), msgAndArgs ...any) {
	t.Helper()
	defer func() {
		if r := recover(); r == nil {
			fail(t, "expected panic, but got none\n", msgAndArgs...)
		}
	}()
	f()
}

func Equal(t HelperTesting, a, b any, msgAndArgs ...any) {
	t.Helper()
	// Handle nil cases first
	if a == nil || b == nil {
		if a != b {
			fail(t, "expected both values to be nil\n", msgAndArgs...)
		}
		return
	}

	// Check if both are byte slices
	if ba, ok := a.([]byte); ok {
		if bb, ok := b.([]byte); ok {
			if !bytes.Equal(ba, bb) {
				fail(t, fmt.Sprintf("expected values to be equal, but got different values: a = %v, b = %v\n", ba, bb), msgAndArgs...)
			}
			return
		}
		// Only fail if one is []byte but not the other
		fail(t, fmt.Sprintf("expected values to be of the same type, but got: a = %T, b = %T\n", a, b), msgAndArgs...)
	}

	// If neither are []byte, use DeepEqual
	if !reflect.DeepEqual(a, b) {
		fail(t, fmt.Sprintf("expected values to be equal, but got different values: a = %v, b = %v\n", a, b), msgAndArgs...)
	}

}

func NoError(t HelperTesting, err error, msgAndArgs ...any) {
	t.Helper()
	if err != nil {
		fail(t, fmt.Sprintf("no error was expected, but got '%v'\n", err), msgAndArgs...)
	}
}

func IsTypeOf[T any](t HelperTesting, v any, msgAndArgs ...any) {
	t.Helper()
	if _, ok := v.(T); !ok {
		fail(t, fmt.Sprintf("expected data to be of type '%T', but got '%T'\n", v, v), msgAndArgs...)
	}
}

func NotTypeOf[T any](t HelperTesting, v any, msgAndArgs ...any) {
	t.Helper()
	if _, ok := v.(T); ok {
		typx := reflect.TypeOf((*T)(nil)).Elem()
		fail(t, fmt.Sprintf("expected data not to be of type '%s', but got '%T'\n", typx, v), msgAndArgs...)
	}
}

func ArrayEquals[T comparable](t HelperTesting, a, b []T, msgAndArgs ...any) {
	t.Helper()
	wa, wb := a, b
	// Check if the slices have different lengths
	if len(wa) != len(wb) {
		fail(t, fmt.Sprintf("expected slices to be equal, but got different lengths: a = %v, b = %v\n", len(wa), len(wb)), msgAndArgs...)
	}

	// Compare each element of the slices
	for i := range wa {
		if wa[i] != wb[i] {
			fail(t, fmt.Sprintf("expected slices to be equal, but got different elements at index %d: a = %v, b = %v\n", i, wa[i], wb[i]), msgAndArgs...)
		}
	}
}

func NotNil(t HelperTesting, value any, msgAndArgs ...any) {
	t.Helper()
	if value == nil {
		fail(t, "expected value not to be nil\n", msgAndArgs...)
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v.IsNil() {
			fail(t, "expected value not to be nil\n", msgAndArgs...)
		}
	default:
	}
}

func ErrorIs(t HelperTesting, err error, expected error, msgAndArgs ...any) {
	t.Helper()
	if err == nil {
		fail(t, "expected error, but got nil\n", msgAndArgs...)
	}
	if !errors.Is(err, expected) {
		fail(t, fmt.Sprintf("expected error to be '%v', but got '%v'\n", expected, err), msgAndArgs...)
	}
}
