// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package osutil

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/pkg/fsutil"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
	"os"
	"sync"
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

func ReadFiles(paths []fsutil.Path, sem *gtools.Semaphore) ([][]byte, error) {
	var wg sync.WaitGroup
	stack := &gtools.ConcurrentStackableError{}
	dataSlices := make([][]byte, len(paths))
	wg.Add(len(paths))

	for i, path := range paths {
		go func(i int, path fsutil.Path) {
			defer wg.Done()

			sem.Acq()
			defer sem.Rls()

			data, err := ReadFile(path.Type, path.Path)
			if err != nil {
				stack.Stack(fmt.Errorf("error reading file '%s': %v", path.Path, err))
				return
			}
			dataSlices[i] = data
		}(i, path)
	}

	wg.Wait()

	if !stack.Empty() {
		return nil, stack
	}
	return dataSlices, nil
}

func ReadPartial(mode fsutil.PathType, path string, start int64, end int64) ([]byte, error) {
	fp, err := fsutil.BuildPath(mode, path)
	file, err := os.Open(fp)
	if err != nil {
		return nil, fmt.Errorf("error opening file '%s': %v", path, err)
	}
	defer file.Close()

	size := end - start
	if size <= 0 {
		return nil, fmt.Errorf("invalid range: end (%d) must be greater than start (%d)", end, start)
	}

	if _, err := file.Seek(start, 0); err != nil {
		return nil, fmt.Errorf("error seeking to start position (%d): %v", start, err)
	}

	data := make([]byte, size)
	n, err := file.Read(data)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	if int64(n) < size {
		return data[:n], nil
	}
	return data, nil
}
