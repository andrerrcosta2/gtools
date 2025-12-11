// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

import (
	"reflect"

	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/internal/tracker"
	"github.com/andrerrcosta2/gtools/reflect4/op"
	"github.com/andrerrcosta2/gtools/reflect4/op/clone"
	"github.com/andrerrcosta2/gtools/reflect4/op/read"
	"github.com/andrerrcosta2/gtools/reflect4/op/write"
)

// NewCopyStrategy creates a new CopyStrategy based on the given options parameters
func NewCopyStrategy[O internal.Option](o ...O) *CopyStrategy {
	s := DefaultCopyStrat()
	for _, opt := range o {
		switch t := any(opt).(type) {
		case op.Read:
			switch t {
			case read.SkipUnexportedFields:
				s.ReadField = readSettableField
				s.RideFields = rideSettableFields
			case read.SkipChannels:
				s.CopyChan = skipCopy
			case read.SkipFunctions:
				s.CopyFunc = skipCopy
			default:
				break
			}
		case op.Write:
			switch t {
			case write.SkipChannels:
				s.CopyChan = skipCopy
			case write.SkipFunctions:
				s.CopyFunc = skipCopy
			case write.SkipUnexportedFields:
				s.ReadField = readSettableField
				s.RideFields = rideSettableFields
			default:
				break
			}
		case op.Clone:
			switch t {
			case clone.SkipUnexportedFields:
				s.ReadField = readSettableField
				s.RideFields = rideSettableFields
			case clone.ChanIdentity:
				s.CopyChan = identityCopy
			case clone.FuncIdentity:
				s.CopyFunc = identityCopy
			case clone.PtrIdentity:
				s.CopyPtr = identityCopy
			case clone.SkipChan:
				s.CopyChan = skipCopy
			case clone.SkipFunc:
				s.CopyFunc = skipCopy
			default:
				break
			}
		default:
			break
		}
	}
	return s
}

type CopyStrategy struct {
	Cache      func(v, cache reflect.Value) reflect.Value
	Check      func(v reflect.Value) (reflect.Value, bool)
	CopyChan   func(v reflect.Value, _ *CopyStrategy) (reflect.Value, error)
	CopyFunc   func(v reflect.Value, _ *CopyStrategy) (reflect.Value, error)
	CopyPtr    func(v reflect.Value, s *CopyStrategy) (reflect.Value, error)
	ReadField  func(v reflect.Value, i int) (field reflect.Value, canRead bool)
	RideFields func(a, b reflect.Value, fn functions.TriPredicate[int, reflect.Value, reflect.Value])
}

func DefaultCopyStrat() *CopyStrategy {
	t := tracker.Reference()
	return &CopyStrategy{
		Cache:      t.Mark,
		Check:      t.Get,
		CopyChan:   defaultDeepCopyChan,
		CopyFunc:   defaultDeepCopyFunc,
		CopyPtr:    defaultDeepCopyPointer,
		ReadField:  readAllFields,
		RideFields: rideAllFields,
	}
}
