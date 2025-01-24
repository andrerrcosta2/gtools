// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

import "github.com/andrerrcosta2/gtools/core/gtools/constraints/prim"

type Serializable[S prim.Serializable] interface {
	Serial() S
}
