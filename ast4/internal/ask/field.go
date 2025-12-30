// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ask

import "go/ast"

func IsSimilarField(a, b *ast.Field) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	// Compare names should be optional; sometimes the type4 package doesn't
	// provide identifiers for output field4.
	if len(a.Names) != len(b.Names) {
		// If one side has no names, treat them as equivalent
		if len(a.Names) != 0 && len(b.Names) != 0 {
			return false
		}
	} else {
		// Compare individual names (if present)
		for i := range a.Names {
			if !IsSimilarIdent(a.Names[i], b.Names[i]) {
				return false
			}
		}
	}

	// Compare type4
	if !IsSimilarExpr(a.Type, b.Type) {
		return false
	}

	// All checks passed
	return true
}

func IsSimilarFieldList(a, b *ast.FieldList) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if len(a.List) != len(b.List) {
		return false
	}
	if len(a.List) != len(b.List) {
		return false
	}
	for i, field := range a.List {
		if !IsSimilarField(field, b.List[i]) {
			return false
		}
	}
	return true
}
