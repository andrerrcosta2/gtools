// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import (
	"reflect"

	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/internal/tracker"
	"github.com/andrerrcosta2/gtools/reflect4/op"
	"github.com/andrerrcosta2/gtools/reflect4/op/compare"
	"github.com/andrerrcosta2/gtools/reflect4/op/read"
)

// NewStrategy creates a new differ strategy based on the options provided
func NewStrategy[O internal.Option](o ...O) *Strategy {
	s := DefaultStrat()

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
			case read.SkipPtr:
				s.pointers = skip
			case read.SkipUnsafePtr:
				s.unsafe = skip
			default:
				break
			}
		case op.Compare:
			switch t {
			case compare.ChanIdentity:
				s.channels = sameChanDiff
				s.idChan = true
			case compare.FuncIdentity:
				s.functions = sameFuncDiff
				s.idFunc = true
			case compare.PtrIdentity:
				s.pointers = samePtrDiff
				s.idPtr = true
			case compare.AllowNilVsEmpty:
				s.maps = serializableMapDiff
				s.slices = serializableSliceDiff
			case compare.ReflectSemantics:
				s = ReflectionStrat() // this option doesn't accept override
				break outer
			case compare.Strict:
				s = StrictStrat() // this option doesn't accept override
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
	unsafe       Diff
	readField    func(v reflect.Value, i int) (field reflect.Value, canRead bool)
	rideFields   func(a, b reflect.Value, fn functions.TriPredicate[int, reflect.Value, reflect.Value])
	stringFormat op.Compare
	idChan       bool
	idFunc       bool
	idPtr        bool
}

func (s *Strategy) shouldDeepCompareKeys(keyType reflect.Type) bool {
	switch keyType.Kind() {
	case reflect.Chan:
		return !s.idChan
	case reflect.Func:
		return !s.idFunc
	case reflect.Ptr, reflect.UnsafePointer:
		return !s.idPtr
	case reflect.Interface:
		return true // always deep compare interface keys
	case reflect.Array:
		// check element type recursively
		return s.shouldDeepCompareKeys(keyType.Elem())
	case reflect.Struct:
		for i := 0; i < keyType.NumField(); i++ {
			if s.shouldDeepCompareKeys(keyType.Field(i).Type) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func DefaultStrat() *Strategy {
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
		unsafe:       defaultUnsafePtrs,
		readField:    readAnyField,
		rideFields:   rideAllFields,
		stringFormat: 0,
	}
}

func ReflectionStrat() *Strategy {
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
		unsafe:       defaultUnsafePtrs,
		readField:    readAnyField,
		rideFields:   rideAllFields,
		stringFormat: 0,
	}
}

func SerializableStrat() *Strategy {
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

func StrictStrat() *Strategy {
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
