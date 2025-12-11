// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

// Package read provides options for controlling reflection read operations
package read

import "github.com/andrerrcosta2/gtools/reflect4/op"

const (
	// Default enables logical equivalence suitable for testing and diffing, rather than strict memory identity..
	//
	// Behavior:
	//   - Structs: Unexported fields are accessible; unaddressable values are made addressable.
	//   - channels: Two channels are compare if they have the same element type and buffer capacity.
	//   - functions: Two functions are compare if they share the same signature - parameter and result types.
	//   - Interfaces: Compares the underlying dynamic values deeply, not just the interface headers.
	//   - nil vs Empty: Treats nil slices/maps as different to empty ones - nil []T != []T{}.
	//
	// This mode is ideal for asserting semantic correctness in tests or config diffing.
	Default op.Read = 0
	// SkipUnexportedFields skips reading unexported struct fields
	SkipUnexportedFields op.Read = 1 << iota
	// SkipChannels skips reading operations over channels
	SkipChannels
	// SkipFunctions skips reading operations over functions
	SkipFunctions
	// SkipPtr skips reading operations over pointers
	SkipPtr
	// SkipUnsafePtr skips reading operations over unsafe pointers
	SkipUnsafePtr
)
