// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fsys

import (
	"log"
	"strings"
	"testing"
)

// mod.go must be findable on a backward recursive search.
func TestBuildPath(t *testing.T) {
	// Construct paths for testing
	pkgDir := "core/io/fsys"

	tests := []struct {
		name      string
		pathType  PathType
		dir       string
		expected  string
		expectErr bool
	}{
		{
			name:     "Relative to pkg",
			pathType: Relative,
			dir:      "../",
			expected: "core/io",
		},
		{
			name:     "Relative to fsutil",
			pathType: Relative,
			dir:      "./",
			expected: pkgDir,
		},
		{
			name:      "Unknown PathType",
			pathType:  999, // Invalid PathType
			dir:       "any/dir",
			expected:  "",
			expectErr: true,
		},
		{
			name:     "Literal",
			pathType: Literal,
			dir:      "any/dir",
			expected: "any/dir",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := BuildPath(tt.pathType, tt.dir)
			if (err != nil) != tt.expectErr {
				t.Errorf("Error building path: \nGot: %v, \nWanted: %v", result, tt.expected)
				return
			}
			if !strings.HasSuffix(result, tt.expected) {
				t.Errorf("Error building path: \nGot: %v, \nWanted: %v", result, tt.expected)
			}
			log.Printf("Result: %v", result)
		})
	}
}
