// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package mod

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// CleanCache clears the Go module cache for the specified module.
func CleanCache(modulePath string) error {
	cmd := exec.Command("go", "clean", "-modcache", modulePath)
	return cmd.Run()
}

// Dependencies returns the list of dependencies of a Go module.
func Dependencies(modulePath string) ([]string, error) {
	cmd := exec.Command("go", "list", "-m", "all")
	cmd.Dir = modulePath // Run in the module's directory
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list dependencies: %w", err)
	}
	// Parse and return the list of dependencies
	lines := strings.Split(string(out), "\n")
	return lines, nil
}

// Exists checks if a Go module exists in the current environment.
func Exists(modulePath string) (bool, error) {
	_, err := ResolvePath(modulePath)
	if err != nil {
		return false, nil // Module does not exist
	}
	return true, nil
}

// DependencyGraph generates a dependency graph for the given module.
func DependencyGraph(modulePath string) (string, error) {
	cmd := exec.Command("go", "mod", "graph")
	cmd.Dir = modulePath
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to generate dependency graph: %w", err)
	}
	return string(out), nil
}

// ModuleMetadata represents metadata about a Go module.
type ModuleMetadata struct {
	Path    string
	Version string
	Replace string
}

// Metadata retrieves metadata about a Go module.
func Metadata(modulePath string) (*ModuleMetadata, error) {
	cmd := exec.Command("go", "list", "-m", "-json", modulePath)
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get metadata for module %s: %w", modulePath, err)
	}

	var metadata ModuleMetadata
	if err := json.Unmarshal(out, &metadata); err != nil {
		return nil, fmt.Errorf("failed to parse module metadata: %w", err)
	}
	return &metadata, nil
}

// ResolvePath returns the absolute path to the root directory of a Go module.
func ResolvePath(modulePath string) (string, error) {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}", modulePath)
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to resolve module %s: %w", modulePath, err)
	}
	return strings.TrimSpace(string(out)), nil
}

// ResolveRepl returns the replacement path for a module, if any.
func ResolveRepl(modulePath string) (string, error) {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Replace}}", modulePath)
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to resolve replacement path for module %s: %w", modulePath, err)
	}
	replacement := strings.TrimSpace(string(out))
	if replacement == "" {
		return "", nil // No replacement
	}
	return replacement, nil
}

// UpdateDependencies updates all dependencies to their latest versions.
func UpdateDependencies(modulePath string) error {
	cmd := exec.Command("go", "get", "-u", "./...")
	cmd.Dir = modulePath
	return cmd.Run()
}

// Validate checks if a directory contains a valid Go module.
func Validate(dir string) (bool, error) {
	cmd := exec.Command("go", "list", "-m")
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		return false, fmt.Errorf("failed to validate module in %s: %w", dir, err)
	}
	return true, nil
}

// Version returns the version of the specified module.
func Version(modulePath string) (string, error) {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Version}}", modulePath)
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get version for module %s: %w", modulePath, err)
	}
	return strings.TrimSpace(string(out)), nil
}
