// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package types

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/ast4/internal/std"
)

func Stmt(stmt ast.Stmt) string {
	switch stmt.(type) {
	case *ast.AssignStmt:
		return "*ast.AssignStmt"
	case *ast.BadStmt:
		return "*ast.BadStmt"
	case *ast.BlockStmt:
		return "*ast.BlockStmt"
	case *ast.BranchStmt:
		return "*ast.BranchStmt"
	case *ast.CaseClause:
		return "*ast.CaseClause"
	case *ast.CommClause:
		return "*ast.CommClause"
	case *ast.DeclStmt:
		return "*ast.DeclStmt"
	case *ast.DeferStmt:
		return "*ast.DeferStmt"
	case *ast.EmptyStmt:
		return "*ast.EmptyStmt"
	case *ast.ExprStmt:
		return "*ast.ExprStmt"
	case *ast.ForStmt:
		return "*ast.ForStmt"
	case *ast.GoStmt:
		return "*ast.GoStmt"
	case *ast.IfStmt:
		return "*ast.IfStmt"
	case *ast.IncDecStmt:
		return "*ast.IncDecStmt"
	case *ast.LabeledStmt:
		return "*ast.LabeledStmt"
	case *ast.RangeStmt:
		return "*ast.RangeStmt"
	case *ast.ReturnStmt:
		return "*ast.ReturnStmt"
	case *ast.SelectStmt:
		return "*ast.SelectStmt"
	case *ast.SendStmt:
		return "*ast.SendStmt"
	case *ast.SwitchStmt:
		return "*ast.SwitchStmt"
	case *ast.TypeSwitchStmt:
		return "*ast.TypeSwitchStmt"
	default:
		return std.StmtTypeError(std.Zero, stmt)
	}
}
