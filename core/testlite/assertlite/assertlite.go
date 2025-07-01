// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package assertlite

import (
	"errors"
	"fmt"
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/testlite/internal/ask"
	"reflect"
)

// AllFieldsAreNil asserts all the struct fields are nil.
// This assertion performs only shallow inspection.
// It flags the test as failed if the value isn't a struct or if it contains any non nil field.
func AllFieldsAreNil(t HelperTesting, checkUnexported bool, target any, msgAndArgs ...any) bool {
	var has bool
	var nf []string
	var err error
	var tv = reflect.ValueOf(target)
	if checkUnexported {
		has, nf, err = ask.HasAnyNonNilField(tv)
	} else {
		has, nf, err = ask.HasExportedNilField(tv)
	}
	if err != nil {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. %v", err), msgAndArgs...)
		return false
	} else if has {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. expected no nil fields but got nil for "+
			"'%v' on '%s%v'\n", nf, tv.String(), target), msgAndArgs...)
		return false
	}
	return true
}

func AreFalse[T any](t HelperTesting, data []T, fn functions.Function[T, bool], msgAndArgs ...any) bool {
	for _, d := range data {
		if fn(d) {
			t.Helper()
			fail(t, fmt.Sprintf("expected all elements to be false, but got true for '%v'\n", d), msgAndArgs...)
			return false
		}
	}
	return true
}

func AreNotNil[T any](t HelperTesting, values []T, msgAndArgs ...any) bool {
	if values == nil {
		t.Helper()
		fail(t, "❌ assertion failed. slice is nil\n", msgAndArgs...)
		return false
	}

	var ni = make([]int, 0, len(values))
	for i, value := range values {
		if !ask.IsNil(reflect.ValueOf(value)) {
			ni = append(ni, i)
		}
	}
	if len(ni) > 0 {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. expected no nil values but found at indexes %v", ni), msgAndArgs...)
		return false
	}
	return true
}

func AreTrue[T any](t HelperTesting, data []T, fn functions.Function[T, bool], msgAndArgs ...any) bool {
	for _, d := range data {
		if !fn(d) {
			t.Helper()
			fail(t, fmt.Sprintf("expected all elements to be true, but got false for '%v'\n", d), msgAndArgs...)
			return false
		}
	}
	return true
}

func EqualMaps[M ~map[K]V, K comparable, V any](t HelperTesting, a, b M, msgAndArgs ...any) bool {
	if !ask.AreEqualMaps(a, b) {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. maps are not equals: a = <%T>%v, "+
			"b = <%T>%v\n", a, a, b, b), msgAndArgs...)
		return false
	}
	return true
}

func Equals(t HelperTesting, a, b any, msgAndArgs ...any) bool {
	if !ask.AreEquals(a, b) {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. values are not equals: a = <%T>%v, "+
			"b = <%T>%v\n", a, a, b, b), msgAndArgs...)
		return false
	}
	return true
}

func EqualSlices[T any](t HelperTesting, a, b []T, msgAndArgs ...any) bool {
	if a == nil || b == nil {
		if a == nil && b == nil {
			return true
		}
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. expected slices to be equal, "+
			"but one of them is nil: a: '%v', b: '%v'", a, b))
		return false
	}
	// Check if the slices have different lengths
	if len(a) != len(b) {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed.  slices to be equal, but got different"+
			" lengths: a = %v, b = %v\n",
			len(a), len(b)), msgAndArgs...)
		return false
	}

	diff := make([]struct {
		idx int
		a   T
		b   T
	}, 0)

	va, vb := reflect.ValueOf(a), reflect.ValueOf(b)
	knd := va.Type().Elem().Kind()
	switch knd {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		for i := 0; i < len(a); i++ {
			if va.Index(i).IsNil() || vb.Index(i).IsNil() {
				if !va.Index(i).IsNil() || !vb.Index(i).IsNil() {
					diff = append(diff, struct {
						idx int
						a   T
						b   T
					}{
						idx: i,
						a:   a[i],
						b:   b[i],
					})
				}
			}
			if !reflect.DeepEqual(a[i], b[i]) {
				t.Helper()
				diff = append(diff, struct {
					idx int
					a   T
					b   T
				}{
					idx: i,
					a:   a[i],
					b:   b[i],
				})
			}
		}
	default:
		for i := 0; i < len(a); i++ {
			if !reflect.DeepEqual(a[i], b[i]) {
				t.Helper()
				diff = append(diff, struct {
					idx int
					a   T
					b   T
				}{
					idx: i,
					a:   a[i],
					b:   b[i],
				})
			}
		}
	}
	if len(diff) > 0 {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. expected slices to be equal, "+
			"but got different values at %v\n", diff), msgAndArgs...)
		return false
	}
	return true
}

// False asserts that the given condition is false.
// Returns true if the given condition is false
// Otherwise returns false and flags the test as failed
func False(t HelperTesting, b bool, msgAndArgs ...any) bool {
	if b {
		t.Helper()
		fail(t, "expected false, but got true\n", msgAndArgs...)
		return false
	}
	return true
}

// IsErrorOf asserts that the error is not nil and any error in error's tree matches
// the target.
func IsErrorOf(t HelperTesting, err error, expected error, msgAndArgs ...any) {
	t.Helper()
	if err == nil {
		fail(t, "expected error, but got nil\n", msgAndArgs...)
	}
	if !errors.Is(err, expected) {
		fail(t, fmt.Sprintf("expected error to be '%v', but got '%v'\n", expected, err), msgAndArgs...)
	}
}

// IsNil asserts the value is nil or flag the test as failed.
func IsNil(t HelperTesting, value any, msgAndArgs ...any) bool {
	if value == nil {
		return true
	}
	if !ask.IsNil(reflect.ValueOf(value)) {
		t.Helper()
		fail(t, "❌ assertion failed. target is not nil", msgAndArgs...)
		return false
	}
	return true
}

// IsTypeOf asserts the value is of the given type or flag the test as failed.
func IsTypeOf[T any](t HelperTesting, v any, msgAndArgs ...any) bool {
	if _, ok := v.(T); !ok {
		t.Helper()
		typx := reflect.TypeOf((*T)(nil)).Elem()
		fail(t, fmt.Sprintf("❌ assertion failed. expected data to be of type '%s',"+
			" but got '%T'\n", typx, v), msgAndArgs...)
		return false
	}
	return true
}

// NoError asserts that the given error is nil or flag the test as failed.
func NoError(t HelperTesting, err error, msgAndArgs ...any) bool {
	if err != nil {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. no error was expected, but got '%v'\n", err), msgAndArgs...)
		return false
	}
	return true
}

// NoNilFields asserts the struct has no nil fields.
// This assertion performs only shallow inspection.
// It flags the test as failed if the value isn't a struct or if it contains any nil field.
func NoNilFields(t HelperTesting, checkUnexported bool, target any, msgAndArgs ...any) bool {
	var has bool
	var nf []string
	var err error
	var tv = reflect.ValueOf(target)
	if checkUnexported {
		has, nf, err = ask.HasAnyNilField(tv)
	} else {
		has, nf, err = ask.HasExportedNilField(tv)
	}
	if err != nil {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. %v", err), msgAndArgs...)
		return false
	} else if has {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. expected no nil fields but got nil for "+
			"'%v' on '%s%v'\n", nf, tv.String(), target), msgAndArgs...)
		return false
	}
	return true
}

// NoNilNonInterfaceFields asserts the target has no nil fields except for interfaces with at least one method
func NoNilNonInterfaceFields(t HelperTesting, checkUnexported bool, target any, msgAndArgs ...any) bool {
	var has bool
	var nf []string
	var err error
	var tv = reflect.ValueOf(target)
	if checkUnexported {
		has, nf, err = ask.HasAnyNonInterfaceNilField(tv)
	} else {
		has, nf, err = ask.HasExportedNonInterfaceNilField(tv)
	}
	if err != nil {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. %v", err), msgAndArgs...)
		return false
	} else if has {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. expected no nil fields but got nil for "+
			"'%v' on '%s%v'\n", nf, tv.String(), target), msgAndArgs...)
		return false
	}
	return true
}

// NotEmpty asserts that the given string is not empty or flag the test as failed.
func NotEmpty(t HelperTesting, v string, msgAndArgs ...any) bool {
	if v == "" {
		t.Helper()
		fail(t, "❌ assertion failed. string is empty", msgAndArgs...)
		return false
	}
	return true
}

// NotEquals asserts the given values are not equals or flag the test as failed.
func NotEquals(t HelperTesting, a, b any, msgAndArgs ...any) bool {
	if ask.AreEquals(a, b) {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. values are equals: a = <%T>%v, "+
			"b = <%T>%v\n", a, a, b, b), msgAndArgs...)
		return false
	}
	return true
}

// NotNil asserts that the given value is not nil or flag the test as failed.
func NotNil(t HelperTesting, value any, msgAndArgs ...any) bool {
	if value == nil {
		t.Helper()
		fail(t, "❌ assertion failed. target is nil\n", msgAndArgs...)
		return false
	}
	if ask.IsNil(reflect.ValueOf(value)) {
		t.Helper()
		fail(t, "❌ assertion failed. target is nil", msgAndArgs...)
		return false
	}
	return true
}

// NoPanic asserts that the function doesn't panic or flag the test as failed
func NoPanic(t HelperTesting, f func(), msgAndArgs ...any) (res bool) {
	t.Helper()
	defer noPanic(t, &res, msgAndArgs...)
	f()
	return
}

func noPanic(t HelperTesting, res *bool, msgAndArgs ...any) {
	t.Helper()
	o := true
	if r := recover(); r != nil {
		fail(t, fmt.Sprintf("expected no panic, but got 'panic: %v'\n", r), msgAndArgs...)
		o = false
	}
	res = &o
}

// NotSame asserts that two references aren't pointing to the same object.
// It'll never return true for non-reference types
func NotSame(t HelperTesting, a, b any, msgAndArgs ...any) bool {
	same, addr, _ := ask.AreSameInstances(a, b)
	if same {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. expected data to be different instances,"+
			" but got same addresses: a = <%#x>, b = <%#x>\n", addr[0], addr[1]), msgAndArgs...)
		return false
	}
	return true
}

// NotTypeOf asserts that the value is not of the given type or flag the test as failed.
func NotTypeOf[T any](t HelperTesting, v any, msgAndArgs ...any) bool {
	if _, ok := v.(T); ok {
		t.Helper()
		typx := reflect.TypeOf((*T)(nil)).Elem()
		fail(t, fmt.Sprintf("❌ assertion failed. expected data not to be of "+
			"type '%s', but got '%T'\n", typx, v), msgAndArgs...)
		return false
	}
	return true
}

// Panic asserts that the function panics or flag the test as failed
func Panic(t HelperTesting, f func(), msgAndArgs ...any) (res bool) {
	t.Helper()
	var panicErr any
	defer func() {
		if r := recover(); r != nil {
			panicErr = r
		}
	}()
	f()

	if panicErr == nil {
		fail(t, "expected a panic, but got none\n", msgAndArgs...)
		return false
	}
	return true
}

// Same asserts that two references are pointing to the same object.
// It'll never return true for non-reference types
func Same(t HelperTesting, a, b any, msgAndArgs ...any) bool {
	same, addr, err := ask.AreSameInstances(a, b)
	if err != nil {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. %v", err), msgAndArgs...)
		return false
	} else if !same {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. expected data to be the same instance,"+
			" but got different instances: a = <%#x>, b = <%#x>\n", addr[0], addr[1]), msgAndArgs...)
		return false
	}
	return true
}

// True asserts the condition is true or flag the test as failed.
func True(t HelperTesting, condition bool, msgAndArgs ...any) bool {
	if !condition {
		t.Helper()
		fail(t, "expected true, but got false\n", msgAndArgs...)
		return false
	}
	return true
}
