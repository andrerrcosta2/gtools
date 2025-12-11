// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ask

import (
	"go/ast"
)

func IsSimilarDecl(a, b ast.Decl) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	switch decl := a.(type) {
	case *ast.GenDecl:
		other, ok := b.(*ast.GenDecl)
		if !ok {
			return false
		}
		return IsSimilarGenDecl(decl, other)
	case *ast.FuncDecl:
		other, ok := b.(*ast.FuncDecl)
		if !ok {
			return false
		}
		return IsSimilarFuncDecl(decl, other)
	default:
		return false
	}
}

func IsSimilarGenDecl(decl, other *ast.GenDecl) bool {
	if decl == nil || other == nil {
		return decl == nil && other == nil
	}

	if decl.Tok != other.Tok {
		return false
	}
	if len(decl.Specs) != len(other.Specs) {
		return false
	}
	for _, spec := range decl.Specs {
		for _, oSpec := range other.Specs {
			if !spec.IsSimilarSpec(spec, oSpec) {
				return false
			}
		}
	}
	return true
}

func IsSimilarFuncDecl(decl, other *ast.FuncDecl) bool {
	if decl == nil || other == nil {
		return decl == nil && other == nil
	}
	if !expr.IsSimilarIdent(decl.Name, other.Name) {
		return false
	}
	return IsSimilarBlockStmt(decl.Body, other.Body)
}

func IsDecl(decl ast.Node) bool {
	switch decl.(type) {
	case *ast.GenDecl, *ast.FuncDecl:
		return true
	default:
		return false
	}
}

func IsGenDecl(decl ast.Decl) bool {
	_, ok := decl.(*ast.GenDecl)
	return ok
}

func IsFuncDecl(decl ast.Decl) bool {
	_, ok := decl.(*ast.FuncDecl)
	return ok
}
