// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package osutil

import (
	"com.github/andrerrcosta2/gtools/pkg/fsutil"
	"os"
)

func ReadFile(mode fsutil.PathType, path string) ([]byte, error) {
	fp, err := fsutil.BuildPath(mode, path)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	return data, nil
}
