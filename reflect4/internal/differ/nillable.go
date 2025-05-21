// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/reflect4/internal/reflect4"
	"github.com/andrerrcosta2/gtools/reflect4/internal/types"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"reflect"
)

// Nullables returns the difference between two values only on terms of nullability
//
// # Important
//   - This function does not compare the value types
func Nullables(tab indent.Tab, a, b reflect.Value) (msg string, hasNil bool, equals bool, err error) {
	// Ensure both values are valid
	if !a.IsValid() || !b.IsValid() {
		return reflect4.ErrInvalidValue.Error(), false, false, reflect4.ErrInvalidValue
	}

	// Check if both values can be nil
	canANil := values.CanNil(a)
	canBNil := values.CanNil(b)

	// If one value can be nil and the other cannot, they are of different types
	if canANil != canBNil {
		if canANil && a.IsNil() {
			bb := types.Name(b.Type())
			return differs.NilReceived(tab, bb, a.Kind().String()), true, false, nil
		}
		aa := types.Name(a.Type())
		return differs.NilExpected(tab, aa, b.Kind().String()), true, false, nil
	}

	// If neither value can be nil, return empty strings
	if !canANil && !canBNil {
		return "", false, false, nil
	}

	// Check if both values are nil
	if a.IsNil() && b.IsNil() {
		aa, bb := types.Name(a.Type()), types.Name(b.Type())
		return differs.BothNils(tab, a.Kind().String(), aa, bb), true,
			aa == bb, nil
	}

	// Check if one of the values is nil
	if a.IsNil() || b.IsNil() {
		if a.IsNil() {
			bb := types.Name(b.Type())
			return differs.NilReceived(tab, bb, a.Kind().String()), true, false, nil
		}
		aa := types.Name(a.Type())
		return differs.NilExpected(tab, aa, b.Kind().String()), true, false, nil

	}

	// Neither value is nil
	return "", false, false, nil
}

func nullables(a, b reflect.Value) (msg string, nils, equals bool) {
	// Check if both values are nil
	if a.IsNil() && b.IsNil() {
		aa, bb := types.Name(a.Type()), types.Name(b.Type())
		return differs.BothNils(indent.Zero(), a.Kind().String(), aa, bb), true, aa == bb
	}

	// Check if one of the values is nil
	if a.IsNil() || b.IsNil() {
		if a.IsNil() {
			bb := types.Name(b.Type())
			return differs.NilReceived(indent.Zero(), b.Kind().String(), bb), true, false
		}
		aa := types.Name(a.Type())
		return differs.NilExpected(indent.Zero(), a.Kind().String(), aa), true, false
	}

	// No nil values
	return "", false, false
}

func invalids(a, b reflect.Value) (msg string, invalids, equals bool) {
	if !a.IsValid() || !b.IsValid() {
		if !a.IsValid() && !b.IsValid() {
			return differs.BothInvalid(indent.Zero()), true, true
		}
		if a.IsValid() {
			return differs.InvalidExpected(indent.Zero(), types.Name(a.Type())), true, false
		}
		return differs.InvalidReceived(indent.Zero(), types.Name(b.Type())), true, false
	}
	return differs.Empty(indent.Zero()), false, false
}
