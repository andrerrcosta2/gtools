// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"reflect"

	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/internal/tracker"
	"github.com/andrerrcosta2/gtools/reflect4/op"
	"github.com/andrerrcosta2/gtools/reflect4/op/read"
)

type DeepFunc func(tab indent.Indentor, v reflect.Value, s *Strategy) string
type ShallowFunc func(tab indent.Indentor, v reflect.Value) string

func NewStrategy[O internal.Option](o ...O) *Strategy {
	s := DefaultStrat()
	for _, opt := range o {
		switch t := any(opt).(type) {
		case op.Read:
			switch t {
			case read.SkipUnexportedFields:
				s.Fields = exportedFields
			case read.SkipChannels:
				s.Chan = skipShallow
			case read.SkipFunctions:
				s.Func = skipShallow
			case read.SkipPtr:
				s.Ptr = skipDeep
			case read.SkipUnsafePtr:
				s.Unsafe = skipShallow
			default:
				break
			}
		}
	}
	return s
}

type Strategy struct {
	Check     func(v reflect.Value) (string, bool)
	Mark      func(v reflect.Value, s string) string
	Array     DeepFunc
	Chan      ShallowFunc
	Func      ShallowFunc
	Fields    DeepFunc
	Interface DeepFunc
	Map       DeepFunc
	Ptr       DeepFunc
	Slice     DeepFunc
	Struct    DeepFunc
	Unsafe    ShallowFunc
}

func DefaultStrat() *Strategy {
	t := tracker.Sprint()
	return &Strategy{
		Check:     t.Get,
		Mark:      t.Mark,
		Array:     defaultArray,
		Chan:      defaultChan,
		Func:      defaultFunc,
		Fields:    defaultFields,
		Interface: defaultInterface,
		Map:       defaultMap,
		Ptr:       defaultPointer,
		Slice:     defaultSlice,
		Struct:    defaultStruct,
		Unsafe:    defaultUnsafe,
	}
}
