// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fsutil

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

type PathType int

const (
	Relative PathType = iota
	Root
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
		dir = path.Join(workDir, dir)
	case Root:
		// Get the module path
		mod, err := modPath(workDir)
		if err != nil {
			return "", fmt.Errorf("error getting module path: %v", err)
		}
		// Construct a path relative to the module path
		dir = path.Join(mod, dir)
	default:
		// Return an error if the pathType is unknown
		return "", fmt.Errorf("unknown PathType: %v", pathType)
	}
	return dir, nil
}

// modPath searches for the directory containing the go.mod file starting from the given directory.
// It takes a startDir of type string as input and returns the directory path as a string and an error if any.
// If the go.mod file is found, the function returns the directory path.
// If the go.mod file isn't found, the function continues searching in the parent directories until the root directory is reached.
// If the go.mod file isn't found in any directory, an error is returned.
func modPath(startDir string) (string, error) {
	for {
		// Construct the path to the go.mod file in the current directory
		goModPath := filepath.Join(startDir, "go.mod")

		// Check if the go.mod file exists
		if _, err := os.Stat(goModPath); !os.IsNotExist(err) {
			// If the go.mod file exists, return the directory path
			return startDir, nil
		}

		// Get the parent directory of the current directory
		parentDir := filepath.Dir(startDir)

		// Check if the parent directory is the same as the current directory
		if parentDir == startDir {
			// If the parent directory is the same as the current directory, it means the go.mod file was not found in any directory
			return "", fmt.Errorf("go.mod not found")
		}

		// Update the start directory to the parent directory
		startDir = parentDir
	}
}
