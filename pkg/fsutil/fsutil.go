// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fsutil

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// BuildPath constructs a file path based on the given path type and directory.
// It takes a pathType of type PathType and a dir of type string as input.
// The function returns the constructed path as a string and an error if any.
// The pathType parameter determines the type of path to be constructed.
// If pathType is Relative, the function constructs a path relative to the current working directory.
// If pathType is Root, the function constructs a path relative to the module path.
// If pathType is neither Relative nor Root, an error is returned.
func BuildPath(pathType PathType, dir string) (string, error) {
	// Get the current working directory
	workDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current working directory: %v", err)
	}
	switch pathType {
	case Relative:
		// Construct a path relative to the current working directory
		dir = filepath.ToSlash(filepath.Join(workDir, dir))
	case Root:
		// Get the module path
		mod, err := FindPathContainingFileRecursivelyBackward(workDir, "go.mod")
		if err != nil {
			return "", fmt.Errorf("error getting module path: %v", err)
		}
		// Construct a path relative to the module path
		dir = filepath.ToSlash(filepath.Join(mod, dir))
	default:
		// Return an error if the pathType is unknown
		return "", fmt.Errorf("unknown PathType: %v", pathType)
	}
	return dir, nil
}

// FindPathContainingFileRecursivelyBackward searches for a file recursively in the parent directories of the given start directory.
// It takes a string startDir  and a string filename as input and returns the directory path where the file is found (string) and an error if the file is not found (error).
// The function starts from the given start directory and recursively checks the parent directories until the file is found or there are no more parent directories.
func FindPathContainingFileRecursivelyBackward(startDir string, filename string) (string, error) {
	for {
		// Construct the path to the file in the current directory
		goModPath := filepath.Join(startDir, filename)

		// Check if the file exists
		if _, err := os.Stat(goModPath); !os.IsNotExist(err) {
			// If the file exists, return the directory path
			return startDir, nil
		}

		// Get the parent directory of the current directory
		parentDir := filepath.Dir(startDir)

		// Check if the parent directory is the same as the current directory
		if parentDir == startDir {
			// If the parent directory is the same as the current directory, it means the file was not found in any directory
			return "", fmt.Errorf("%s not found", filename)
		}

		// Update the start directory to the parent directory
		startDir = parentDir
	}
}

// IsValidPath checks if the given path is valid.
func IsValidPath(path string) bool {
	// TODO:
	var invalidChars string
	// Determine the invalid characters based on the operating system
	if runtime.GOOS == "windows" {
		// Windows doesn't allow: < > : " / \ | ? *
		invalidChars = `<>:"/\|?*`
	} else {
		// Common invalid characters on UNIX-like systems (e.g., Linux, macOS)
		// Generally, only null characters (0x00) are not allowed in filenames
		// However, we also avoid using control characters and '/' in filenames.
		invalidChars = "/\x00"
	}

	for _, ch := range path {
		if strings.ContainsRune(invalidChars, ch) {
			return false
		}
	}
	return true
}
