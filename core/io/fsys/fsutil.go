// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fsys

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

// IsValidPath checks if the given path is valid based on operating system-specific rules.
//
// Notes:
//   - On Windows, invalid characters include: < > : " / \ | ? *
//     Additionally, reserved filenames (e.g., CON, NUL) and paths exceeding the maximum length (260 characters by default) are considered invalid.
//   - On UNIX-like systems, null characters ('\x00') and '/' in filenames are invalid.
//   - This function does not check if the path exists, only its validity.
//
// Parameters:
//   - path: The path string to validate.
//
// Returns:
//   - A boolean indicating whether the path is valid.
func IsValidPath(path string) bool {
	if strings.TrimSpace(path) == "" {
		// Paths with only whitespace or empty strings are invalid
		return false
	}

	// Define invalid characters
	var invalidChars string
	if runtime.GOOS == "windows" {
		// Windows doesn't allow: < > : " / \ | ? *
		invalidChars = `<>:"/\|?*`

		// Additional checks for Windows
		// Check if the path length exceeds the maximum allowed
		if len(path) > 260 {
			return false
		}

		// Check for reserved names
		reservedNames := map[string]struct{}{
			"CON": {}, "PRN": {}, "AUX": {}, "NUL": {}, "COM1": {}, "COM2": {}, "COM3": {}, "COM4": {}, "COM5": {},
			"COM6": {}, "COM7": {}, "COM8": {}, "COM9": {}, "LPT1": {}, "LPT2": {}, "LPT3": {}, "LPT4": {}, "LPT5": {},
			"LPT6": {}, "LPT7": {}, "LPT8": {}, "LPT9": {},
		}
		baseName := strings.ToUpper(filepath.Base(path))
		if _, exists := reservedNames[baseName]; exists {
			return false
		}
	} else {
		// Common invalid characters for UNIX-like systems
		invalidChars = "/\x00"
	}

	// Check for invalid characters
	invalidCharSet := make(map[rune]struct{}, len(invalidChars))
	for _, ch := range invalidChars {
		invalidCharSet[ch] = struct{}{}
	}
	for _, ch := range path {
		if _, exists := invalidCharSet[ch]; exists {
			return false
		}
	}

	return true
}
