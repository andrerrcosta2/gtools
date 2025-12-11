// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"

	"go/ast"
)

func CommentGroup(tab indent.Indentor, doc *ast.CommentGroup) string {
	if doc == nil {
		return "CommentGroup<nil>"
	}
	return sprints.Closedobj(tab, "CommentGroup",
		sprints.Ltfield(tab, "List", sprints.Slicef(tab, doc.List, Comment)))
}

func Comment(tab indent.Indentor, doc *ast.Comment) string {
	if doc == nil {
		return "Comment<nil>"
	}
	return sprints.Closedobj(tab, "Comment",
		sprints.Ltfield(tab, "Slash", sprints.Value(indent.Zero(), doc.Slash)),
		sprints.Ltfield(tab, "Text", sprints.Value(indent.Zero(), doc.Text)))
}
