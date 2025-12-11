// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differs

import (
	"go/ast"

	"github.com/andrerrcosta2/gtools/ast4/internal/std"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
)

func File(a, b *ast.File) (equals bool, diff string) {
	return file(std.Zero, a, b)
}

func file(t indent.Indentor, a, b *ast.File) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "File", "*ast.File"); has {
		return eq, d
	}
	if a.Package != b.Package {
		return false, differs.ValuesOf(t, "Package",
			sprints.Field(std.Zero, "Package", sprints.Digit(a.Package)),
			sprints.Field(std.Zero, "Package", sprints.Digit(b.Package)),
		)
	}
	if a.FileStart != b.FileStart {
		return false, differs.ValuesOf(t, "FileStart",
			sprints.Field(std.Zero, "FileStart", sprints.Digit(a.FileStart)),
			sprints.Field(std.Zero, "FileStart", sprints.Digit(b.FileStart)),
		)
	}
	if a.FileEnd != b.FileEnd {
		return false, differs.ValuesOf(t, "FileEnd",
			sprints.Field(std.Zero, "FileEnd", sprints.Digit(a.FileEnd)),
			sprints.Field(std.Zero, "FileEnd", sprints.Digit(b.FileEnd)),
		)
	}
	if a.GoVersion != b.GoVersion {
		return false, differs.ValuesOf(t, "GoVersion",
			sprints.Field(std.Zero, "GoVersion", a.GoVersion),
			sprints.Field(std.Zero, "GoVersion", b.GoVersion),
		)
	}
	if eq, dif := commentGroup(t.Inc(), a.Doc, b.Doc); !eq {
		return false, t.Sprintf("'Doc' values mismatch:\n%s", dif)
	}
	if eq, dif := ident(t.Inc(), a.Name, b.Name); !eq {
		return false, t.Sprintf("'Name' values mismatch:\n%s", dif)
	}
	for i, d := range a.Decls {
		if eq, dif := decl(t.Inc(), d, b.Decls[i]); !eq {
			return false, t.Sprintf("'Decls[%d]' values mismatch:\n%s", i, dif)
		}
	}

	for i, is := range a.Imports {
		if eq, dif := importSpec(t.Inc(), is, b.Imports[i]); !eq {
			return false, t.Sprintf("'Imports[%d]' values mismatch:\n%s", i, dif)
		}
	}
	for i, unresolved := range a.Unresolved {
		if eq, dif := ident(t.Inc(), unresolved, b.Unresolved[i]); !eq {
			return false, t.Sprintf("'Unresolved[%d'] values mismatch:\n%s", i, dif)
		}
	}
	for i, cg := range a.Comments {
		if eq, dif := commentGroup(t.Inc(), cg, b.Comments[i]); !eq {
			return false, t.Sprintf("C'omments[%d]' values mismatch:\n%s", i, dif)
		}
	}
	return true, diff
}
