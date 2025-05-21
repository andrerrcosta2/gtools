// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fsys

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Reader interface {
	fs.FS
	ReadDir(name string) ([]fs.DirEntry, error)
	ReadFile(name string) ([]byte, error)
}

// BuildPath constructs a file path based on the given path type and directory.
// It takes a pathType of type PathType and a dir of type string as input.
// The function returns the constructed path as a string and an error if any.
// The pathType parameter determines the type of path to be constructed.
// If pathType is Literal, the function returns the directory as-is.
// If pathType is Relative, the function constructs a path relative to the current working directory.
// If pathType is neither Literal, Relative an error is returned.
func BuildPath(pathType PathType, dir string) (string, error) {
	// Get the current working directory
	workDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current working directory: %v", err)
	}

	switch pathType {
	case Literal:
		return filepath.ToSlash(dir), nil
	case Relative:
		// Construct a path relative to the current working directory
		return filepath.ToSlash(filepath.Join(workDir, dir)), nil
	case ModRoot:
		modRoot, e := FindPathContainingFileRecursivelyBackward(workDir, "go.mod")
		if e != nil {
			return "", fmt.Errorf("error finding mod descriptor: %v", e)
		}
		modRoot = filepath.Clean(modRoot)
		if !strings.HasSuffix(modRoot, string(filepath.Separator)) {
			modRoot += string(filepath.Separator) // Append separator if not present
		}

		// Return full path
		return filepath.ToSlash(filepath.Join(modRoot, dir)), nil
	default:
		// Return an error if the pathType is unknown
		return "", fmt.Errorf("unknown PathType: %v", pathType)
	}
}

// FindPathContainingFileRecursivelyBackward searches for a file by name, starting
// from the given directory and traversing up through its parent directories until
// the file is found or the root directory is reached.
//
// Notes:
//   - This function does not resolve symbolic links. If `startDir` is a symlink,
//     the search will follow the symlink path as provided without dereferencing.
//   - On Linux/Unix systems, this function respects case sensitivity in file names.
//   - On Windows systems, it handles absolute paths and volume root detection,
//     including UNC paths.
//   - The function does not handle network-related errors or permissions issues
//     beyond reporting them via the returned error.
//
// Parameters:
//   - startDir: The directory to start the search from.
//   - filename: The name of the file to locate.
//
// Returns:
//   - The directory containing the specified file, or an error if not found.
func FindPathContainingFileRecursivelyBackward(startDir string, filename string) (string, error) {
	// Clean the start directory path
	startDir = filepath.Clean(startDir)

	for {
		// Construct the path to the file in the current directory
		goModPath := filepath.Join(startDir, filename)

		// Check if the file exists
		// Repeated calls to os.Stat in a deep directory tree might incur noticeable performance overhead in large projects.
		// If performance becomes an issue, caching intermediate directory results might help.
		if _, err := os.Stat(goModPath); !os.IsNotExist(err) {
			// If the file exists, return the directory path
			return startDir, nil
		}

		// Get the parent directory of the current directory
		parentDir := filepath.Dir(startDir)

		/// Check if we've reached the root of the file system or a UNC root
		if filepath.VolumeName(startDir) != "" && parentDir == startDir {
			// No further parent directories to check
			return "", fmt.Errorf("%s not found", filename)
		}

		// Update the start directory to the parent directory
		startDir = parentDir
	}
}

// ImportPath constructs a full filesystem path by combining the module directory and additional path components.
func ImportPath(modDir string, parts ...string) (string, error) {
	// Start with the module directory
	fullPath := modDir

	// Append additional path components
	for _, part := range parts {
		fullPath = path.Join(fullPath, part)
	}

	// Normalize the path: Replace backslashes with forward slashes
	normalizedPath := strings.ReplaceAll(fullPath, "\\", "/")

	// Validate the normalized path
	for i, r := range normalizedPath {
		if !isValidFileSystemChar(r) {
			return "", errors.New("invalid filesystem path: contains forbidden characters: " + fmt.Sprintf("'%c' at position %d", r, i))
		}
	}

	return normalizedPath, nil
}

// isValidFileSystemChar checks if a character is allowed in a filesystem path.
func isValidFileSystemChar(r rune) bool {
	return (r >= 'a' && r <= 'z') || // Lowercase letters
		(r >= 'A' && r <= 'Z') || // Uppercase letters
		(r >= '0' && r <= '9') || // Digits
		r == '/' || // Path separator
		//r == '\\' || // Backslash
		r == '.' || // Dot (used in filenames)
		r == '_' || // Underscore
		r == '-' || // Hyphen
		r == '!' || // Exclamation mark (allowed in Windows paths)
		r == '@' || // At symbol (used in Go module cache paths)
		r == ':' //
}
