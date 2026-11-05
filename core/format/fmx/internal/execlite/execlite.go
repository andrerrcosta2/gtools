// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package execlite

import (
	"os"
	"path/filepath"
)

var SupportsTerminalHyperlink bool

func init() {
	term := os.Getenv("TERM")
	wtSession := os.Getenv("WT_SESSION")
	termProgram := os.Getenv("TERM_PROGRAM")

	// Keep previous behavior while avoiding dependency on core/domain/env.
	SupportsTerminalHyperlink = term == "" && (wtSession == "" || termProgram == "vscode")
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

	return filepath.Clean(rel)
}
