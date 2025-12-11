// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package decl

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/ast4/internal/ask"
)

func IsSimilar(a, b ast.Decl) bool {
	return ask.IsSimilarDecl(a, b)
}

func IsSimilarGenDecl(decl, other *ast.GenDecl) bool {
	return ask.IsSimilarGenDecl(decl, other)
}

func IsSimilarFuncDecl(decl, other *ast.FuncDecl) bool {
	return ask.IsSimilarFuncDecl(decl, other)
}

func IsDecl(decl ast.Node) bool {
	return ask.IsDecl(decl)
}

func IsGenDecl(decl ast.Decl) bool {
	return ask.IsGenDecl(decl)
}

func IsFuncDecl(decl ast.Decl) bool {
	return ask.IsFuncDecl(decl)
}
