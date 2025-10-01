// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package osys

import (
	"github.com/andrerrcosta2/gtools/core/domain/gerrors"
	"github.com/andrerrcosta2/gtools/core/io/fsys"
	"os"
	"path/filepath"
	"testing"
)

const seedDir = "./"

func TestReadFiles_Success(t *testing.T) {
	pattern := filepath.Join(seedDir, "*.txt")

	files, err := mckGlob(pattern)
	if err != nil {
		t.Fatalf("failed to list files: %v", err)
	}

	if len(files) == 0 {
		t.Fatal("no test files found in directory")
	}

	paths := make([]fsys.Path, len(files))
	for i, file := range files {
		paths[i] = fsys.Path{Type: fsys.Relative, Path: file}
	}

	// Mock OS ops
	readOsFile = mckReadOsFiles
	defer func() {
		readOsFile = os.ReadFile
	}()

	// Call ReadFiles
	data, err := ReadFiles(paths, 6)
	if err != nil {
		if stack, ok := gerrors.AsStackable(err); ok {
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
	paths := []fsys.Path{
		{Type: fsys.Relative, Path: "./nonexistent/file1.txt"},
		{Type: fsys.Relative, Path: "./nonexistent/file2.txt"},
	}

	// Mock OS ops
	readOsFile = mckReadOsFiles
	defer func() {
		readOsFile = os.ReadFile
	}()

	// Call ReadFiles
	data, err := ReadFiles(paths, 6)
	if err != nil {
		if stack, ok := gerrors.AsStackable(err); !ok {
			t.Fatalf("Error should be stackable but is not: %v", err)
		} else {
			t.Logf("Error Stack: %s", stack.Trace())
		}
	}

	if data != nil {
		t.Fatalf("expected no data but got %v", data)
	}
}

func TestReadFiles_Concurrency(t *testing.T) {
	// Collect file paths
	files, err := mckGlob(filepath.Join(seedDir, "*.txt"))
	if err != nil {
		if stack, ok := gerrors.AsStackable(err); ok {
			t.Fatalf("failed to list files: %v", stack.Trace())
		} else {
			t.Fatalf("Error should be stackable but is not: %v", err)
		}
	}
	if len(files) < 3 {
		t.Fatal("not enough test files for concurrency test")
	}

	// Mock OS ops
	readOsFile = mckReadOsFiles
	defer func() {
		readOsFile = os.ReadFile
	}()

	// Select only the first 3 files for testing
	paths := []fsys.Path{
		{Type: fsys.Relative, Path: files[0]},
		{Type: fsys.Relative, Path: files[1]},
		{Type: fsys.Relative, Path: files[2]},
	}

	// Call ReadFiles
	data, err := ReadFiles(paths, 2)
	if err != nil {
		if stack, ok := gerrors.AsStackable(err); ok {
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
