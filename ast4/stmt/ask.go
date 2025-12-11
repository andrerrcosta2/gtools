// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package stmt

import (
	"github.com/andrerrcosta2/gtools/ast4/internal/ask"

	"go/ast"
)

func IsSimilar(a, b ast.Stmt) bool {
	return ask.IsSimilarStmt(a, b)
}

func IsSimilarAssignStmt(a, b *ast.AssignStmt) bool {
	return ask.IsSimilarAssignStmt(a, b)
}

func IsSimilarBlockStmt(a, b *ast.BlockStmt) bool {
	return ask.IsSimilarBlockStmt(a, b)
}

func IsSimilarBranchStmt(a, b *ast.BranchStmt) bool {
	return ask.IsSimilarBranchStmt(a, b)
}

func IsSimilarCaseClause(a, b *ast.CaseClause) bool {
	return ask.IsSimilarCaseClause(a, b)
}

func IsSimilarCommClause(a, b *ast.CommClause) bool {
	return ask.IsSimilarCommClause(a, b)
}

func IsSimilarDeclStmt(a, b *ast.DeclStmt) bool {
	return ask.IsSimilarDeclStmt(a, b)
}

func IsSimilarDeferStmt(a, b *ast.DeferStmt) bool {
	return ask.IsSimilarDeferStmt(a, b)
}

func IsSimilarEmptyStmt(a, b *ast.EmptyStmt) bool {
	return ask.IsSimilarEmptyStmt(a, b)
}

func IsSimilarExprStmt(a, b *ast.ExprStmt) bool {
	return ask.IsSimilarExprStmt(a, b)
}

func IsSimilarForStmt(a, b *ast.ForStmt) bool {
	return ask.IsSimilarForStmt(a, b)
}

func IsSimilarGoStmt(a, b *ast.GoStmt) bool {
	return ask.IsSimilarGoStmt(a, b)
}

func IsSimilarIfStmt(a, b *ast.IfStmt) bool {
	return ask.IsSimilarIfStmt(a, b)
}

func IsSimilarIncDecStmt(a, b *ast.IncDecStmt) bool {
	return ask.IsSimilarIncDecStmt(a, b)
}

func IsSimilarLabeledStmt(a, b *ast.LabeledStmt) bool {
	return ask.IsSimilarLabeledStmt(a, b)
}

func IsSimilarRangeStmt(a, b *ast.RangeStmt) bool {
	return ask.IsSimilarRangeStmt(a, b)
}

func IsSimilarReturnStmt(a, b *ast.ReturnStmt) bool {
	return ask.IsSimilarReturnStmt(a, b)
}

func IsSimilarSelectStmt(a, b *ast.SelectStmt) bool {
	return ask.IsSimilarSelectStmt(a, b)
}

func IsSimilarSendStmt(a, b *ast.SendStmt) bool {
	return ask.IsSimilarSendStmt(a, b)
}

func IsSimilarSwitchStmt(a, b *ast.SwitchStmt) bool {
	return ask.IsSimilarSwitchStmt(a, b)
}

func IsSimilarTypeSwitchStmt(a, b *ast.TypeSwitchStmt) bool {
	return ask.IsSimilarTypeSwitchStmt(a, b)
}

func IsStmt(stmt ast.Node) bool {
	return ask.IsStmt(stmt)
}

func IsAssignStmt(stmt ast.Stmt) bool {
	return ask.IsAssignStmt(stmt)
}

func IsBadStmt(stmt ast.Stmt) bool {
	return ask.IsBadStmt(stmt)
}

func IsBlockStmt(stmt ast.Stmt) bool {
	return ask.IsBlockStmt(stmt)
}

func IsBranchStmt(stmt ast.Stmt) bool {
	return ask.IsBranchStmt(stmt)
}

func IsCaseClause(stmt ast.Stmt) bool {
	return ask.IsCaseClause(stmt)
}

func IsCommClause(stmt ast.Stmt) bool {
	return ask.IsCommClause(stmt)
}

func IsDeclStmt(stmt ast.Stmt) bool {
	return ask.IsDeclStmt(stmt)
}

func IsDeferStmt(stmt ast.Stmt) bool {
	return ask.IsDeferStmt(stmt)
}

func IsEmptyStmt(stmt ast.Stmt) bool {
	return ask.IsEmptyStmt(stmt)
}

func IsExprStmt(stmt ast.Stmt) bool {
	return ask.IsExprStmt(stmt)
}

func IsForStmt(stmt ast.Stmt) bool {
	return ask.IsForStmt(stmt)
}

func IsGoStmt(stmt ast.Stmt) bool {
	return ask.IsGoStmt(stmt)
}

func IsIfStmt(stmt ast.Stmt) bool {
	return ask.IsIfStmt(stmt)
}

func IsIncDecStmt(stmt ast.Stmt) bool {
	return ask.IsIncDecStmt(stmt)
}

func IsLabeledStmt(stmt ast.Stmt) bool {
	return ask.IsLabeledStmt(stmt)
}

func IsRangeStmt(stmt ast.Stmt) bool {
	return ask.IsRangeStmt(stmt)
}

func IsReturnStmt(stmt ast.Stmt) bool {
	return ask.IsReturnStmt(stmt)
}

func IsSelectStmt(stmt ast.Stmt) bool {
	return ask.IsSelectStmt(stmt)
}

func IsSendStmt(stmt ast.Stmt) bool {
	return ask.IsSendStmt(stmt)
}

func IsSwitchStmt(stmt ast.Stmt) bool {
	return ask.IsSwitchStmt(stmt)
}

func IsTypeSwitchStmt(stmt ast.Stmt) bool {
	return ask.IsTypeSwitchStmt(stmt)
}
