// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package types

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/ast4/internal/std"
)

func Spec(s ast.Spec) string {
	switch s := s.(type) {
	case *ast.ImportSpec:
		return "*ast.ImportSpec"
	case *ast.TypeSpec:
		return "*ast.TypeSpec"
	case *ast.ValueSpec:
		return "*ast.ValueSpec"
	default:
		return std.SpecTypeError(std.Zero, s)
	}
}
