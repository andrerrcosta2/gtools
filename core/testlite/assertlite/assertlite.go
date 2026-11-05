// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package assertlite

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/andrerrcosta2/gtools/core/domain/function"
	"github.com/andrerrcosta2/gtools/core/testlite"
	"github.com/andrerrcosta2/gtools/core/testlite/internal/ask"
)

// AllFieldsAreNil asserts all the struct fields are nil.
// This assertion performs only shallow inspection.
// It flags the test as failed if the value isn't a struct or if it contains any non nil field.
func AllFieldsAreNil(t testlite.HelperTesting, checkUnexported bool, target any, msgAndArgs ...any) bool {
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

// AreFalse asserts the predicate returns false for every element.
func AreFalse[T any](t testlite.HelperTesting, data []T, fn function.Function[T, bool], msgAndArgs ...any) bool {
	for _, d := range data {
		if fn(d) {
			t.Helper()
			fail(t, fmt.Sprintf("❌ assertion failed. expected all elements to be false, but got true for '%v'\n", d), msgAndArgs...)
			return false
		}
	}
	return true
}

// AreNotNil asserts the slice itself is non-nil and has no nil elements.
func AreNotNil[S ~[]E, E any](t testlite.HelperTesting, values S, msgAndArgs ...any) bool {
	if values == nil {
		t.Helper()
		fail(t, "❌ assertion failed. slice is nil\n", msgAndArgs...)
		return false
	}

	var ni = make([]int, 0, len(values))
	for i, value := range values {
		if ask.IsNil(reflect.ValueOf(value)) {
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

// AreTrue asserts the predicate returns true for every element.
func AreTrue[T any](t testlite.HelperTesting, data []T, fn function.Function[T, bool], msgAndArgs ...any) bool {
	for _, d := range data {
		if !fn(d) {
			t.Helper()
			fail(t, fmt.Sprintf("❌ assertion failed. expected all elements to be true, but got "+
				"false for '%v'\n", d), msgAndArgs...)
			return false
		}
	}
	return true
}

// Contains asserts that elem exists in slice a.
func Contains[S ~[]E, E comparable](t testlite.HelperTesting,
	a S, elem E,
	msgAndArgs ...any) bool {
	for _, e := range a {
		if elem == e {
			return true
		}
	}
	t.Helper()
	fail(t, fmt.Sprintf("❌ assertion failed. expected elements to exist inside slice, but element '%v' "+
		"wasn't found\n", elem), msgAndArgs...)
	return false
}

type sliceDiff[T any] struct {
	idx int
	a   T
	b   T
}

func (s sliceDiff[T]) String() string {
	return fmt.Sprintf("{ idx: %v, a: %v, b: %v }", s.idx, s.a, s.b)
}

// ElementsMatch asserts two slices contain the same elements, regardless of order.
// Element multiplicity is respected.
func ElementsMatch[S ~[]E, E any](
	t testlite.HelperTesting,
	a, b S,
	msgAndArgs ...any,
) bool {
	t.Helper()

	// nil handling
	if a == nil || b == nil {
		if a == nil && b == nil {
			return true
		}
		fail(t, fmt.Sprintf(
			"❌ assertion failed. expected slices to match elements, but one is nil: a=%v, b=%v",
			a, b,
		), msgAndArgs...)
		return false
	}

	// length check
	if len(a) != len(b) {
		fail(t, fmt.Sprintf(
			"❌ assertion failed. expected slices to have same elements, but lengths differ: a=%d, b=%d",
			len(a), len(b),
		), msgAndArgs...)
		return false
	}

	used := make([]bool, len(b))
	unmatchedA := make([]E, 0)

	for _, av := range a {
		found := false
		for i, bv := range b {
			if used[i] {
				continue
			}
			if reflect.DeepEqual(av, bv) {
				used[i] = true
				found = true
				break
			}
		}
		if !found {
			unmatchedA = append(unmatchedA, av)
		}
	}

	if len(unmatchedA) > 0 {
		unmatchedB := make([]E, 0)
		for i, used := range used {
			if !used {
				unmatchedB = append(unmatchedB, b[i])
			}
		}

		fail(t, fmt.Sprintf(
			"❌ assertion failed. slices do not match elements.\n"+
				"unmatched in a: %v\n"+
				"unmatched in b: %v",
			unmatchedA,
			unmatchedB,
		), msgAndArgs...)
		return false
	}

	return true
}

// EqualMaps asserts two dicts have equal keys and values.
func EqualMaps[M ~map[K]V, K comparable, V any](t testlite.HelperTesting, a, b M, msgAndArgs ...any) bool {
	if !ask.AreEqualMaps(a, b) {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. dicts are not equals: a = <%T>%v, "+
			"b = <%T>%v\n", a, a, b, b), msgAndArgs...)
		return false
	}
	return true
}

// Equals asserts two values are deeply equal.
func Equals(t testlite.HelperTesting, a, b any, msgAndArgs ...any) bool {
	if !ask.AreEquals(a, b) {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. values are not equals: a = <%T>%v, "+
			"b = <%T>%v\n", a, a, b, b), msgAndArgs...)
		return false
	}
	return true
}

// EqualSlices asserts two slices are equal by length and element order.
// Nil and empty slices are considered different.
func EqualSlices[S ~[]E, E any](t testlite.HelperTesting, a, b S, msgAndArgs ...any) bool {
	if a == nil || b == nil {
		if a == nil && b == nil {
			return true
		}
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. expected slices to be equals, "+
			"but one of them is nil: a: '%v', b: '%v'", a, b))
		return false
	}
	// Check if the slices have different lengths
	if len(a) != len(b) {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. expected slices to be equals, but got different"+
			" lengths: a = %d:%v, b = %d:%v\n",
			len(a), a, len(b), b), msgAndArgs...)
		return false
	}

	diff := make([]sliceDiff[E], 0)

	va, vb := reflect.ValueOf(a), reflect.ValueOf(b)
	knd := va.Type().Elem().Kind()
	switch knd {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		for i := 0; i < len(a); i++ {
			if va.Index(i).IsNil() || vb.Index(i).IsNil() {
				if !va.Index(i).IsNil() || !vb.Index(i).IsNil() {
					diff = append(diff, sliceDiff[E]{
						idx: i,
						a:   a[i],
						b:   b[i],
					})
				}
			}
			if !reflect.DeepEqual(a[i], b[i]) {
				t.Helper()
				diff = append(diff, sliceDiff[E]{
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
				diff = append(diff, sliceDiff[E]{
					idx: i,
					a:   a[i],
					b:   b[i],
				})
			}
		}
	}
	if len(diff) > 0 {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. expected slices to be equals, "+
			"but got different values at %v\n", diff), msgAndArgs...)
		return false
	}
	return true
}

// Eventually asserts condition becomes true before timeout.
// The condition is polled every 5ms.
func Eventually(t testlite.HelperTesting, timeout time.Duration, condition func() bool, msgAndArgs ...any) bool {
	start := time.Now()
	for {
		if condition() {
			return true
		}
		if time.Since(start) >= timeout {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}
	fail(t, fmt.Sprintf("❌ assertion failed. condition did not become true within '%v'", timeout),
		msgAndArgs...)
	return false
}

// False asserts that the given condition is false.
// Returns true if the given condition is false
// Otherwise returns false and flags the test as failed
func False(t testlite.HelperTesting, b bool, msgAndArgs ...any) bool {
	if b {
		t.Helper()
		fail(t, "❌ assertion failed. expected false, but got true\n", msgAndArgs...)
		return false
	}
	return true
}

// IsEmpty asserts the slice isn't nil and doesn't contain any element
func IsEmpty[S ~[]E, E any](t testlite.HelperTesting, s S, msgAndArgs ...any) bool {
	size := len(s)
	if size != 0 {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. expected slice to be empty, but it has '%d' element(s)", size),
			msgAndArgs...)
		return false
	}
	return true
}

// IsErrorOf asserts that the error is not nil and any error in error's tree matches
// the target.
func IsErrorOf(t testlite.HelperTesting, err error, expected error, msgAndArgs ...any) bool {
	t.Helper()
	if err == nil {
		fail(t, fmt.Sprintf("❌ assertion failed. expected error of '%v', but got nil\n", expected), msgAndArgs...)
	}
	if !errors.Is(err, expected) {
		fail(t, fmt.Sprintf("❌ assertion failed. expected error to be '%v', but got '%v'\n", expected, err), msgAndArgs...)
		return false
	}
	return true
}

// IsErrorMessage asserts that the error is not nil and it has the same message of the target
func IsErrorMessage(t testlite.HelperTesting, err error, expected string) bool {
	t.Helper()
	if err == nil || err.Error() != expected {
		fail(t, "expected error message '%s', got '%v'", expected, err)
		return false
	}
	return true
}

// IsNil asserts the value is nil or flag the test as failed.
func IsNil(t testlite.HelperTesting, value any, msgAndArgs ...any) bool {
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
func IsTypeOf[T any](t testlite.HelperTesting, v any, msgAndArgs ...any) bool {
	if _, ok := v.(T); !ok {
		t.Helper()
		typx := reflect.TypeOf((*T)(nil)).Elem()
		fail(t, fmt.Sprintf("❌ assertion failed. expected data to be of type '%s',"+
			" but got '%T'\n", typx, v), msgAndArgs...)
		return false
	}
	return true
}

// Len asserts the slice length equals the expected length.
func Len[S ~[]E, E any](t testlite.HelperTesting, s S, length int, msgAndArgs ...any) bool {
	slen := len(s)
	if slen != length {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. expected length to be '%d', but got '%d.'", length, slen),
			msgAndArgs...)
		return false
	}
	return true
}

// NoError asserts that the given error is nil or flag the test as failed.
func NoError(t testlite.HelperTesting, err error, msgAndArgs ...any) bool {
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
func NoNilFields(t testlite.HelperTesting, checkUnexported bool, target any, msgAndArgs ...any) bool {
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
func NoNilNonInterfaceFields(t testlite.HelperTesting, checkUnexported bool, target any, msgAndArgs ...any) bool {
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
func NotEmpty(t testlite.HelperTesting, v string, msgAndArgs ...any) bool {
	if v == "" {
		t.Helper()
		fail(t, "❌ assertion failed. string is empty", msgAndArgs...)
		return false
	}
	return true
}

// NotEquals asserts the given values are not equals or flag the test as failed.
func NotEquals(t testlite.HelperTesting, a, b any, msgAndArgs ...any) bool {
	if ask.AreEquals(a, b) {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. values are equals: a = <%T>%v, "+
			"b = <%T>%v\n", a, a, b, b), msgAndArgs...)
		return false
	}
	return true
}

// NotNil asserts that the given value is not nil or flag the test as failed.
func NotNil(t testlite.HelperTesting, value any, msgAndArgs ...any) bool {
	if value == nil || ask.IsNil(reflect.ValueOf(value)) {
		t.Helper()
		fail(t, "❌ assertion failed. target is nil", msgAndArgs...)
		return false
	}
	return true
}

// NoPanics asserts that the function doesn't panic or flag the test as failed
func NoPanics(t testlite.HelperTesting, f func(), msgAndArgs ...any) (res bool) {
	t.Helper()
	defer noPanic(t, &res, msgAndArgs...)
	f()
	return
}

func noPanic(t testlite.HelperTesting, res *bool, msgAndArgs ...any) {
	t.Helper()
	o := true
	if r := recover(); r != nil {
		fail(t, fmt.Sprintf("❌ assertion failed. expected no panic, but got 'panic: %v'\n", r), msgAndArgs...)
		o = false
	}
	res = &o
}

// NotSame asserts that two references aren't pointing to the same object.
// It'll never return true for non-reference types
func NotSame(t testlite.HelperTesting, a, b any, msgAndArgs ...any) bool {
	same, addr, _ := ask.AreSameInstances(a, b)
	if same {
		t.Helper()
		fail(t, fmt.Sprintf("❌ assertion failed. expected data to be different instances,"+
			" but got same address: a = <%#x>, b = <%#x>\n", addr[0], addr[1]), msgAndArgs...)
		return false
	}
	return true
}

// NotTypeOf asserts that the value is not of the given type or flag the test as failed.
func NotTypeOf[T any](t testlite.HelperTesting, v any, msgAndArgs ...any) bool {
	if _, ok := v.(T); ok {
		t.Helper()
		typx := reflect.TypeOf((*T)(nil)).Elem()
		fail(t, fmt.Sprintf("❌ assertion failed. expected data not to be of "+
			"type '%s', but got '%T'\n", typx, v), msgAndArgs...)
		return false
	}
	return true
}

// Panics asserts that the function panics or flag the test as failed
func Panics(t testlite.HelperTesting, f func(), msgAndArgs ...any) (res bool) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			res = true
		}
	}()
	f()

	if !res {
		fail(t, "❌ assertion failed. expected a panic, but got none\n", msgAndArgs...)
		return false
	}
	return
}

// Same asserts that two references are pointing to the same object.
// It'll never return true for non-reference types
func Same(t testlite.HelperTesting, a, b any, msgAndArgs ...any) bool {
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
func True(t testlite.HelperTesting, condition bool, msgAndArgs ...any) bool {
	if !condition {
		t.Helper()
		fail(t, "❌ assertion failed. expected true, but got false\n", msgAndArgs...)
		return false
	}
	return true
}
