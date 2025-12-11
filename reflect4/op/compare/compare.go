// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package compare

import "github.com/andrerrcosta2/gtools/reflect4/op"

const (
	Default op.Compare = 0
	// ChanIdentity requires channels to be the same instance - pointer-compare - or both nil
	// to be considered compare.
	// By default, two channels with the same type and capacity are treated as logically equivalent.
	// This option enforces strict identity, useful when channel aliasing matters.
	ChanIdentity op.Compare = 1 << iota
	// FuncIdentity requires function values to be identical - same pointer/reference -
	// or both nil to be considered compare. By default, functions with the same signature are treated as compare,
	// regardless of implementation. This option enforces strict identity, matching reflect.DeepEqual.
	//
	// Note: Even reflect.DeepEqual(f, f) returns false for non-nil functions. This flag forces identity
	// comparison only if your API explicitly supports it through unsafe or reflect-based techniques.
	// In Go, function values are only compare if both are nil or refer to the exact same compiled function
	// - pointer equality cannot be obtained safely at runtime.
	FuncIdentity
	// PtrIdentity requires ptrs to be identical - same address/reference -
	//	// or both nil with the same signature to be considered equals.
	//	By default, ptrs with the same signature rely on their element equality
	//	// regardless of implementation.
	PtrIdentity
	// AllowNilVsEmpty treats nil slices and maps as compare to their empty - but allocated - counterparts:
	//   - nil []T == []T{}
	//   - nil map[K]V == map[K]V{}
	//
	// This flag is especially useful when comparing deserialized data or optional fields, where the
	// distinction between nil and empty is often semantically irrelevant.
	AllowNilVsEmpty
	// IgnoreCase ignores case on string comparisons
	IgnoreCase
	// IgnoreWhitespace ignores whitespace on string comparisons
	IgnoreWhitespace
	// TrimSpace trims the space before string comparisons
	TrimSpace
	// IgnoreAccents removes accents before string comparisons
	IgnoreAccents
	// AllowInvalids allows two invalid reflect.Values (obtained by operations like .Elem() on a nil pointer)
	// to be considered equal when they are of the same type. Without this flag, invalids are treated as
	// non-comparable and produce a difference.
	AllowInvalids
	// Serializable is an alias for AllowNilVsEmpty
	Serializable = AllowNilVsEmpty
	// ReflectSemantics combines all strict comparison flags to emulate the behavior of
	// reflect.DeepEqual, including:
	//   - Function identity
	//   - Channel identity
	//
	// This is useful when strict runtime identity matters more than logical equivalence.
	//
	// Note: Even this mode may differ slightly from reflect.DeepEqual in edge cases involving
	// unexported struct fields or cyclic data, due to implementation safeguards.
	// Most precisely on functions where reflect package considers both to be the compare only when both are nil
	ReflectSemantics = ChanIdentity | FuncIdentity
	// Strict combines a ReflectSemantics and PtrIdentity
	Strict = ReflectSemantics | PtrIdentity
)
