// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package clone

import "github.com/andrerrcosta2/gtools/reflect4/op"

const (
	Default op.Clone = 0

	// PtrIdentity clones all ptrs addresses - shallow clone.
	PtrIdentity op.Clone = 1 << iota
	// ChanIdentity clones all channels underlying memory address - shallow clone.
	ChanIdentity
	// FuncIdentity clones all functions underlying memory address - shallow clone.
	FuncIdentity

	// SkipUnexportedFields skips copying unexported struct field4.
	SkipUnexportedFields
	// SkipChan skip copying channels
	SkipChan
	// SkipFunc skip copying functions
	SkipFunc
	// SkipPtr skip copying pointer4
	SkipPtr
	// SkipUnsafePtr skip copying unsafe pointer4
	SkipUnsafePtr
)
