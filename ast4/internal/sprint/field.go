// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"

	"go/ast"
)

func FieldList(tab indent.Indentor, fields *ast.FieldList) string {
	if fields == nil {
		return "FieldList<nil>"
	}
	return sprints.Closedobj(tab, "FieldList",
		sprints.Ltfield(tab, "Opening", sprints.Value(tab, fields.Opening)),
		sprints.Ltfield(tab, "List", sprints.Slicef(tab, fields.List, Field)),
		sprints.Ltfield(tab, "Closing", sprints.Value(tab, fields.Closing)))
}

func Field(tab indent.Indentor, field *ast.Field) string {
	if field == nil {
		return "Field<nil>"
	}
	return sprints.Closedobj(tab, "Field",
		sprints.Ltfield(tab, "Doc", CommentGroup(tab, field.Doc)),
		sprints.Ltfield(tab, "Names", sprints.Slicef(tab, field.Names, Ident)),
		sprints.Ltfield(tab, "Type", Expr(tab, field.Type)),
		sprints.Ltfield(tab, "Tag", Expr(tab, field.Tag)),
		sprints.Ltfield(tab, "Comment", CommentGroup(tab, field.Comment)))
}
