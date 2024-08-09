// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package fsutil

import (
	"log"
	"strings"
	"testing"
)

// mod.go must be findable on a backward recursive search.
func TestBuildPath(t *testing.T) {
	// Construct paths for testing
	pkgDir := "gtools/pkg"
	fsutil := "gtools/pkg/fsutil"

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
			expected: pkgDir,
		},
		{
			name:     "Relative to fsutil",
			pathType: Relative,
			dir:      "./",
			expected: fsutil,
		},
		{
			name:     "Root to pkg",
			pathType: Root,
			dir:      "./pkg",
			expected: pkgDir,
		},
		{
			name:     "Root to fsutil",
			pathType: Root,
			dir:      "./pkg/fsutil",
			expected: fsutil,
		},
		{
			name:      "Unknown PathType",
			pathType:  999, // Invalid PathType
			dir:       "any/dir",
			expected:  "",
			expectErr: true,
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
