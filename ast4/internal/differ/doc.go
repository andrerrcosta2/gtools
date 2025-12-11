// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differs

import (
	"github.com/andrerrcosta2/gtools/ast4/internal/std"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/sprints"

	"go/ast"
)

func CommentGroup(doc, other *ast.CommentGroup) (equals bool, diff string) {
	return commentGroup(std.Zero, doc, other)
}

func commentGroup(t indent.Indentor, a, b *ast.CommentGroup) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "CommentGroup", "*ast.CommentGroup"); has {
		return eq, d
	}
	for i := 0; i < len(a.List); i++ {
		if eq, dif := comment(t.Inc(), a.List[i], b.List[i]); !eq {
			return false, t.Sprintf("List[%d] values mismatch: %s", i, dif)
		}
	}
	return true, diff
}

func Comment(doc, other *ast.Comment) (equals bool, diff string) {
	return comment(std.Zero, doc, other)
}

func comment(t indent.Indentor, a, b *ast.Comment) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "Comment", "*ast.Comment"); has {
		return eq, d
	}
	if a.Slash != b.Slash {
		return false, differs.ValuesOf(t, "Slash",
			sprints.Field(t, "Slash", sprints.Digit(a.Slash)),
			sprints.Field(t, "Slash", sprints.Digit(b.Slash)))
	}
	if a.Text != b.Text {
		return false, differs.ValuesOf(t, "Text",
			sprints.Field(t, "Text", a.Text),
			sprints.Field(t, "Text", b.Text))
	}
	return true, diff
}
