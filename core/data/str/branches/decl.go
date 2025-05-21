// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package branches

import (
	"github.com/andrerrcosta2/gtools/core/data"
	"github.com/andrerrcosta2/gtools/core/data/str/nodes"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

type abstractBranch[B any] interface {
	data.Branchable[B]
	// New returns a new branch having this current as origin.
	New() B
}

type Of[B any, D any] interface {
	abstractBranch[B]
	// Data returns the data of the branch.
	Data() D
}

type OfNodes[B any, N any] interface {
	abstractBranch[B]
	// Nodes returns the nodes of the branch.
	Nodes() []N
}

type OfUnsortedNodes[N nodes.Typed[V], V any] OfNodes[OfUnsortedNodes[N, V], N]

type OfSerialized[B any, D any, S prim.Serializable] interface {
	Of[B, D]
	data.Serializable[S]
}

type SerialNodes[N nodes.Serializable[V, S], V any, S prim.Serializable, B OfNodes[B, N]] interface {
	OfNodes[B, N]
	data.Serializable[S]
}

type OfSequential[B any, D any] OfSerialized[B, D, int]
type OfSequentialNodes[V any] SerialNodes[nodes.Sequential[V], V, int, OfSequentialNodes[V]]

var _ data.SerializableBranchable[int, int] = (OfSequential[int, int])(nil)
