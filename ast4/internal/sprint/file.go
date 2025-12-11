// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/sprints"

	"go/ast"
)

func File(tab indent.Indentor, f *ast.File) string {
	if f == nil {
		return "File<nil>"
	}
	return sprints.Closedobj(tab, "File",
		sprints.Ltfield(tab, "Doc", CommentGroup(tab, f.Doc)),
		sprints.Ltfield(tab, "Package", sprints.Value(tab, f.Package)),
		sprints.Ltfield(tab, "Name", Expr(tab, f.Name)),
		sprints.Ltfield(tab, "Decls", sprints.Slicef(tab, f.Decls, Decl)),
		sprints.Ltfield(tab, "FileStart", sprints.Value(tab, f.FileStart)),
		sprints.Ltfield(tab, "FileEnd", sprints.Value(tab, f.FileEnd)),
		sprints.Ltfield(tab, "Imports", sprints.Slicef(tab, f.Imports, ImportSpec)),
		sprints.Ltfield(tab, "Unresolved", sprints.Slicef(tab, f.Unresolved, Ident)),
		sprints.Ltfield(tab, "Comments", sprints.Slicef(tab, f.Comments, CommentGroup)),
		sprints.Ltfield(tab, "GoVersion", f.GoVersion))
}
