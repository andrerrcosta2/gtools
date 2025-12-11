// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ask

import (
	"go/ast"
)

func IsSimilarSpec(a, b ast.Spec) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	switch expr := a.(type) {
	case *ast.ImportSpec:
		other, ok := b.(*ast.ImportSpec)
		if !ok {
			return false
		}
		return IsSimilarImportSpec(expr, other)
	case *ast.TypeSpec:
		other, ok := b.(*ast.TypeSpec)
		if !ok {
			return false
		}
		return IsSimilarTypeSpec(expr, other)
	case *ast.ValueSpec:
		other, ok := b.(*ast.ValueSpec)
		if !ok {
			return false
		}
		return IsSimilarValueSpec(expr, other)
	default:
		return false
	}
}

func IsSimilarImportSpec(a, b *ast.ImportSpec) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return a.Name.Name == b.Name.Name &&
		IsSimilarExpr(a.Path, b.Path)
}

func IsSimilarTypeSpec(a, b *ast.TypeSpec) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return a.Name.Name == b.Name.Name &&
		IsSimilarExpr(a.Type, b.Type) &&
		IsSimilarFieldList(a.TypeParams, b.TypeParams)
}

func IsSimilarValueSpec(a, b *ast.ValueSpec) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if len(a.Names) != len(b.Names) {
		return false
	}
	for i := range a.Names {
		if a.Names[i].Name != b.Names[i].Name {
			return false
		}
	}
	if len(a.Values) != len(b.Values) {
		return false
	}
	for i := range a.Values {
		if !IsSimilarExpr(a.Values[i], b.Values[i]) {
			return false
		}
	}
	return IsSimilarExpr(a.Type, b.Type)
}

func IsSpec(s ast.Node) bool {
	switch s.(type) {
	case *ast.ImportSpec, *ast.TypeSpec, *ast.ValueSpec:
		return true
	default:
		return false
	}
}

func IsImportSpec(s ast.Spec) bool {
	_, ok := s.(*ast.ImportSpec)
	return ok
}

func IsTypeSpec(s ast.Spec) bool {
	_, ok := s.(*ast.TypeSpec)
	return ok
}

func IsValueSpec(s ast.Spec) bool {
	_, ok := s.(*ast.ValueSpec)
	return ok
}
