// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package osutil

import (
	"github.com/andrerrcosta2/gtools/pkg/fsutil"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
	"path/filepath"
	"testing"
)

const seedDir = "./../_test/seed/txt/"

func TestReadFiles_Success(t *testing.T) {
	pattern := filepath.Join(seedDir, "*.txt")

	files, err := filepath.Glob(pattern)
	if err != nil {
		t.Fatalf("failed to list files: %v", err)
	}
	if len(files) == 0 {
		t.Fatal("no test files found in directory")
	}

	paths := make([]fsutil.Path, len(files))
	for i, file := range files {
		paths[i] = fsutil.Path{Type: fsutil.Relative, Path: file}
	}

	// Call ReadFiles
	data, err := ReadFiles(paths, gtools.NewSemaphore(10))
	if err != nil {
		if stack, ok := gtools.AsStackable(err); ok {
			t.Fatalf("ReadFiles returned an error: %v", stack.Trace())
		} else {
			t.Fatalf("Error should be stackable but is not: %v", err)
		}
	}

	if len(data) != len(files) {
		t.Fatalf("expected %d files to be read, got %d", len(files), len(data))
	}

	for i, d := range data {
		if len(d) == 0 {
			t.Errorf("file %s was read but has no data", files[i])
		}
	}
}

func TestReadFiles_Error(t *testing.T) {
	paths := []fsutil.Path{
		{Type: fsutil.Root, Path: "./_test/seed/txt/nonexistent_file1.txt"},
		{Type: fsutil.Root, Path: "./_test/seed/txt/nonexistent_file2.txt"},
	}

	// Semaphore to allow all goroutines to run concurrently
	sem := gtools.NewSemaphore(2) // Limit concurrency to 2

	// Call ReadFiles
	data, err := ReadFiles(paths, sem)
	if err == nil {
		t.Fatal("expected an error but got none")
	}
	if data != nil {
		t.Fatal("expected no data but got some")
	}
}

func TestReadFiles_Concurrency(t *testing.T) {
	// Collect file paths
	files, err := filepath.Glob(filepath.Join(seedDir, "*.txt"))
	if err != nil {
		if stack, ok := gtools.AsStackable(err); ok {
			t.Fatalf("failed to list files: %v", stack.Trace())
		} else {
			t.Fatalf("Error should be stackable but is not: %v", err)
		}
	}
	if len(files) < 3 {
		t.Fatal("not enough test files for concurrency test")
	}

	// Select only the first 3 files for testing
	paths := []fsutil.Path{
		{Type: fsutil.Relative, Path: files[0]},
		{Type: fsutil.Relative, Path: files[1]},
		{Type: fsutil.Relative, Path: files[2]},
	}

	sem := gtools.NewSemaphore(2) // Limit concurrency to 2

	// Call ReadFiles
	data, err := ReadFiles(paths, sem)
	if err != nil {
		if stack, ok := gtools.AsStackable(err); ok {
			t.Fatalf("ReadFiles returned an error: %v", stack.Trace())
		} else {
			t.Fatalf("Error should be stackable but is not: %v", err)
		}
	}

	if len(data) != len(paths) {
		t.Fatalf("expected %d files to be read, got %d", len(paths), len(data))
	}

	for i, d := range data {
		if len(d) == 0 {
			t.Errorf("file %s was read but has no data", paths[i].Path)
		}
	}
}

func TestReadFiles_SemaphoreLimit(t *testing.T) {
	// Collect file paths
	files, err := filepath.Glob(filepath.Join(seedDir, "*.txt"))
	if err != nil {
		t.Fatalf("failed to list files: %v", err)
	}
	if len(files) < 10 {
		t.Fatal("not enough test files for semaphore limit test")
	}

	// Create paths with a large number of files
	paths := make([]fsutil.Path, 20)
	for i := 0; i < 20; i++ {
		paths[i] = fsutil.Path{Type: fsutil.Relative, Path: files[i]}
	}

	sem := gtools.NewSemaphore(5) // Limit concurrency to 5

	// Call ReadFiles
	data, err := ReadFiles(paths, sem)
	if err != nil {
		if stack, ok := gtools.AsStackable(err); ok {
			t.Fatalf("ReadFiles returned an error: %v", stack.Trace())
		} else {
			t.Fatalf("Error should be stackable but is not: %v", err)
		}
	}

	if len(data) != len(paths) {
		t.Fatalf("expected %d files to be read, got %d", len(paths), len(data))
	}

	for i, d := range data {
		if len(d) == 0 {
			t.Errorf("file %s was read but has no data", paths[i].Path)
		}
	}
}
