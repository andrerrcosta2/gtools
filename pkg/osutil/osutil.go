// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package osutil

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/pkg/fsutil"
	"github.com/andrerrcosta2/gtools/pkg/gtools"
	"io/fs"
	"os"
	"sync"
)

// ReadFile reads the contents of a file at the specified path.
// The path is resolved based on the provided mode, which can be either Relative or Root.
//
// Args:
//
//	mode (fsutil.PathType): The type of path to resolve.
//	path (string): The path to the file.
//
// Returns:
//
//	([]byte, error): The contents of the file as a byte slice, or an error if the file cannot be read.
func ReadFile(mode fsutil.PathType, path string) ([]byte, error) {
	// Resolve the file path based on the provided mode
	fp, err := fsutil.BuildPath(mode, path)
	if err != nil {
		// Return the error if the path cannot be resolved
		return nil, err
	}

	// Attempt to read the file at the resolved path
	data, err := os.ReadFile(fp)
	if err != nil {
		// Return the error if the file cannot be read
		return nil, err
	}
	// Return the file contents and no error
	return data, nil
}

// ReadFiles reads the contents of multiple files concurrently.
//
// Args:
//
//	paths ([]fsutil.Path): A slice of file paths to read.
//	sem (gtools.Semaphore): A semaphore to limit concurrency.
//
// Returns:
//
//	([][]byte, error): A slice of byte slices containing the file contents, or an error if any file cannot be read.
func ReadFiles(paths []fsutil.Path, sem gtools.Semaphore) ([][]byte, error) {
	// Create a wait group to synchronize the goroutines
	var wg sync.WaitGroup

	// Create a stackable error to collect any errors that occur
	stack := &gtools.ConcurrentStackableError{}

	// Initialize a slice to store the file contents
	dataSlices := make([][]byte, len(paths))

	// Add the number of paths to the wait group
	wg.Add(len(paths))

	// Iterate over the paths and start a goroutine for each one
	for i, path := range paths {
		go func(i int, path fsutil.Path) {
			// Defer the wait group done call to ensure it's called even if an error occurs
			defer wg.Done()

			// Acquire the semaphore to limit concurrency
			sem.Acq()
			defer sem.Rls()

			// Read the file at the specified path
			data, err := ReadFile(path.Type, path.Path)
			if err != nil {
				// Stack the error if the file cannot be read
				stack.Stack(fmt.Errorf("error reading file '%s': %v", path.Path, err))
				return
			}
			// Store the file contents in the data slice
			dataSlices[i] = data
		}(i, path)
	}

	// Wait for all the goroutines to finish
	wg.Wait()

	// Check if any errors occurred
	if !stack.Empty() {
		// Return the stacked error
		return nil, stack
	}
	// Return the file contents and no error
	return dataSlices, nil
}

// ReadPartial reads a partial content of a file.
//
// It takes a file path, a start position, and an end position as input, and returns the content of the file
// within the specified range. If an error occurs during the read operation, it returns an error.
func ReadPartial(mode fsutil.PathType, path string, start int64, end int64) ([]byte, error) {
	// Build the full file path based on the provided mode and path
	fp, err := fsutil.BuildPath(mode, path)
	if err != nil {
		// Return an error if the file path cannot be built
		return nil, fmt.Errorf("error building file path '%s': %v", path, err)
	}

	// Open the file in read-only mode
	file, err := os.Open(fp)
	if err != nil {
		// Return an error if the file cannot be opened
		return nil, fmt.Errorf("error opening file '%s': %v", path, err)
	}
	defer file.Close() // Close the file when we're done with it

	// Calculate the size of the range to read
	size := end - start
	if size <= 0 {
		// Return an error if the range is invalid
		return nil, fmt.Errorf("invalid range: end (%d) must be greater than start (%d)", end, start)
	}

	// Seek to the start position of the range
	if _, err := file.Seek(start, 0); err != nil {
		// Return an error if the seek operation fails
		return nil, fmt.Errorf("error seeking to start position (%d): %v", start, err)
	}

	// Create a byte slice to store the read data
	data := make([]byte, size)

	// Read the data from the file
	n, err := file.Read(data)
	if err != nil {
		// Return an error if the read operation fails
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// If the read data is shorter than the expected size, return the truncated data
	if int64(n) < size {
		return data[:n], nil
	}
	return data, nil
}

// MkdirAll creates a directory with the given package name and operation.
// It joins the operation path with the package name, replacing any '/' with the OS-specific path separator.
// The directory is created with permissions to read, write, and execute for all users.
//
// Args:
//
//	mode (fsutil.PathType): The type of path to construct (e.g., relative or root).
//	path (string): The path to the directory to create.
//	perm (fs.FileMode): The permissions to apply to the created directory.
//
// Returns:
//
//	error: An error if the directory cannot be created, or nil if successful.
func MkdirAll(mode fsutil.PathType, path string, perm fs.FileMode) error {
	// Construct the full path to the directory by joining the mode and path
	// This step is necessary to ensure the correct path separator is used for the operating system
	dir, err := fsutil.BuildPath(mode, path)
	if err != nil {
		// If an error occurs during path construction, return it immediately
		// This allows the caller to handle the error as needed
		return err
	}

	// Create the directory and all its parents if they don't exist
	// The os.MkdirAll function will create all necessary parent directories
	err = os.MkdirAll(dir, perm)
	if err != nil {
		// If an error occurs during directory creation, return it with a descriptive message
		// This provides more context to the caller about the nature of the error
		return fmt.Errorf("failed to create directory '%s': %v", dir, err)
	}
	// If no errors occur, return nil to indicate success
	// This allows the caller to assume the directory was created successfully
	return nil
}
