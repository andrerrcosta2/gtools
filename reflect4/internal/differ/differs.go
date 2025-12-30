// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"errors"
	"reflect"

	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
)

var (
	ErrUnexportedFieldStruct = errors.New("cannot diff between structs as unexported fields")
)

// Between returns the difference between two values
func Between[O internal.Option](value, expected reflect.Value, o ...O) (diff string, equals bool) {
	var tab indent.Branch
	d := between(tab, value, expected, NewStrategy(o...))
	if !d.Equals {
		return differs.Message(d.Message, d.Diff), false
	}
	return "", true
}

// BetweenByStrat returns the difference between two values using a predefined strategy
func BetweenByStrat(value, expected reflect.Value, strategy *Strategy) (diff string, equals bool) {
	var tab indent.Branch
	d := between(tab, value, expected, strategy)
	if !d.Equals {
		return differs.Message(d.Message, d.Diff), false
	}
	return "", true
}
