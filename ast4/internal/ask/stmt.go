// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ask

import (
	"go/ast"
)

func IsSimilarStmt(a, b ast.Stmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	switch stmt := a.(type) {
	case *ast.AssignStmt:
		other, ok := b.(*ast.AssignStmt)
		if !ok {
			return false
		}
		return IsSimilarAssignStmt(stmt, other)
	case *ast.BadStmt:
		_, ok := b.(*ast.BadStmt)
		return ok
	case *ast.BlockStmt:
		other, ok := b.(*ast.BlockStmt)
		if !ok {
			return false
		}
		return IsSimilarBlockStmt(stmt, other)
	case *ast.BranchStmt:
		other, ok := b.(*ast.BranchStmt)
		if !ok {
			return false
		}
		return IsSimilarBranchStmt(stmt, other)
	case *ast.CaseClause:
		other, ok := b.(*ast.CaseClause)
		if !ok {
			return false
		}
		return IsSimilarCaseClause(stmt, other)
	case *ast.CommClause:
		other, ok := b.(*ast.CommClause)
		if !ok {
			return false
		}
		return IsSimilarCommClause(stmt, other)
	case *ast.DeclStmt:
		other, ok := b.(*ast.DeclStmt)
		if !ok {
			return false
		}
		return IsSimilarDeclStmt(stmt, other)
	case *ast.DeferStmt:
		other, ok := b.(*ast.DeferStmt)
		if !ok {
			return false
		}
		return IsSimilarDeferStmt(stmt, other)
	case *ast.EmptyStmt:
		other, ok := b.(*ast.EmptyStmt)
		if !ok {
			return false
		}
		return IsSimilarEmptyStmt(stmt, other)
	case *ast.ExprStmt:
		other, ok := b.(*ast.ExprStmt)
		if !ok {
			return false
		}
		return IsSimilarExprStmt(stmt, other)
	case *ast.ForStmt:
		other, ok := b.(*ast.ForStmt)
		if !ok {
			return false
		}
		return IsSimilarForStmt(stmt, other)
	case *ast.GoStmt:
		other, ok := b.(*ast.GoStmt)
		if !ok {
			return false
		}
		return IsSimilarGoStmt(stmt, other)
	case *ast.IfStmt:
		other, ok := b.(*ast.IfStmt)
		if !ok {
			return false
		}
		return IsSimilarIfStmt(stmt, other)
	case *ast.IncDecStmt:
		other, ok := b.(*ast.IncDecStmt)
		if !ok {
			return false
		}
		return IsSimilarIncDecStmt(stmt, other)
	case *ast.LabeledStmt:
		other, ok := b.(*ast.LabeledStmt)
		if !ok {
			return false
		}
		return IsSimilarLabeledStmt(stmt, other)
	case *ast.RangeStmt:
		other, ok := b.(*ast.RangeStmt)
		if !ok {
			return false
		}
		return IsSimilarRangeStmt(stmt, other)
	case *ast.ReturnStmt:
		other, ok := b.(*ast.ReturnStmt)
		if !ok {
			return false
		}
		return IsSimilarReturnStmt(stmt, other)
	case *ast.SelectStmt:
		other, ok := b.(*ast.SelectStmt)
		if !ok {
			return false
		}
		return IsSimilarSelectStmt(stmt, other)
	case *ast.SendStmt:
		other, ok := b.(*ast.SendStmt)
		if !ok {
			return false
		}
		return IsSimilarSendStmt(stmt, other)
	case *ast.SwitchStmt:
		other, ok := b.(*ast.SwitchStmt)
		if !ok {
			return false
		}
		return IsSimilarSwitchStmt(stmt, other)
	case *ast.TypeSwitchStmt:
		other, ok := b.(*ast.TypeSwitchStmt)
		if !ok {
			return false
		}
		return IsSimilarTypeSwitchStmt(stmt, other)
	default:
		return false
	}
}

func IsSimilarAssignStmt(a, b *ast.AssignStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	if len(a.Lhs) != len(b.Lhs) {
		return false
	}
	for i := range a.Lhs {
		if !IsSimilarExpr(a.Lhs[i], b.Lhs[i]) {
			return false
		}
	}
	if len(a.Rhs) != len(b.Rhs) {
		return false
	}
	for i := range a.Rhs {
		if !IsSimilarExpr(a.Rhs[i], b.Rhs[i]) {
			return false
		}
	}
	return true
}

func IsSimilarBlockStmt(a, b *ast.BlockStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if len(a.List) != len(b.List) {
		return false
	}
	for i := range a.List {
		if !IsSimilarStmt(a.List[i], b.List[i]) {
			return false
		}
	}
	return true
}

func IsSimilarBranchStmt(a, b *ast.BranchStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return a.Tok == b.Tok && a.Label.Name == b.Label.Name
}

func IsSimilarCaseClause(a, b *ast.CaseClause) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	if len(a.List) != len(b.List) {
		return false
	}
	for i := range a.List {
		if !IsSimilarExpr(a.List[i], b.List[i]) {
			return false
		}
	}
	if len(a.Body) != len(b.Body) {
		return false
	}
	for i := range a.Body {
		if !IsSimilarStmt(a.Body[i], b.Body[i]) {
			return false
		}
	}
	return true
}

func IsSimilarCommClause(a, b *ast.CommClause) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	if len(a.Body) != len(b.Body) {
		return false
	}
	for i := range a.Body {
		if !IsSimilarStmt(a.Body[i], b.Body[i]) {
			return false
		}
	}
	return IsSimilarStmt(a.Comm, b.Comm)
}

func IsSimilarDeclStmt(a, b *ast.DeclStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return IsSimilarDecl(a.Decl, b.Decl)
}

func IsSimilarDeferStmt(a, b *ast.DeferStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return IsSimilarCallExpr(a.Call, b.Call)
}

func IsSimilarEmptyStmt(a, b *ast.EmptyStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return a.Implicit == b.Implicit
}

func IsSimilarExprStmt(a, b *ast.ExprStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return IsSimilarExpr(a.X, b.X)
}

func IsSimilarForStmt(a, b *ast.ForStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return IsSimilarStmt(a.Init, b.Init) &&
		IsSimilarExpr(a.Cond, b.Cond) &&
		IsSimilarStmt(a.Post, b.Post) &&
		IsSimilarBlockStmt(a.Body, b.Body)
}

func IsSimilarGoStmt(a, b *ast.GoStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return IsSimilarCallExpr(a.Call, b.Call)
}

func IsSimilarIfStmt(a, b *ast.IfStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return IsSimilarStmt(a.Init, b.Init) &&
		IsSimilarExpr(a.Cond, b.Cond) &&
		IsSimilarBlockStmt(a.Body, b.Body) &&
		IsSimilarStmt(a.Else, b.Else)
}

func IsSimilarIncDecStmt(a, b *ast.IncDecStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return a.Tok == b.Tok && IsSimilarExpr(a.X, b.X)
}

func IsSimilarLabeledStmt(a, b *ast.LabeledStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return a.Label.Name == b.Label.Name && IsSimilarStmt(a.Stmt, b.Stmt)
}

func IsSimilarRangeStmt(a, b *ast.RangeStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return IsSimilarExpr(a.Key, b.Key) &&
		IsSimilarExpr(a.Value, b.Value) &&
		a.Tok == b.Tok &&
		IsSimilarExpr(a.X, b.X) &&
		IsSimilarBlockStmt(a.Body, b.Body)
}

func IsSimilarReturnStmt(a, b *ast.ReturnStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	if len(a.Results) != len(b.Results) {
		return false
	}
	for i := range a.Results {
		if !IsSimilarExpr(a.Results[i], b.Results[i]) {
			return false
		}
	}
	return true
}

func IsSimilarSelectStmt(a, b *ast.SelectStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return IsSimilarBlockStmt(a.Body, b.Body)
}

func IsSimilarSendStmt(a, b *ast.SendStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return IsSimilarExpr(a.Chan, b.Chan) &&
		IsSimilarExpr(a.Value, b.Value)
}

func IsSimilarSwitchStmt(a, b *ast.SwitchStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return IsSimilarStmt(a.Init, b.Init) &&
		IsSimilarExpr(a.Tag, b.Tag) &&
		IsSimilarBlockStmt(a.Body, b.Body)
}

func IsSimilarTypeSwitchStmt(a, b *ast.TypeSwitchStmt) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return IsSimilarStmt(a.Init, b.Init) &&
		IsSimilarStmt(a.Assign, b.Assign) &&
		IsSimilarBlockStmt(a.Body, b.Body)
}

func IsStmt(stmt ast.Node) bool {
	switch stmt.(type) {
	case *ast.AssignStmt, *ast.BadStmt, *ast.BlockStmt, *ast.BranchStmt, *ast.CaseClause,
		*ast.CommClause, *ast.DeclStmt, *ast.DeferStmt, *ast.EmptyStmt, *ast.ExprStmt,
		*ast.ForStmt, *ast.GoStmt, *ast.IfStmt, *ast.IncDecStmt, *ast.LabeledStmt,
		*ast.RangeStmt, *ast.ReturnStmt, *ast.SelectStmt, *ast.SendStmt, *ast.SwitchStmt,
		*ast.TypeSwitchStmt:
		return true
	default:
		return false
	}
}

func IsAssignStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.AssignStmt)
	return ok
}

func IsBadStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.BadStmt)
	return ok
}

func IsBlockStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.BlockStmt)
	return ok
}

func IsBranchStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.BranchStmt)
	return ok
}

func IsCaseClause(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.CaseClause)
	return ok
}

func IsCommClause(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.CommClause)
	return ok
}

func IsDeclStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.DeclStmt)
	return ok
}

func IsDeferStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.DeferStmt)
	return ok
}

func IsEmptyStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.EmptyStmt)
	return ok
}

func IsExprStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.ExprStmt)
	return ok
}

func IsForStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.ForStmt)
	return ok
}

func IsGoStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.GoStmt)
	return ok
}

func IsIfStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.IfStmt)
	return ok
}

func IsIncDecStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.IncDecStmt)
	return ok
}

func IsLabeledStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.LabeledStmt)
	return ok
}

func IsRangeStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.RangeStmt)
	return ok
}

func IsReturnStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.ReturnStmt)
	return ok
}

func IsSelectStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.SelectStmt)
	return ok
}

func IsSendStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.SendStmt)
	return ok
}

func IsSwitchStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.SwitchStmt)
	return ok
}

func IsTypeSwitchStmt(stmt ast.Stmt) bool {
	_, ok := stmt.(*ast.TypeSwitchStmt)
	return ok
}
