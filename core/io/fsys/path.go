// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fsys

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type PathType int

const (
	Relative PathType = iota
	Literal
	ModRoot
)

type Path struct {
	Type PathType
	Path string
}

func PathOf(pathType PathType, path string) Path {
	return Path{Type: pathType, Path: path}
}

// IsValidPath checks if the given path is valid based on operating system-specific rules.
//
// Notes:
//   - On Windows, invalid characters include: < > : " / \ | ? *
//     Additionally, reserved filenames (e.g., CON, NUL) and paths exceeding the maximum length
//     (260 characters by default) are considered invalid.
//   - On UNIX-like systems, null characters ('\x00') and '/' in filenames are invalid.
//   - This function does not check if the path exists, only its validity.
//
// Parameters:
//   - path: The path string to validate.
//
// Returns:
//   - A boolean indicating whether the path is valid.
//
// IsValidPath validates individual path components for invalid characters, reserved names, and path traversal.
func IsValidPath(path string) bool {
	return ValidatePath(path) == nil
}

func ValidatePath(path string) error {
	if strings.TrimSpace(path) == "" {
		// Paths with only whitespace or empty strings are invalid
		return errors.New("path cannot be empty")
	}

	// Define invalid characters
	var invalidChars string
	if runtime.GOOS == "windows" {
		// Windows doesn't allow: < > : " / \ | ? *
		invalidChars = `<>:"|?*`

		// Additional checks for Windows
		// Check if the path length exceeds the maximum allowed
		if len(path) > 260 {
			return fmt.Errorf("path length exceeds the maximum allowed (260 characters): %s", path)
		}

		// Check for reserved names
		reservedNames := map[string]struct{}{
			"CON": {}, "PRN": {}, "AUX": {}, "NUL": {}, "COM1": {}, "COM2": {}, "COM3": {}, "COM4": {}, "COM5": {},
			"COM6": {}, "COM7": {}, "COM8": {}, "COM9": {}, "LPT1": {}, "LPT2": {}, "LPT3": {}, "LPT4": {}, "LPT5": {},
			"LPT6": {}, "LPT7": {}, "LPT8": {}, "LPT9": {},
		}
		baseName := strings.ToUpper(filepath.Base(path))
		if _, exists := reservedNames[baseName]; exists {
			return fmt.Errorf("reserved name: %s", path)
		}
	} else {
		// Common invalid characters for UNIX-like systems
		invalidChars = "/\x00"
	}

	// Check for invalid characters
	for _, ch := range path {
		if strings.ContainsRune(invalidChars, ch) {
			return fmt.Errorf("path '%s' contains an invalid character: %s", path, string(ch))
		}
	}

	// Prevent path traversal
	cleanedPath := filepath.Clean(path)
	if strings.HasPrefix(cleanedPath, "..") || strings.Contains(cleanedPath, "../") {
		return fmt.Errorf("path traversal detected: %s", path)
	}

	return nil
}

func SanitizePath(p string) string {
	// Replace backslashes with forward slashes
	p = strings.ReplaceAll(p, "\\", "/")

	// Remove version suffix (e.g., "@v0.0.0-...")
	if idx := strings.Index(p, "@"); idx != -1 {
		p = p[:idx]
	}

	// Remove invalid characters
	var sb strings.Builder
	for _, r := range p {
		if isValidImportChar(r) {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// isValidImportChar checks if a character is allowed in a Go import path.
func isValidImportChar(r rune) bool {
	return (r >= 'a' && r <= 'z') || // Lowercase letters
		(r >= 'A' && r <= 'Z') || // Uppercase letters
		(r >= '0' && r <= '9') || // Digits
		r == '/' || // Path separator
		r == '.' || // Dot (used in domain names)
		r == '_' || // Underscore
		r == '-' // Hyphen
}
