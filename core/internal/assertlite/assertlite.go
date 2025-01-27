// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

//go:build tt

package assertlite

import (
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
