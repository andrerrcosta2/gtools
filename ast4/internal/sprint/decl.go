// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/ast4/internal/std"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
)

func Decl(tab indent.Indentor, decl ast.Decl) string {
	if decl == nil {
		return "ast.Decl<nil>"
	}
	switch d := decl.(type) {
	case *ast.GenDecl:
		return GenDecl(tab, d)
	case *ast.FuncDecl:
		return FuncDecl(tab, d)
	default:
		return std.DeclTypeError(tab, decl)
	}
}

func FuncDecl(tab indent.Indentor, decl *ast.FuncDecl) string {
	if decl == nil {
		return "FuncDecl<nil>"
	}
	return sprints.Closedobj(tab, "FuncDecl",
		sprints.Ltfield(tab, "Doc", CommentGroup(tab.Inc(), decl.Doc)),
		sprints.Ltfield(tab, "Recv", FieldList(tab.Inc(), decl.Recv)),
		sprints.Ltfield(tab, "Name", Ident(tab.Inc(), decl.Name)),
		sprints.Ltfield(tab, "Type", FuncType(tab.Inc(), decl.Type)),
		sprints.Ltfield(tab, "Body", Stmt(tab.Inc(), decl.Body)))
}

func GenDecl(tab indent.Indentor, decl *ast.GenDecl) string {
	if decl == nil {
		return "GenDecl<nil>"
	}
	return sprints.Closedobj(tab, "GenDecl",
		sprints.Ltfield(tab, "Doc", CommentGroup(tab.Inc(), decl.Doc)),
		sprints.Ltfield(tab, "TokPos", sprints.Value(indent.Zero(), decl.TokPos)),
		sprints.Ltfield(tab, "Tok", sprints.Value(indent.Zero(), decl.Tok)),
		sprints.Ltfield(tab, "Lparen", "("),
		sprints.Ltfield(tab, "Specs", sprints.Slicef(tab.Inc(), decl.Specs, Spec)),
		sprints.Ltfield(tab, "Rparen", ")"))
}
