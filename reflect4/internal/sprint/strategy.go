// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/internal/tracker"
	"reflect"
)

type DeepFunc func(tab indent.Tab, v reflect.Value, s *Strategy) string
type ShallowFunc func(tab indent.Tab, v reflect.Value) string

func NewStrategy(o ...internal.Option) *Strategy {
	s := defaultStrat()

	return s
}

type Strategy struct {
	Check     func(v reflect.Value) (string, bool)
	Mark      func(v reflect.Value, s string) string
	Array     DeepFunc
	Chan      ShallowFunc
	Func      ShallowFunc
	Interface DeepFunc
	Map       DeepFunc
	Ptr       DeepFunc
	Slice     DeepFunc
	Struct    DeepFunc
	Unsafe    ShallowFunc
}

func defaultStrat() *Strategy {
	t := tracker.Sprint()
	return &Strategy{
		Check:     t.Get,
		Mark:      t.Mark,
		Array:     defaultArray,
		Chan:      defaultChan,
		Func:      defaultFunc,
		Interface: defaultInterface,
		Map:       defaultMap,
		Ptr:       defaultPointer,
		Slice:     defaultSlice,
		Struct:    defaultStruct,
		Unsafe:    defaultUnsafe,
	}
}
