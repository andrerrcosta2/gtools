// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differs

import (
	"github.com/andrerrcosta2/gtools/ast4/internal/std"
	"github.com/andrerrcosta2/gtools/ast4/internal/types"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/sprints"

	"go/ast"
)

func Stmt(a, b ast.Stmt) (equals bool, diff string) {
	return stmt(std.Zero, a, b)
}

func stmt(t indent.Indentor, a, b ast.Stmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "Stmt", "ast.Stmt"); has {
		return eq, d
	}
	switch this := a.(type) {
	case *ast.AssignStmt:
		if other, ok := b.(*ast.AssignStmt); ok {
			return assignStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.AssignStmt", types.Stmt(b))
	case *ast.BadStmt:
		if other, ok := b.(*ast.BadStmt); ok {
			return badStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.BadStmt", types.Stmt(b))
	case *ast.BlockStmt:
		if other, ok := b.(*ast.BlockStmt); ok {
			return blockStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.BlockStmt", types.Stmt(b))
	case *ast.BranchStmt:
		if other, ok := b.(*ast.BranchStmt); ok {
			return branchStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.BranchStmt", types.Stmt(b))
	case *ast.CaseClause:
		if other, ok := b.(*ast.CaseClause); ok {
			return caseClause(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.CaseClause", types.Stmt(b))
	case *ast.CommClause:
		if other, ok := b.(*ast.CommClause); ok {
			return commClause(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.CommClause", types.Stmt(b))
	case *ast.DeclStmt:
		if other, ok := b.(*ast.DeclStmt); ok {
			return declStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.DeclStmt", types.Stmt(b))
	case *ast.DeferStmt:
		if other, ok := b.(*ast.DeferStmt); ok {
			return deferStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.DeferStmt", types.Stmt(b))
	case *ast.EmptyStmt:
		if other, ok := b.(*ast.EmptyStmt); ok {
			return emptyStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.EmptyStmt", types.Stmt(b))
	case *ast.ExprStmt:
		if other, ok := b.(*ast.ExprStmt); ok {
			return exprStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.ExprStmt", types.Stmt(b))
	case *ast.ForStmt:
		if other, ok := b.(*ast.ForStmt); ok {
			return forStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.ForStmt", types.Stmt(b))
	case *ast.GoStmt:
		if other, ok := b.(*ast.GoStmt); ok {
			return goStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.GoStmt", types.Stmt(b))
	case *ast.IfStmt:
		if other, ok := b.(*ast.IfStmt); ok {
			return ifStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.IfStmt", types.Stmt(b))
	case *ast.IncDecStmt:
		if other, ok := b.(*ast.IncDecStmt); ok {
			return incDecStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.IncDecStmt", types.Stmt(b))
	case *ast.LabeledStmt:
		if other, ok := b.(*ast.LabeledStmt); ok {
			return labeledStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.LabeledStmt", types.Stmt(b))
	case *ast.RangeStmt:
		if other, ok := b.(*ast.RangeStmt); ok {
			return rangeStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.RangeStmt", types.Stmt(b))
	case *ast.ReturnStmt:
		if other, ok := b.(*ast.ReturnStmt); ok {
			return returnStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.ReturnStmt", types.Stmt(b))
	case *ast.SelectStmt:
		if other, ok := b.(*ast.SelectStmt); ok {
			return selectStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.SelectStmt", types.Stmt(b))
	case *ast.SendStmt:
		if other, ok := b.(*ast.SendStmt); ok {
			return sendStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.SendStmt", types.Stmt(b))
	case *ast.SwitchStmt:
		if other, ok := b.(*ast.SwitchStmt); ok {
			return switchStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.SwitchStmt", types.Stmt(b))
	case *ast.TypeSwitchStmt:
		if other, ok := b.(*ast.TypeSwitchStmt); ok {
			return typeSwitchStmt(t.Inc(), this, other)
		}
		return false, differs.TypesMismatch(t, "*ast.TypeSwitchStmt", types.Stmt(b))
	default:
		return false, std.StmtTypeDiffError(t.Inc(), a, b)
	}
}

func AssignStmt(a, b *ast.AssignStmt) (equals bool, diff string) {
	return assignStmt(std.Zero, a, b)
}

func assignStmt(t indent.Indentor, a, b *ast.AssignStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "AssignStmt", "*ast.AssignStmt"); has {
		return eq, d
	}
	if a.TokPos != b.TokPos {
		return false, differs.ValuesOf(t, "TokPos",
			sprints.Field(std.Zero, "TokPos", sprints.Digit(a.TokPos)),
			sprints.Field(std.Zero, "TokPos", sprints.Digit(b.TokPos)),
		)
	}
	if a.Tok != b.Tok {
		return false, differs.ValuesOf(t, "Tok",
			sprints.Field(std.Zero, "Tok", sprints.Digit(a.Tok)),
			sprints.Field(std.Zero, "Tok", sprints.Digit(b.Tok)),
		)
	}
	if len(a.Lhs) != len(b.Lhs) {
		return false, differs.SizesOf(t, "Lhs",
			sprints.Digit(len(a.Lhs)), sprints.Digit(len(b.Lhs)),
		)
	}
	if len(a.Rhs) != len(b.Rhs) {
		return false, differs.SizesOf(t, "Rhs",
			sprints.Digit(len(a.Rhs)), sprints.Digit(len(b.Rhs)),
		)
	}
	for i := 0; i < len(a.Lhs); i++ {
		if eq, dif := expr(t.Inc(), a.Lhs[i], b.Lhs[i]); !eq {
			return false, t.Sprintf("'Lhs[%d]' values mismatch:\n%s", i, dif)
		}
	}
	for i := 0; i < len(a.Rhs); i++ {
		if eq, dif := expr(t.Inc(), a.Rhs[i], b.Rhs[i]); !eq {
			return false, t.Sprintf("'Rhs[%d]' values mismatch:\n%s", i, dif)
		}
	}
	return true, diff
}

func BadStmt(stmt, other *ast.BadStmt) (equals bool, diff string) {
	return badStmt(std.Zero, stmt, other)
}

func badStmt(t indent.Indentor, a, b *ast.BadStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "BadStmt", "*ast.BadStmt"); has {
		return eq, d
	}
	if a.From != b.From {
		return false, differs.ValuesOf(t, "From",
			sprints.Field(std.Zero, "From", sprints.Digit(a.From)),
			sprints.Field(std.Zero, "From", sprints.Digit(b.From)),
		)
	}
	if a.To != b.To {
		return false, differs.ValuesOf(t, "To",
			sprints.Field(std.Zero, "To", sprints.Digit(a.To)),
			sprints.Field(std.Zero, "To", sprints.Digit(b.To)),
		)
	}
	return true, diff
}

func BlockStmt(a, b *ast.BlockStmt) (equals bool, diff string) {
	return blockStmt(std.Zero, a, b)
}

func blockStmt(t indent.Indentor, a, b *ast.BlockStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "BlockStmt", "*ast.BlockStmt"); has {
		return eq, d
	}
	if a.Lbrace != b.Lbrace {
		return false, differs.ValuesOf(t, "Lbrace",
			sprints.Field(std.Zero, "Lbrace", sprints.Digit(a.Lbrace)),
			sprints.Field(std.Zero, "Lbrace", sprints.Digit(b.Lbrace)),
		)
	}
	if a.Rbrace != b.Rbrace {
		return false, differs.ValuesOf(t, "Rbrace",
			sprints.Field(std.Zero, "Rbrace", sprints.Digit(a.Rbrace)),
			sprints.Field(std.Zero, "Rbrace", sprints.Digit(b.Rbrace)),
		)
	}
	if len(a.List) != len(b.List) {
		return false, differs.SizesOf(t, "List",
			sprints.Digit(len(a.List)), sprints.Digit(len(b.List)),
		)
	}
	for i := 0; i < len(a.List); i++ {
		if eq, dif := stmt(t.Inc(), a.List[i], b.List[i]); !eq {
			return false, t.Sprintf("'List[%d]' values mismatch:\n%s", i, dif)
		}
	}
	return true, diff
}

func BranchStmt(a, b *ast.BranchStmt) (equals bool, diff string) {
	return branchStmt(std.Zero, a, b)
}

func branchStmt(t indent.Indentor, a, b *ast.BranchStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "BranchStmt", "*ast.BranchStmt"); has {
		return eq, d
	}
	if a.TokPos != b.TokPos {
		return false, differs.ValuesOf(t, "TokPos",
			sprints.Field(std.Zero, "TokPos", sprints.Digit(a.TokPos)),
			sprints.Field(std.Zero, "TokPos", sprints.Digit(b.TokPos)),
		)
	}
	if a.Tok != b.Tok {
		return false, differs.ValuesOf(t, "Tok",
			sprints.Field(std.Zero, "Tok", sprints.Digit(a.Tok)),
			sprints.Field(std.Zero, "Tok", sprints.Digit(b.Tok)),
		)
	}
	if eq, dif := ident(t.Inc(), a.Label, b.Label); !eq {
		return false, t.Sprintf("'Label' values mismatch:\n%s", dif)
	}
	return true, diff
}

func CaseClause(a, b *ast.CaseClause) (equals bool, diff string) {
	return caseClause(std.Zero, a, b)
}

func caseClause(t indent.Indentor, a, b *ast.CaseClause) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "CaseClause", "*ast.CaseClause"); has {
		return eq, d
	}
	if a.Case != b.Case {
		return false, differs.ValuesOf(t, "Case",
			sprints.Field(std.Zero, "Case", sprints.Digit(a.Case)),
			sprints.Field(std.Zero, "Case", sprints.Digit(b.Case)),
		)
	}
	if a.Colon != b.Colon {
		return false, differs.ValuesOf(t, "Colon",
			sprints.Field(std.Zero, "Colon", sprints.Digit(a.Colon)),
			sprints.Field(std.Zero, "Colon", sprints.Digit(b.Colon)),
		)
	}
	if len(a.List) != len(b.List) {
		return false, differs.SizesOf(t, "List",
			sprints.Digit(len(a.List)), sprints.Digit(len(b.List)),
		)
	}
	for i := 0; i < len(a.List); i++ {
		if eq, dif := expr(t.Inc(), a.List[i], b.List[i]); !eq {
			return false, t.Sprintf("'List[%d]' values mismatch:\n%s", i, dif)
		}
	}
	if len(a.Body) != len(b.Body) {
		return false, differs.SizesOf(t, "Body", sprints.Digit(len(a.Body)), sprints.Digit(len(b.Body)))
	}
	for i := 0; i < len(a.Body); i++ {
		if eq, dif := stmt(t.Inc(), a.Body[i], b.Body[i]); !eq {
			return false, t.Sprintf("'Body[%d]' values mismatch:\n%s", i, dif)
		}
	}
	return true, diff
}

func CommClause(a, b *ast.CommClause) (equals bool, diff string) {
	return commClause(std.Zero, a, b)
}

func commClause(t indent.Indentor, a, b *ast.CommClause) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "CommClause", "*ast.CommClause"); has {
		return eq, d
	}
	if a.Case != b.Case {
		return false, differs.ValuesOf(t, "Case",
			sprints.Field(std.Zero, "Case", sprints.Digit(a.Case)),
			sprints.Field(std.Zero, "Case", sprints.Digit(b.Case)),
		)
	}
	if a.Colon != b.Colon {
		return false, differs.ValuesOf(t, "Colon",
			sprints.Field(std.Zero, "Colon", sprints.Digit(a.Colon)),
			sprints.Field(std.Zero, "Colon", sprints.Digit(b.Colon)),
		)
	}
	if len(a.Body) != len(b.Body) {
		return false, differs.SizesOf(t, "Body",
			sprints.Digit(len(a.Body)), sprints.Digit(len(b.Body)),
		)
	}
	for i := 0; i < len(a.Body); i++ {
		if eq, dif := stmt(t.Inc(), a.Body[i], b.Body[i]); !eq {
			return false, t.Sprintf("'Body[%d]' values mismatch:\n%s", i, dif)
		}
	}
	if eq, dif := stmt(t.Inc(), a.Comm, b.Comm); !eq {
		return false, t.Sprintf("'Comm' values mismatch:\n%s", dif)
	}
	return true, diff
}

func DeclStmt(a, b *ast.DeclStmt) (equals bool, diff string) {
	return declStmt(std.Zero, a, b)
}

func declStmt(t indent.Indentor, a, b *ast.DeclStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "DeclStmt", "*ast.DeclStmt"); has {
		return eq, d
	}
	if eq, dif := decl(t.Inc(), a.Decl, b.Decl); !eq {
		return false, t.Sprintf("'Decl' values mismatch:\n%s", dif)
	}
	return true, diff
}

func DeferStmt(a, b *ast.DeferStmt) (equals bool, diff string) {
	return deferStmt(std.Zero, a, b)
}

func deferStmt(t indent.Indentor, a, b *ast.DeferStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "DeferStmt", "*ast.DeferStmt"); has {
		return eq, d
	}
	if a.Defer != b.Defer {
		return false, differs.ValuesOf(t, "Defer",
			sprints.Field(std.Zero, "Defer", sprints.Digit(a.Defer)),
			sprints.Field(std.Zero, "Defer", sprints.Digit(b.Defer)),
		)
	}
	if eq, dif := callExpr(t.Inc(), a.Call, b.Call); !eq {
		return false, t.Sprintf("'Call' values mismatch:\n%s", dif)
	}
	return true, diff
}

func EmptyStmt(a, b *ast.EmptyStmt) (equals bool, diff string) {
	return emptyStmt(std.Zero, a, b)
}

func emptyStmt(t indent.Indentor, a, b *ast.EmptyStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "EmptyStmt", "*ast.EmptyStmt"); has {
		return eq, d
	}
	if a.Implicit != b.Implicit {
		return false, differs.ValuesOf(t, "Implicit",
			sprints.Field(std.Zero, "Implicit", sprints.Bool(std.Zero, a.Implicit)),
			sprints.Field(std.Zero, "Implicit", sprints.Bool(std.Zero, b.Implicit)),
		)
	}
	if a.Semicolon != b.Semicolon {
		return false, differs.ValuesOf(t, "Semicolon",
			sprints.Field(std.Zero, "Semicolon", sprints.Digit(a.Semicolon)),
			sprints.Field(std.Zero, "Semicolon", sprints.Digit(b.Semicolon)),
		)
	}
	return true, diff
}

func ExprStmt(a, b *ast.ExprStmt) (equals bool, diff string) {
	return exprStmt(std.Zero, a, b)
}

func exprStmt(t indent.Indentor, a, b *ast.ExprStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "ExprStmt", "*ast.ExprStmt"); has {
		return eq, d
	}
	if eq, dif := expr(t.Inc(), a.X, b.X); !eq {
		return false, t.Sprintf("'X' values mismatch:\n%s", dif)
	}
	return true, diff
}

func ForStmt(a, b *ast.ForStmt) (equals bool, diff string) {
	return forStmt(std.Zero, a, b)
}

func forStmt(t indent.Indentor, a, b *ast.ForStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "ForStmt", "*ast.ForStmt"); has {
		return eq, d
	}
	if a.For != b.For {
		return false, differs.ValuesOf(t, "For",
			sprints.Field(std.Zero, "For", sprints.Digit(a.For)),
			sprints.Field(std.Zero, "For", sprints.Digit(b.For)),
		)
	}
	if eq, dif := stmt(t.Inc(), a.Init, b.Init); !eq {
		return false, t.Sprintf("Init values mismatch:\n%s", dif)
	}
	if eq, dif := expr(t.Inc(), a.Cond, b.Cond); !eq {
		return false, t.Sprintf("Cond values mismatch:\n%s", dif)
	}
	if eq, dif := stmt(t.Inc(), a.Post, b.Post); !eq {
		return false, t.Sprintf("Post values mismatch:\n%s", dif)
	}
	if eq, dif := blockStmt(t.Inc(), a.Body, b.Body); !eq {
		return false, t.Sprintf("Body values mismatch:\n%s", dif)
	}
	return true, diff
}

func GoStmt(a, b *ast.GoStmt) (equals bool, diff string) {
	return goStmt(std.Zero, a, b)
}

func goStmt(t indent.Indentor, a, b *ast.GoStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "GoStmt", "*ast.GoStmt"); has {
		return eq, d
	}
	if a.Go != b.Go {
		return false, differs.ValuesOf(t, "Go",
			sprints.Field(std.Zero, "Go", sprints.Digit(a.Go)),
			sprints.Field(std.Zero, "Go", sprints.Digit(b.Go)),
		)
	}
	if eq, dif := callExpr(t.Inc(), a.Call, b.Call); !eq {
		return false, t.Sprintf("Call values mismatch:\n%s", dif)
	}
	return true, diff
}

func IfStmt(a, b *ast.IfStmt) (equals bool, diff string) {
	return ifStmt(std.Zero, a, b)
}

func ifStmt(t indent.Indentor, a, b *ast.IfStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "IfStmt", "*ast.IfStmt"); has {
		return eq, d
	}
	if a.If != b.If {
		return false, differs.ValuesOf(t, "If",
			sprints.Field(std.Zero, "If", sprints.Digit(a.If)),
			sprints.Field(std.Zero, "If", sprints.Digit(b.If)),
		)
	}
	if eq, dif := expr(t.Inc(), a.Cond, b.Cond); !eq {
		return false, t.Sprintf("'Cond' values mismatch:\n%s", dif)
	}
	if eq, dif := blockStmt(t.Inc(), a.Body, b.Body); !eq {
		return false, t.Sprintf("'Body' values mismatch:\n%s", dif)
	}
	if eq, dif := stmt(t.Inc(), a.Else, b.Else); !eq {
		return false, t.Sprintf("'Else' values mismatch:\n%s", dif)
	}
	if eq, dif := stmt(t.Inc(), a.Init, b.Init); !eq {
		return false, t.Sprintf("'Init' values mismatch:\n%s", dif)
	}
	return true, diff
}

func IncDecStmt(a, b *ast.IncDecStmt) (equals bool, diff string) {
	return incDecStmt(std.Zero, a, b)
}

func incDecStmt(t indent.Indentor, a, b *ast.IncDecStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "IncDecStmt", "*ast.IncDecStmt"); has {
		return eq, d
	}
	if a.TokPos != b.TokPos {
		return false, differs.ValuesOf(t, "TokPos",
			sprints.Field(std.Zero, "TokPos", sprints.Digit(a.TokPos)),
			sprints.Field(std.Zero, "TokPos", sprints.Digit(b.TokPos)),
		)
	}
	if a.Tok != b.Tok {
		return false, differs.ValuesOf(t, "Tok",
			sprints.Field(std.Zero, "Tok", sprints.Digit(a.Tok)),
			sprints.Field(std.Zero, "Tok", sprints.Digit(b.Tok)),
		)
	}
	if eq, dif := expr(t.Inc(), a.X, b.X); !eq {
		return false, t.Sprintf("'X' values mismatch:\n%s", dif)
	}
	return true, diff
}

func LabeledStmt(a, b *ast.LabeledStmt) (equals bool, diff string) {
	return labeledStmt(std.Zero, a, b)
}

func labeledStmt(t indent.Indentor, a, b *ast.LabeledStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "LabeledStmt", "*ast.LabeledStmt"); has {
		return eq, d
	}
	if a.Colon != b.Colon {
		return false, differs.ValuesOf(t, "Colon",
			sprints.Field(std.Zero, "Colon", sprints.Digit(a.Colon)),
			sprints.Field(std.Zero, "Colon", sprints.Digit(b.Colon)),
		)
	}
	if eq, dif := ident(t.Inc(), a.Label, b.Label); !eq {
		return false, t.Sprintf("'Label' values mismatch:\n%s", dif)
	}
	if eq, dif := stmt(t.Inc(), a.Stmt, b.Stmt); !eq {
		return false, t.Sprintf("'Stmt' values mismatch:\n%s", dif)
	}
	return true, diff
}

func RangeStmt(a, b *ast.RangeStmt) (equals bool, diff string) {
	return rangeStmt(std.Zero, a, b)
}

func rangeStmt(t indent.Indentor, a, b *ast.RangeStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "RangeStmt", "*ast.RangeStmt"); has {
		return eq, d
	}
	if a.For != b.For {
		return false, differs.ValuesOf(t, "For",
			sprints.Field(std.Zero, "For", sprints.Digit(a.For)),
			sprints.Field(std.Zero, "For", sprints.Digit(b.For)),
		)
	}
	if a.Range != b.Range {
		return false, differs.ValuesOf(t, "Range",
			sprints.Field(std.Zero, "Range", sprints.Digit(a.Range)),
			sprints.Field(std.Zero, "Range", sprints.Digit(b.Range)),
		)
	}
	if a.TokPos != b.TokPos {
		return false, differs.ValuesOf(t, "TokPos",
			sprints.Field(std.Zero, "TokPos", sprints.Digit(a.TokPos)),
			sprints.Field(std.Zero, "TokPos", sprints.Digit(b.TokPos)),
		)
	}
	if a.Tok != b.Tok {
		return false, differs.ValuesOf(t, "Tok",
			sprints.Field(std.Zero, "Tok", sprints.Digit(a.Tok)),
			sprints.Field(std.Zero, "Tok", sprints.Digit(b.Tok)),
		)
	}
	if eq, dif := expr(t.Inc(), a.Key, b.Key); !eq {
		return false, t.Sprintf("'Key' values mismatch:\n%s", dif)
	}
	if eq, dif := expr(t.Inc(), a.Value, b.Value); !eq {
		return false, t.Sprintf("'Value' values mismatch:\n%s", dif)
	}
	if eq, dif := expr(t.Inc(), a.X, b.X); !eq {
		return false, t.Sprintf("'X' values mismatch:\n%s", dif)
	}
	if eq, dif := stmt(t.Inc(), a.Body, b.Body); !eq {
		return false, t.Sprintf("'Body' values mismatch:\n%s", dif)
	}
	return true, diff
}

func ReturnStmt(a, b *ast.ReturnStmt) (equals bool, diff string) {
	return returnStmt(std.Zero, a, b)
}

func returnStmt(t indent.Indentor, a, b *ast.ReturnStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "ReturnStmt", "*ast.ReturnStmt"); has {
		return eq, d
	}
	if a.Return != b.Return {
		return false, differs.ValuesOf(t, "Return",
			sprints.Field(std.Zero, "Return", sprints.Digit(a.Return)),
			sprints.Field(std.Zero, "Return", sprints.Digit(b.Return)),
		)
	}
	for i := range a.Results {
		if eq, dif := expr(t.Inc(), a.Results[i], b.Results[i]); !eq {
			return false, t.Sprintf("'Results[%d]' values mismatch:\n%s", i, dif)
		}
	}
	return true, diff
}

func SelectStmt(a, b *ast.SelectStmt) (equals bool, diff string) {
	return selectStmt(std.Zero, a, b)
}

func selectStmt(t indent.Indentor, a, b *ast.SelectStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "SelectStmt", "*ast.SelectStmt"); has {
		return eq, d
	}
	if a.Select != b.Select {
		return false, differs.ValuesOf(t, "Select",
			sprints.Field(std.Zero, "Select", sprints.Digit(a.Select)),
			sprints.Field(std.Zero, "Select", sprints.Digit(b.Select)),
		)
	}
	if eq, dif := stmt(t.Inc(), a.Body, b.Body); !eq {
		return false, t.Sprintf("'Body' values mismatch:\n%s", dif)
	}
	return true, diff
}

func SendStmt(a, b *ast.SendStmt) (equals bool, diff string) {
	return sendStmt(std.Zero, a, b)
}

func sendStmt(t indent.Indentor, a, b *ast.SendStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "SendStmt", "*ast.SendStmt"); has {
		return eq, d
	}
	if a.Arrow != b.Arrow {
		return false, differs.ValuesOf(t, "Arrow",
			sprints.Field(std.Zero, "Arrow", sprints.Digit(a.Arrow)),
			sprints.Field(std.Zero, "Arrow", sprints.Digit(b.Arrow)),
		)
	}
	if eq, dif := expr(t.Inc(), a.Chan, b.Chan); !eq {
		return false, t.Sprintf("Chan values mismatch:\n%s", dif)
	}
	if eq, dif := expr(t.Inc(), a.Value, b.Value); !eq {
		return false, t.Sprintf("Value values mismatch:\n%s", dif)
	}
	return true, diff
}

func SwitchStmt(a, b *ast.SwitchStmt) (equals bool, diff string) {
	return switchStmt(std.Zero, a, b)
}

func switchStmt(t indent.Indentor, a, b *ast.SwitchStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "SwitchStmt", "*ast.SwitchStmt"); has {
		return eq, d
	}
	if a.Switch != b.Switch {
		return false, differs.ValuesOf(t, "Switch",
			sprints.Field(std.Zero, "Switch", sprints.Digit(a.Switch)),
			sprints.Field(std.Zero, "Switch", sprints.Digit(b.Switch)),
		)
	}
	if eq, dif := stmt(t.Inc(), a.Init, b.Init); !eq {
		return false, t.Sprintf("'Init' values mismatch:\n%s", dif)
	}
	if eq, dif := stmt(t.Inc(), a.Body, b.Body); !eq {
		return false, t.Sprintf("'Body' values mismatch:\n%s", dif)
	}
	if eq, dif := expr(t.Inc(), a.Tag, b.Tag); !eq {
		return false, t.Sprintf("'Tag' values mismatch:\n%s", dif)
	}
	return true, diff
}

func TypeSwitchStmt(a, b *ast.TypeSwitchStmt) (equals bool, diff string) {
	return typeSwitchStmt(std.Zero, a, b)
}

func typeSwitchStmt(t indent.Indentor, a, b *ast.TypeSwitchStmt) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "TypeSwitchStmt", "*ast.TypeSwitchStmt"); has {
		return eq, d
	}
	if a.Switch != b.Switch {
		return false, differs.ValuesOf(t, "Switch",
			sprints.Field(std.Zero, "Switch", sprints.Digit(a.Switch)),
			sprints.Field(std.Zero, "Switch", sprints.Digit(b.Switch)),
		)
	}
	if eq, dif := stmt(t.Inc(), a.Assign, b.Assign); !eq {
		return false, t.Sprintf("'Assign' values mismatch:\n%s", dif)
	}
	if eq, dif := stmt(t.Inc(), a.Body, b.Body); !eq {
		return false, t.Sprintf("'Body' values mismatch:\n%s", dif)
	}
	if eq, dif := stmt(t.Inc(), a.Init, b.Init); !eq {
		return false, t.Sprintf("'Init' values mismatch:\n%s", dif)
	}
	return true, diff
}
