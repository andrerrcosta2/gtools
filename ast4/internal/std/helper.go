// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package std

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
)

var Zero = indent.Zero()

func DeclTypeError(tab indent.Indentor, decl ast.Decl) string {
	return sprints.Errorf(tab, "unsupported decl type: '%T'", decl)
}

func DeclTypeDiffError(tab indent.Indentor, a, b ast.Decl) string {
	return sprints.Errorf(tab, "unsupported decl type4: a = '%T', b = '%T'", a, b)
}

func ExprTypeError(tab indent.Indentor, expr ast.Expr) string {
	return sprints.Errorf(tab, "unsupported expression type: '%T'", expr)
}

func ExprTypeDiffError(tab indent.Indentor, a, b ast.Expr) string {
	return sprints.Errorf(tab, "unsupported expression typse: a = '%T', b = '%T'", a, b)
}

func NodeTypeError(tab indent.Indentor, node ast.Node) string {
	return sprints.Errorf(tab, "unsupported node type: '%T'", node)
}

func NodeTypeDiffError(tab indent.Indentor, a, b ast.Node) string {
	return sprints.Errorf(tab, "unsupported node type4: a = '%T', b = '%T'", a, b)
}

func SpecTypeError(tab indent.Indentor, spec ast.Spec) string {
	return sprints.Errorf(tab, "unsupported spec type: '%T'", spec)
}

func SpecTypeDiffError(tab indent.Indentor, a, b ast.Spec) string {
	return sprints.Errorf(tab, "unsupported spec type4: a = '%T', b = '%T'", a, b)
}

func StmtTypeError(tab indent.Indentor, stmt ast.Stmt) string {
	return sprints.Errorf(tab, "unsupported statement type: '%T'", stmt)
}

func StmtTypeDiffError(tab indent.Indentor, a, b ast.Stmt) string {
	return sprints.Errorf(tab, "unsupported statement type4: a = '%T', b = '%T'", a, b)
}
