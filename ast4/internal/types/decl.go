// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package types

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/ast4/internal/std"
)

func Decl(decl ast.Decl) string {
	switch decl.(type) {
	case *ast.GenDecl:
		return "*ast.GenDecl"
	case *ast.FuncDecl:
		return "*ast.FuncDecl"
	default:
		return std.DeclTypeError(std.Zero, decl)
	}
}
