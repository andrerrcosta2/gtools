// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package equals

import (
	"reflect"

	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/stringutil"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/internal/tracker"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"github.com/andrerrcosta2/gtools/reflect4/op"
	"github.com/andrerrcosta2/gtools/reflect4/op/compare"
	"github.com/andrerrcosta2/gtools/reflect4/op/read"
)

func NewStrategy[O internal.Option](o ...O) *Strategy {
	differ := DefaultStrat()

outer:
	for _, opt := range o {
		switch t := any(opt).(type) {
		case op.Read:
			switch t {
			case read.SkipUnexportedFields:
				differ.rideFields = values.RideExportedFields
			case read.SkipChannels:
				differ.channels = shallowSkip
			case read.SkipFunctions:
				differ.functions = shallowSkip
			case read.SkipPtr:
				differ.pointers = deepSkip
			case read.SkipUnsafePtr:
				differ.unsafe = shallowSkip
			default:
				break
			}
		case op.Compare:
			switch t {
			case compare.ChanIdentity:
				differ.channels = identityShallow
				differ.idChan = true
			case compare.FuncIdentity:
				differ.functions = identityShallow
				differ.idFunc = true
			case compare.PtrIdentity:
				differ.pointers = identityDeep
				differ.idPtr = true
			case compare.AllowNilVsEmpty:
				differ.channels = serialChan
				differ.maps = serialMap
				differ.slices = serialSlice
			case compare.ReflectSemantics:
				differ = ReflStrat() // this option doesn't accept override
				break outer
			case compare.Strict:
				differ = StrictStrat() // this option doesn't accept override
				break outer
			case compare.IgnoreCase:
				differ.strings = stringutil.EqualsIgnoreCase
			default:
				differ.stringFormat |= t
			}
		}
	}
	return differ
}

type DeepFunc func(a, b reflect.Value, s *Strategy) bool
type ShallowFunc func(a, b reflect.Value) bool

type Strategy struct {
	check      func(a, b reflect.Value) bool
	mark       func(a, b reflect.Value)
	arrays     DeepFunc
	channels   ShallowFunc
	functions  ShallowFunc
	interfaces DeepFunc
	pointers   DeepFunc
	maps       DeepFunc
	slices     DeepFunc
	strings    func(a, b string) bool
	structs    DeepFunc
	unsafe     ShallowFunc

	rideFields   func(v reflect.Value, fn functions.BiPredicate[int, reflect.Value])
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
	t := tracker.Eq()
	return &Strategy{
		check:      t.Check,
		mark:       t.Mark,
		arrays:     defaultArray,
		channels:   defaultChan,
		functions:  defaultFunc,
		interfaces: defaultInterface,
		pointers:   defaultPtr,
		maps:       defaultMap,
		slices:     defaultSlice,
		strings:    stringutil.Equals,
		structs:    defaultStruct,
		unsafe:     defaultUnsafePointer,
		rideFields: values.RideFields,
	}
}

func ReflStrat() *Strategy {
	t := tracker.Eq()
	return &Strategy{
		check:      t.Check,
		mark:       t.Mark,
		arrays:     defaultArray,
		channels:   nativeChan,
		functions:  nativeFunc,
		interfaces: defaultInterface,
		pointers:   defaultPtr,
		maps:       defaultMap,
		slices:     defaultSlice,
		strings:    stringutil.Equals,
		structs:    defaultStruct,
		unsafe:     defaultUnsafePointer,
		rideFields: values.RideFields,
	}
}

func StrictStrat() *Strategy {
	t := tracker.Eq()
	return &Strategy{
		check:      t.Check,
		mark:       t.Mark,
		arrays:     defaultArray,
		channels:   identityShallow,
		functions:  identityShallow,
		interfaces: defaultInterface,
		pointers:   identityDeep,
		maps:       identityDeep,
		slices:     identitySlice,
		strings:    stringutil.Equals,
		structs:    defaultStruct,
		unsafe:     defaultUnsafePointer,
		rideFields: values.RideFields,
	}
}

func serialStrat() *Strategy {
	t := tracker.Eq()
	return &Strategy{
		check:      t.Check,
		mark:       t.Mark,
		arrays:     defaultArray,
		channels:   serialChan,
		functions:  defaultFunc,
		interfaces: defaultInterface,
		pointers:   defaultPtr,
		maps:       serialMap,
		slices:     serialSlice,
		strings:    stringutil.Equals,
		structs:    defaultStruct,
		unsafe:     defaultUnsafePointer,
		rideFields: values.RideFields,
	}
}
