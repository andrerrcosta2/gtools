// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testseed

// ComplexUnsafeCastableBase is a base struct for unsafe casts
// Specifications:
//   - size 56 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableBase struct {
	Nested   NestedStruct   // Size 24 bytes. Offset 0
	ID       int            // Size 8 bytes. Offset 24
	Metadata map[string]any // Size 16 bytes. Offset 32
	Score    float64        // Size 8 bytes. Offset 48
}

// ComplexUnsafeCastableOf is a struct for unsafe casts
// Specifications:
//   - size 56 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableOf struct {
	ComplexUnsafeCastableBase
}

// ComplexUnsafeCastableA is a struct for unsafe casts
// Specifications:
//   - size 80 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableA struct {
	ComplexUnsafeCastableBase          // Size 56 bytes. Offset 0
	Idx                       uint     // Size 8 bytes. Offset 56
	List                      []string // Size 16 bytes. Offset 64
}

// ComplexUnsafeCastableB is a struct for unsafe casts
// Specifications:
//   - size 80 bytes
//   - alignment 8
//   - empty padding 0
type ComplexUnsafeCastableB struct {
	ComplexUnsafeCastableBase                // Size 56 bytes. Offset 0
	Score                     float64        // Size 8 bytes. Offset 56
	Metadata                  map[string]any // Size 16 bytes. Offset 64
}

// NestedStruct specifications:
//   - size 24 bytes
//   - alignment 8
//   - no padding
//
// A string in Go is composed of an address and a length, both of which are 8 bytes each.
// Therefore, the total size of the string field is 16 bytes, but it aligns to 8 bytes.
type NestedStruct struct {
	NestedField1 int    // Size 8 bytes. Offset 0
	NestedField2 string // Size 16 bytes. Offset 8
}
