// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differs

import (
	"fmt"
	"go/ast"

	"github.com/andrerrcosta2/gtools/ast4/internal/std"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
)

func FieldList(a, b *ast.FieldList) (equals bool, diff string) {
	return fieldList(std.Zero, a, b)
}

func fieldList(t indent.Indentor, a, b *ast.FieldList) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "FieldList", "*ast.FieldList"); has {
		return eq, d
	}
	if a.Opening != b.Opening {
		return false, differs.ValuesOf(t, "Opening",
			sprints.Field(std.Zero, "Opening", sprints.Digit(a.Opening)),
			sprints.Field(std.Zero, "Opening", sprints.Digit(b.Opening)),
		)
	}
	if a.Closing != b.Closing {
		return false, differs.ValuesOf(t, "Closing",
			sprints.Field(std.Zero, "Closing", sprints.Digit(a.Closing)),
			sprints.Field(std.Zero, "Closing", sprints.Digit(b.Closing)),
		)
	}
	if len(a.List) != len(b.List) {
		return false, differs.SizesOf(t, "List",
			sprints.Field(std.Zero, "List", sprints.Digit(len(a.List))),
			sprints.Field(std.Zero, "List", sprints.Digit(len(b.List))),
		)
	}
	for i := 0; i < len(a.List); i++ {
		if eq, dif := field(t.Inc(), a.List[i], b.List[i]); !eq {
			return false, fmt.Sprintf("'List[%d]' values mismatch:\n%s", i, dif)
		}
	}
	return true, diff
}

func Field(a, b *ast.Field) (equals bool, diff string) {
	return field(std.Zero, a, b)
}

func field(t indent.Indentor, a, b *ast.Field) (equals bool, diff string) {
	if has, eq, d := nilDiff(t, a, b, "Field", "*ast.Field"); has {
		return eq, d
	}
	if eq, dif := commentGroup(t.Inc(), a.Doc, b.Doc); !eq {
		return false, t.Sprintf("'Doc' values mismatch:\n%s", dif)
	}
	if eq, dif := expr(t.Inc(), a.Type, b.Type); !eq {
		return false, t.Sprintf("'Type' values mismatch:\n%s", dif)
	}
	if eq, dif := basicLit(t.Inc(), a.Tag, b.Tag); !eq {
		return false, t.Sprintf("'Tag' values mismatch:\n%s", dif)
	}
	if eq, dif := commentGroup(t.Inc(), a.Comment, b.Comment); !eq {
		return false, t.Sprintf("'Comment' values mismatch:\n%s", dif)
	}
	for i := 0; i < len(a.Names); i++ {
		if eq, dif := ident(t.Inc(), a.Names[i], b.Names[i]); !eq {
			return false, t.Sprintf("'Names[%d]' values mismatch:\n%s", i, dif)
		}
	}
	return true, diff

}
