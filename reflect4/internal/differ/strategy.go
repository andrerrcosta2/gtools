// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/internal/tracker"
	"github.com/andrerrcosta2/gtools/reflect4/op"
	"github.com/andrerrcosta2/gtools/reflect4/op/compare"
	"github.com/andrerrcosta2/gtools/reflect4/op/read"
	"reflect"
)

// NewStrategy creates a new differ strategy based on the options provided
func NewStrategy[O internal.Option](o ...O) *Strategy {
	s := defaultStrat()

outer:
	for _, opt := range o {
		switch t := any(opt).(type) {
		case op.Read:
			switch t {
			case read.SkipUnexportedFields:
				s.readField = readSettableField
				s.rideFields = rideSettableFields
			case read.SkipChannels:
				s.channels = skip
			case read.SkipFunctions:
				s.functions = skip
			default:
				break
			}
		case op.Compare:
			switch t {
			case compare.ChanIdentity:
				s.channels = sameChanDiff
			case compare.FuncIdentity:
				s.functions = sameFuncDiff
			case compare.PtrIdentity:
				s.pointers = samePtrDiff
			case compare.AllowNilVsEmpty:
				s.maps = serializableMapDiff
				s.slices = serializableSliceDiff
			case compare.ReflectSemantics:
				s = reflStrat() // this option doesn't accept override
				break outer
			case compare.Strict:
				s = strictStrat() // this option doesn't accept override
				break outer
			//case compare.IgnoreCase:
			//	s.strings = ignoreCaseStringDiff
			default:
				s.stringFormat |= t
			}
		}
	}
	return s
}

type Diff func(tab indent.Indentor, a, b reflect.Value, s *Strategy) differs.Difference

type Strategy struct {
	cache        func(a, b reflect.Value, diff differs.Difference) differs.Difference
	check        func(a, b reflect.Value) (diff differs.Difference, done bool)
	channels     Diff
	functions    Diff
	pointers     Diff
	maps         Diff
	slices       Diff
	strings      func(value, expected string) (string, string, int)
	readField    func(v reflect.Value, i int) (field reflect.Value, canRead bool)
	rideFields   func(a, b reflect.Value, fn functions.TriPredicate[int, reflect.Value, reflect.Value])
	stringFormat op.Compare
}

func defaultStrat() *Strategy {
	t := tracker.Diff()
	return &Strategy{
		check:        t.Get,
		cache:        t.Mark,
		channels:     defaultChanDiff,
		functions:    defaultFuncDiff,
		maps:         defaultMapDiff,
		pointers:     defaultPtrDiff,
		slices:       defaultSliceDiff,
		strings:      defaultStringDiff,
		readField:    readAnyField,
		rideFields:   rideAllFields,
		stringFormat: 0,
	}
}

func reflStrat() *Strategy {
	t := tracker.Diff()
	return &Strategy{
		check:        t.Get,
		cache:        t.Mark,
		channels:     sameChanDiff,
		functions:    sameFuncDiff,
		maps:         defaultMapDiff,
		pointers:     defaultPtrDiff,
		slices:       defaultSliceDiff,
		strings:      defaultStringDiff,
		readField:    readAnyField,
		rideFields:   rideAllFields,
		stringFormat: 0,
	}
}

func serStrat() *Strategy {
	t := tracker.Diff()
	return &Strategy{
		check:        t.Get,
		cache:        t.Mark,
		channels:     sameChanDiff,
		functions:    sameFuncDiff,
		maps:         defaultMapDiff,
		pointers:     defaultPtrDiff,
		slices:       defaultSliceDiff,
		strings:      defaultStringDiff,
		readField:    readAnyField,
		rideFields:   rideAllFields,
		stringFormat: 0,
	}
}

func strictStrat() *Strategy {
	t := tracker.Diff()
	return &Strategy{
		check:        t.Get,
		cache:        t.Mark,
		channels:     sameChanDiff,
		functions:    sameFuncDiff,
		maps:         defaultMapDiff,
		pointers:     samePtrDiff,
		slices:       defaultSliceDiff,
		strings:      defaultStringDiff,
		readField:    readAnyField,
		rideFields:   rideAllFields,
		stringFormat: 0,
	}
}
