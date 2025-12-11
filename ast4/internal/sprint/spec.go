// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/ast4/internal/std"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
)

func Spec(tab indent.Indentor, spec ast.Spec) string {
	if spec == nil {
		return "ast.Spec<nil>"
	}
	switch s := spec.(type) {
	case *ast.ImportSpec:
		return ImportSpec(tab, s)
	case *ast.TypeSpec:
		return TypeSpec(tab, s)
	case *ast.ValueSpec:
		return ValueSpec(tab, s)
	default:
		return std.SpecTypeError(tab, s)
	}
}

func ImportSpec(tab indent.Indentor, s *ast.ImportSpec) string {
	if s == nil {
		return "ImportSpec<nil>"
	}
	return sprints.Closedobj(tab, "ImportSpec",
		sprints.Ltfield(tab, "Doc", CommentGroup(tab, s.Doc)),
		sprints.Ltfield(tab, "Name", Ident(tab, s.Name)),
		sprints.Ltfield(tab, "Path", Expr(tab, s.Path)),
		sprints.Ltfield(tab, "Comment", CommentGroup(tab, s.Comment)),
		sprints.Ltfield(tab, "EndPos", sprints.Value(tab, s.EndPos)))
}

func TypeSpec(tab indent.Indentor, s *ast.TypeSpec) string {
	if s == nil {
		return "TypeSpec<nil>"
	}
	return sprints.Closedobj(tab, "TypeSpec",
		sprints.Ltfield(tab, "Doc", CommentGroup(tab, s.Doc)),
		sprints.Ltfield(tab, "Name", Ident(tab, s.Name)),
		sprints.Ltfield(tab, "TypeParams", FieldList(tab, s.TypeParams)),
		sprints.Ltfield(tab, "Assign", sprints.Value(tab, s.Assign)),
		sprints.Ltfield(tab, "Type", Expr(tab, s.Type)),
		sprints.Ltfield(tab, "Comment", CommentGroup(tab, s.Comment)))
}

func ValueSpec(tab indent.Indentor, s *ast.ValueSpec) string {
	if s == nil {
		return "ValueSpec<nil>"
	}
	return sprints.Closedobj(tab, "ValueSpec",
		sprints.Ltfield(tab, "Doc", CommentGroup(tab, s.Doc)),
		sprints.Ltfield(tab, "Names", sprints.Slicef(tab, s.Names, Ident)),
		sprints.Ltfield(tab, "Type", Expr(tab, s.Type)),
		sprints.Ltfield(tab, "Values", sprints.Slicef(tab, s.Values, Expr)),
		sprints.Ltfield(tab, "Comment", CommentGroup(tab, s.Comment)))
}
