// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package execution

import (
	"github.com/andrerrcosta2/gtools/core/domain/env"
	"os"
	"path/filepath"
)

var SupportsTerminalHyperlink bool

func init() {
	// Determine if hyperlinks are supported during initialization
	SupportsTerminalHyperlink = !env.TERM.Exists() &&
		(!env.WT_SESSION.Exists() || env.TERM_PROGRAM.Equals("vscode"))
}

func GetFileDisplayPath(file string, root string) string {
	if root == "" {
		root = "PROJECT_ROOT"
	}
	projectRoot := os.Getenv(root)
	if projectRoot == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return filepath.Base(file)
		}
		projectRoot = cwd
	}

	absFile, err := filepath.Abs(file)
	if err != nil {
		return filepath.Base(file)
	}

	absRoot, err := filepath.Abs(projectRoot)
	if err != nil {
		return filepath.Base(file)
	}

	rel, err := filepath.Rel(absRoot, absFile)
	if err != nil {
		return filepath.Base(file)
	}

	// Normalize and remove excessive ".." parts
	cleaned := filepath.Clean(rel)
	return cleaned
}
