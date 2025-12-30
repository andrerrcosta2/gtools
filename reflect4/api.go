// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package reflect4

import (
	"reflect"

	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/reflect4/internal/differ"
	"github.com/andrerrcosta2/gtools/reflect4/internal/equals"
	"github.com/andrerrcosta2/gtools/reflect4/internal/handlers/data"
	"github.com/andrerrcosta2/gtools/reflect4/internal/sprint"
)

// DeepCopy creates a clone of the value pointed to by `t`.
//
// If `deep` is true, it performs a deep clone—recursively cloning all nested
// values such as structs, slices, maps, and interfaces.
// This ensures the returned clone is completely independent of the original, including all referenced data.
//
// If `deep` is false, it performs a shallow clone—only copying the immediate value,
// sharing references to nested structures like slices, maps, and ptrs.
//
// DeepCopy handles cyclic references safely and preserves the structure of complex
// data graphs.
// It supports all standard Go types, with some limitations:
//
//   - channels: not deep-copied.
//     The same channel reference is reused in the clone.
//     You may implement a custom strategy to replace or close them if needed.
//
//   - functions: cannot be copied.
//     Function values are preserved as-is, but will
//     always compare as unequal across copies, even if they have identical logic.
//
//   - unsafe.Pointers: not traversed or copied, for safety reasons.
//
// Because of these limitations, values containing channels or functions will not
// compare as deep-compare - `reflect.DeepDiffer` returns false - even if logically similar.
//
// Requirements:
//   - The input must be a pointer (`*T`).
//
// Returns:
//   - A copied value of type `T`.
//   - An error if the clone op fails - due to unhandled types or
//     internal reflection issues - or if the result cannot be cast back to `T`.
//
// Example:
//
//	original := &MyStruct{Name: "Alice", Tags: []string{"go", "dev"}}
//	copied, err := DeepCopy(original, true) // deep clone
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// `copied` is now a fully independent clone
func DeepCopy[T any](t *T, o ...Option) (cp T, err error) {
	var val reflect.Value
	var ok bool
	val, err = data.DeepCopy(reflect.ValueOf(t), o...)
	if err != nil {
		return
	}
	cp, ok = val.Interface().(T)
	if !ok {
		return cp, fmx.Errorf("unable to cast value to %T", t)
	}
	return
}

// DeepDiffer compares two arbitrary Go values for **logical equivalence**, offering
// more flexible semantics than `reflect.DeepDiffer`.
// It is designed for use cases like testing, snapshotting, caching validation, and configuration diffing.
//
// ### Key Differences from reflect.DeepDiffer:
//
// - **channels**: Two channels are considered compare if:
//   - They have the same type - including element type.
//   - They have the same buffer capacity.
//   - Internal state - open/closed, data in buffer - and identity are ignored.
//     Example:
//     ch1 := make(chan int, 10)
//     ch2 := make(chan int, 10)
//     DeepDiffer(ch1, ch2) // → true
//
// - **functions**: Considered compare if they have the same **signature**
//   - parameter and return types -, regardless of closure context or implementation.
//     Example:
//     fn1 := func(x int) int { return x + 1 }
//     fn2 := func(x int) int { return x + 100 }
//     DeepDiffer(fn1, fn2) // → true
//
// - **Pointers**: pointer identity is ignored.
// Deep structural comparison is applied
//
//		to the dereferenced values.
//
//	  - **Cyclic References**: handled safely with cycle detection.
//	    Infinite recursion is avoided.
//
//	  - **Unexported Fields**: comparison of unexported fields is controlled via `read.Opt`.
//	    Some options allow skipping or inspecting them depending on need.
//
// ### Options:
// You may pass flags from the `read.Opt` type to:
//   - Ignore specific fields
//   - Enable detailed logging
//   - Toggle comparison rules
//
// See the `read` package for available flags.
//
// ### Returns:
//   - `bool`: `true` if values are deeply logically compare.
//   - `[]string`: a slice of string paths representing mismatches -`.User.Friends[1].ID`.
//
// ### Notes:
// This function prioritizes **logical equivalence** over memory identity.
// It should
// not be used in contexts where reference aliasing, pointer identity, or channel
// state is semantically significant.
//
// ### Example:
//
//	a := map[string]any{
//	    "Name": "Alice",
//	    "OnUpdate": func(string) {},            // any func(string)
//	    "LogCh":    make(chan string, 5),       // same capacity
//	}
//	b := map[string]any{
//	    "Name": "Alice",
//	    "OnUpdate": func(s string) { println(s) },
//	    "LogCh":    make(chan string, 5),
//	}
//
//	compare, diffs := DeepDiffer(a, b)
//	// compare = true, diffs = nil
func DeepDiffer(a, b any, opts ...Option) (diff string, equals bool) {
	return differ.Between(reflect.ValueOf(a), reflect.ValueOf(b), opts...)
}

// DeepEqual reports whether a and b are “deeply equal,” defined by the strategies
// represented by the op.Option
func DeepEqual(a, b any, opts ...Option) bool {
	if a == nil || b == nil {
		return a == b
	}
	return equals.Deep(reflect.ValueOf(a), reflect.ValueOf(b), opts...)
}

func Sprint(a any, opts ...Option) string {
	return sprint.Of(indent.Zero(), reflect.ValueOf(a), opts...)
}
