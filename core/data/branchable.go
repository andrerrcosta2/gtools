// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

import (
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

type Branchable[B any] interface {
	Branch() (B, bool)
}

type SerializableBranchable[B any, S prim.Serializable] interface {
	Branchable[B]
	Serializable[S]
}

func IsBranchable(data any) bool {
	_, ok := data.(Branchable[any])
	return ok
}
