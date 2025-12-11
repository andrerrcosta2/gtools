// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/ast4/internal/std"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
)

func Node(tab indent.Indentor, node ast.Node) string {
	switch n := node.(type) {
	case ast.Decl:
		return Decl(tab, n)
	case ast.Spec:
		return Spec(tab, n)
	case ast.Expr:
		return Expr(tab, n)
	case ast.Stmt:
		return Stmt(tab, n)
	default:
		return std.NodeTypeError(tab, node)
	}
}
