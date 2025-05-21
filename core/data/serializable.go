// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

import (
	"encoding/json"
	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim"
)

type Serializable[S prim.Serializable] interface {
	Serial() S
}

type JSONSerializable interface {
	json.Marshaler
	json.Unmarshaler
}
