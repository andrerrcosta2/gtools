// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package ask

import "go/ast"

func IsSimilarFile(a, b *ast.File) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	if len(a.Decls) != len(b.Decls) {
		return false
	}
	for i, decl := range a.Decls {
		if !IsSimilarDecl(decl, b.Decls[i]) {
			return false
		}
	}
	if len(a.Imports) != len(b.Imports) {
		return false
	}
	for i, imp := range a.Imports {
		if !IsSimilarImportSpec(imp, b.Imports[i]) {
			return false
		}
	}
	if len(a.Unresolved) != len(b.Unresolved) {
		return false
	}
	for i, u := range a.Unresolved {
		if u.Name != b.Unresolved[i].Name {
			return false
		}
	}
	return a.Name.Name == b.Name.Name && a.GoVersion == b.GoVersion
}
