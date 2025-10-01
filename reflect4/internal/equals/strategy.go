// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package equals

import (
	"github.com/andrerrcosta2/gtools/core/domain/functions"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/stringutil"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/internal/tracker"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"github.com/andrerrcosta2/gtools/reflect4/op"
	"github.com/andrerrcosta2/gtools/reflect4/op/compare"
	"github.com/andrerrcosta2/gtools/reflect4/op/read"
	"reflect"
)

func NewStrategy[O internal.Option](o ...O) *Strategy {
	differ := defaultStrat()

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
			default:
				break
			}
		case op.Compare:
			switch t {
			case compare.ChanIdentity:
				differ.channels = identityShallow
			case compare.FuncIdentity:
				differ.functions = identityShallow
			case compare.PtrIdentity:
				differ.pointers = identityDeep
			case compare.AllowNilVsEmpty:
				differ.channels = serialChan
				differ.maps = serialMap
				differ.slices = serialSlice
			case compare.ReflectSemantics:
				differ = reflStrat() // this option doesn't accept override
				break outer
			case compare.Strict:
				differ = strictStrat() // this option doesn't accept override
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
	unsafeptrs ShallowFunc

	rideFields   func(v reflect.Value, fn functions.BiPredicate[int, reflect.Value])
	stringFormat op.Compare
}

func defaultStrat() *Strategy {
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
		unsafeptrs: defaultUnsafePointer,
		rideFields: values.RideFields,
	}
}

func reflStrat() *Strategy {
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
		unsafeptrs: defaultUnsafePointer,
		rideFields: values.RideFields,
	}
}

func strictStrat() *Strategy {
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
		unsafeptrs: defaultUnsafePointer,
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
		unsafeptrs: defaultUnsafePointer,
		rideFields: values.RideFields,
	}
}
