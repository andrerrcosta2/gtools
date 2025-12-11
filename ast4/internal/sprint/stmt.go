// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"github.com/andrerrcosta2/gtools/ast4/internal/std"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"

	"go/ast"
)

func Stmt(tab indent.Indentor, stmt ast.Stmt) string {
	if stmt == nil {
		return "ast.Stmt<nil>"
	}
	switch s := stmt.(type) {
	case *ast.AssignStmt:
		return AssignStmt(tab, s)
	case *ast.BadStmt:
		return BadStmt(tab, s)
	case *ast.BlockStmt:
		return BlockStmt(tab, s)
	case *ast.BranchStmt:
		return BranchStmt(tab, s)
	case *ast.CaseClause:
		return CaseClause(tab, s)
	case *ast.CommClause:
		return CommClause(tab, s)
	case *ast.DeclStmt:
		return DeclStmt(tab, s)
	case *ast.DeferStmt:
		return DeferStmt(tab, s)
	case *ast.EmptyStmt:
		return EmptyStmt(tab, s)
	case *ast.ExprStmt:
		return ExprStmt(tab, s)
	case *ast.ForStmt:
		return ForStmt(tab, s)
	case *ast.GoStmt:
		return GoStmt(tab, s)
	case *ast.IfStmt:
		return IfStmt(tab, s)
	case *ast.IncDecStmt:
		return IncDecStmt(tab, s)
	case *ast.LabeledStmt:
		return LabeledStmt(tab, s)
	case *ast.RangeStmt:
		return RangeStmt(tab, s)
	case *ast.ReturnStmt:
		return ReturnStmt(tab, s)
	case *ast.SelectStmt:
		return SelectStmt(tab, s)
	case *ast.SendStmt:
		return SendStmt(tab, s)
	case *ast.SwitchStmt:
		return SwitchStmt(tab, s)
	case *ast.TypeSwitchStmt:
		return TypeSwitchStmt(tab, s)
	default:
		return std.StmtTypeError(tab, stmt)
	}
}

func AssignStmt(tab indent.Indentor, s *ast.AssignStmt) string {
	if s == nil {
		return "AssignStmt<nil>"
	}
	return sprints.Closedobj(tab, "AssignStmt",
		sprints.Ltfield(tab, "Lhs", sprints.Slicef(tab, s.Lhs, Expr)),
		sprints.Ltfield(tab, "TokPos", sprints.Value(tab, s.TokPos)),
		sprints.Ltfield(tab, "Tok", sprints.Value(tab, s.Tok)),
		sprints.Ltfield(tab, "Rhs", sprints.Slicef(tab, s.Rhs, Expr)))
}

func BadStmt(tab indent.Indentor, s *ast.BadStmt) string {
	if s == nil {
		return sprints.Error(tab, "BadStmt<nil>")
	}
	return sprints.Error(indent.Zero(), sprints.Closedobj(tab, "BadStmt",
		sprints.Ltfield(tab, "From", sprints.Value(tab, s.From)),
		sprints.Ltfield(tab, "To", sprints.Value(tab, s.To))))
}

func BlockStmt(tab indent.Indentor, s *ast.BlockStmt) string {
	if s == nil {
		return "BlockStmt<nil>"
	}
	return sprints.Closedobj(tab, "BlockStmt",
		sprints.Ltfield(tab, "Lbrace", sprints.Value(tab, s.Lbrace)),
		sprints.Ltfield(tab, "List", sprints.Slicef(tab, s.List, Stmt)),
		sprints.Ltfield(tab, "Rbrace", sprints.Value(tab, s.Rbrace)))
}
func BranchStmt(tab indent.Indentor, s *ast.BranchStmt) string {
	if s == nil {
		return "BranchStmt<nil>"
	}
	return sprints.Closedobj(tab, "BranchStmt",
		sprints.Ltfield(tab, "TokPos", sprints.Value(tab, s.TokPos)),
		sprints.Ltfield(tab, "Tok", sprints.Value(tab, s.Tok)),
		sprints.Ltfield(tab, "Label", Ident(tab, s.Label)))
}
func CaseClause(tab indent.Indentor, s *ast.CaseClause) string {
	if s == nil {
		return "CaseClause<nil>"
	}
	return sprints.Closedobj(tab, "CaseClause",
		sprints.Ltfield(tab, "Case", sprints.Value(tab, s.Case)),
		sprints.Ltfield(tab, "List", sprints.Slicef(tab, s.List, Expr)),
		sprints.Ltfield(tab, "Colon", sprints.Value(tab, s.Colon)),
		sprints.Ltfield(tab, "Body", sprints.Slicef(tab, s.Body, Stmt)))
}

func CommClause(tab indent.Indentor, s *ast.CommClause) string {
	if s == nil {
		return "CommClause<nil>"
	}
	return sprints.Closedobj(tab, "CommClause",
		sprints.Ltfield(tab, "Case", sprints.Value(tab, s.Case)),
		sprints.Ltfield(tab, "Comm", Stmt(tab, s.Comm)),
		sprints.Ltfield(tab, "Colon", sprints.Value(tab, s.Colon)),
		sprints.Ltfield(tab, "Body", sprints.Slicef(tab, s.Body, Stmt)))
}

func DeclStmt(tab indent.Indentor, s *ast.DeclStmt) string {
	if s == nil {
		return "DeclStmt<nil>"
	}
	return sprints.Closedobj(tab, "DeclStmt",
		sprints.Ltfield(tab, "Decl", Decl(tab, s.Decl)))
}

func DeferStmt(tab indent.Indentor, s *ast.DeferStmt) string {
	if s == nil {
		return "DeferStmt<nil>"
	}
	return sprints.Closedobj(tab, "DeferStmt",
		sprints.Ltfield(tab, "Defer", sprints.Value(tab, s.Defer)),
		sprints.Ltfield(tab, "Call", Expr(tab, s.Call)))
}

func EmptyStmt(tab indent.Indentor, s *ast.EmptyStmt) string {
	if s == nil {
		return "EmptyStmt<nil>"
	}
	return sprints.Closedobj(tab, "EmptyStmt",
		sprints.Ltfield(tab, "Semicolon", sprints.Value(tab, s.Semicolon)))
}

func ExprStmt(tab indent.Indentor, s *ast.ExprStmt) string {
	if s == nil {
		return "ExprStmt<nil>"
	}
	return sprints.Closedobj(tab, "ExprStmt",
		sprints.Ltfield(tab, "X", Expr(tab, s.X)))
}

func ForStmt(tab indent.Indentor, s *ast.ForStmt) string {
	if s == nil {
		return "ForStmt<nil>"
	}
	return sprints.Closedobj(tab, "ForStmt",
		sprints.Ltfield(tab, "For", sprints.Value(tab, s.For)),
		sprints.Ltfield(tab, "Init", Stmt(tab, s.Init)),
		sprints.Ltfield(tab, "Cond", Expr(tab, s.Cond)),
		sprints.Ltfield(tab, "Post", Stmt(tab, s.Post)),
		sprints.Ltfield(tab, "Body", Stmt(tab, s.Body)))
}

func GoStmt(tab indent.Indentor, s *ast.GoStmt) string {
	if s == nil {
		return "GoStmt<nil>"
	}
	return sprints.Closedobj(tab, "GoStmt",
		sprints.Ltfield(tab, "Go", sprints.Value(tab, s.Go)),
		sprints.Ltfield(tab, "Call", Expr(tab, s.Call)))
}

func IfStmt(tab indent.Indentor, s *ast.IfStmt) string {
	if s == nil {
		return "IfStmt<nil>"
	}
	return sprints.Closedobj(tab, "IfStmt",
		sprints.Ltfield(tab, "If", sprints.Value(tab, s.If)),
		sprints.Ltfield(tab, "Init", Stmt(tab, s.Init)),
		sprints.Ltfield(tab, "Cond", Expr(tab, s.Cond)),
		sprints.Ltfield(tab, "Body", Stmt(tab, s.Body)),
		sprints.Ltfield(tab, "Else", Stmt(tab, s.Else)))
}

func IncDecStmt(tab indent.Indentor, s *ast.IncDecStmt) string {
	if s == nil {
		return "IncDecStmt<nil>"
	}
	return sprints.Closedobj(tab, "IncDecStmt",
		sprints.Ltfield(tab, "X", Expr(tab, s.X)),
		sprints.Ltfield(tab, "TokPos", sprints.Value(tab, s.TokPos)),
		sprints.Ltfield(tab, "Tok", sprints.Value(tab, s.Tok)))
}

func LabeledStmt(tab indent.Indentor, s *ast.LabeledStmt) string {
	if s == nil {
		return "LabeledStmt<nil>"
	}
	return sprints.Closedobj(tab, "LabeledStmt",
		sprints.Ltfield(tab, "Label", Ident(tab, s.Label)),
		sprints.Ltfield(tab, "Colon", sprints.Value(tab, s.Colon)),
		sprints.Ltfield(tab, "Stmt", Stmt(tab, s.Stmt)))
}

func RangeStmt(tab indent.Indentor, s *ast.RangeStmt) string {
	if s == nil {
		return "RangeStmt<nil>"
	}
	return sprints.Closedobj(tab, "RangeStmt",
		sprints.Ltfield(tab, "For", sprints.Value(tab, s.For)),
		sprints.Ltfield(tab, "Key", Expr(tab, s.Key)),
		sprints.Ltfield(tab, "Value", Expr(tab, s.Value)),
		sprints.Ltfield(tab, "TokPos", sprints.Value(tab, s.TokPos)),
		sprints.Ltfield(tab, "Tok", sprints.Value(tab, s.Tok)),
		sprints.Ltfield(tab, "X", Expr(tab, s.X)),
		sprints.Ltfield(tab, "Body", Stmt(tab, s.Body)))
}

func ReturnStmt(tab indent.Indentor, s *ast.ReturnStmt) string {
	if s == nil {
		return "ReturnStmt<nil>"
	}
	return sprints.Closedobj(tab, "ReturnStmt",
		sprints.Ltfield(tab, "Return", sprints.Value(tab, s.Return)),
		sprints.Ltfield(tab, "Results", sprints.Slicef(tab, s.Results, Expr)))
}

func SelectStmt(tab indent.Indentor, s *ast.SelectStmt) string {
	if s == nil {
		return "SelectStmt<nil>"
	}
	return sprints.Closedobj(tab, "SelectStmt",
		sprints.Ltfield(tab, "Select", sprints.Value(tab, s.Select)),
		sprints.Ltfield(tab, "Body", Stmt(tab, s.Body)))
}

func SendStmt(tab indent.Indentor, s *ast.SendStmt) string {
	if s == nil {
		return "SendStmt<nil>"
	}
	return sprints.Closedobj(tab, "SendStmt",
		sprints.Ltfield(tab, "Chan", Expr(tab, s.Chan)),
		sprints.Ltfield(tab, "Arrow", sprints.Value(tab, s.Arrow)),
		sprints.Ltfield(tab, "Value", Expr(tab, s.Value)))
}

func SwitchStmt(tab indent.Indentor, s *ast.SwitchStmt) string {
	if s == nil {
		return "SwitchStmt<nil>"
	}
	return sprints.Closedobj(tab, "SwitchStmt",
		sprints.Ltfield(tab, "Switch", sprints.Value(tab, s.Switch)),
		sprints.Ltfield(tab, "Init", Stmt(tab, s.Init)),
		sprints.Ltfield(tab, "Tag", Expr(tab, s.Tag)),
		sprints.Ltfield(tab, "Body", Stmt(tab, s.Body)))
}

func TypeSwitchStmt(tab indent.Indentor, s *ast.TypeSwitchStmt) string {
	if s == nil {
		return "TypeSwitchStmt<nil>"
	}
	return sprints.Closedobj(tab, "TypeSwitchStmt",
		sprints.Ltfield(tab, "Switch", sprints.Value(tab, s.Switch)),
		sprints.Ltfield(tab, "Init", Stmt(tab, s.Init)),
		sprints.Ltfield(tab, "Assign", Stmt(tab, s.Assign)),
		sprints.Ltfield(tab, "Body", Stmt(tab, s.Body)))
}
