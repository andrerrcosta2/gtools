// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package osys

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/seeders/random"
	"os"
	"path/filepath"
	"strings"
)

func mckGlob(pattern string) (matches []string, err error) {
	ext := filepath.Ext(pattern)
	return random.String(500, 15, 20).
		Map(func(s string) string {
			return fmt.Sprintf("%s%s", s, ext)
		}).
		Values(), nil
}

func mckReadOsFiles(path string) ([]byte, error) {
	if strings.Contains(path, "/nonexistent/") {
		return nil, &os.PathError{
			Op:   "open",
			Path: path,
			Err:  os.ErrNotExist,
		}
	}
	return random.Bytes(1, 1, 5).At(0), nil
}
