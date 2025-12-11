// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package write

import "github.com/andrerrcosta2/gtools/reflect4/op"

const (
	Default              op.Write = 0
	SkipUnexportedFields op.Write = 1 << iota
	SkipChannels
	SkipFunctions
	SkipPtr
	SkipUnsafePtr
)
