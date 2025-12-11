// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package docs

import (
	"go/ast"
	"go/token"
	"strings"
)

// FindByPrefix scans the file and returns all comment groups
// (or individual comments) whose text starts with the given prefix.
func FindByPrefix(fset *token.FileSet, file *ast.File, prefix string) []*ast.CommentGroup {
	var matched []*ast.CommentGroup

	for _, cg := range file.Comments {
		for _, c := range cg.List {
			text := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
			if strings.HasPrefix(text, prefix) {
				matched = append(matched, cg)
				break // we found a match in this group
			}
		}
	}
	return matched
}

// FindCommentLinesByPrefix returns the actual comment lines that start with the prefix.
// Useful if you want precise per-line control rather than groups.
func FindCommentLinesByPrefix(fset *token.FileSet, file *ast.File, prefix string) []*ast.Comment {
	var matched []*ast.Comment

	for _, cg := range file.Comments {
		for _, c := range cg.List {
			text := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
			if strings.HasPrefix(text, prefix) {
				matched = append(matched, c)
			}
		}
	}
	return matched
}

type NodeCommentPair struct {
	Node    ast.Node
	Comment *ast.CommentGroup
	Pos     token.Position
	Text    string
}

// FindNodeByCommentPrefix walks through the file and returns all (node, comment) pairs
// where the node’s leading doc starts with the given prefix.
func FindNodeByCommentPrefix(fset *token.FileSet, file *ast.File, prefix string) []NodeCommentPair {
	var results []NodeCommentPair

	ast.Inspect(file, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		var doc *ast.CommentGroup

		switch x := n.(type) {
		case *ast.GenDecl:
			doc = x.Doc
		case *ast.FuncDecl:
			doc = x.Doc
		case *ast.TypeSpec:
			doc = x.Doc
		case *ast.Field:
			doc = x.Doc
		case *ast.ValueSpec:
			doc = x.Doc
		default:
			return true
		}

		if doc == nil {
			return true
		}

		for _, c := range doc.List {
			text := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
			if strings.HasPrefix(text, prefix) {
				results = append(results, NodeCommentPair{
					Node:    n,
					Comment: doc,
					Text:    text,
					Pos:     fset.Position(c.Pos()),
				})
				break
			}
		}
		return true
	})

	return results
}
