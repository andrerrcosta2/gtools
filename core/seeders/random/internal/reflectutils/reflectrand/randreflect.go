// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package reflectrand

import (
	"math/rand"
	"reflect"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/prng"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/reflectutils"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/charsets"
)

// AnyValue generates a random reflect.Value of random type
func AnyValue(maxDepth ...int) reflect.Value {
	t := Type()
	tt := reflectutils.TrackerOf(maxDepth...)
	switch t.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(prng.Bool())
	case reflect.Int:
		return reflect.ValueOf(prng.Int(-1, 1))
	case reflect.Int8:
		return reflect.ValueOf(prng.Int8(-1, 1))
	case reflect.Int16:
		return reflect.ValueOf(prng.Int16(-1, 1))
	case reflect.Int32:
		return reflect.ValueOf(prng.Int32(-1, 1))
	case reflect.Int64:
		return reflect.ValueOf(prng.Int64(-1, 1))
	case reflect.Uint:
		return reflect.ValueOf(prng.Uint(0, 1))
	case reflect.Uintptr:
		return reflect.ValueOf(uintptr(prng.Uint(0, 1)))
	case reflect.Uint8:
		return reflect.ValueOf(prng.Uint8(0, 1))
	case reflect.Uint16:
		return reflect.ValueOf(prng.Uint16(0, 1))
	case reflect.Uint32:
		return reflect.ValueOf(prng.Uint32(0, 1))
	case reflect.Uint64:
		return reflect.ValueOf(prng.Uint64(0, 1))
	case reflect.Float32:
		return reflect.ValueOf(prng.Float32(-1, 1))
	case reflect.Float64:
		return reflect.ValueOf(prng.Float64(-1, 1))
	case reflect.Complex64:
		return reflect.ValueOf(prng.Complex64(-1, 1, -1, 1))
	case reflect.Complex128:
		return reflect.ValueOf(prng.Complex128(-1, 1, -1, 1))
	case reflect.Array:
		return arrayOf(t, tt)
	case reflect.Chan:
		return ChanOf(t)
	case reflect.Func:
		return funcOf(t, tt)
	case reflect.Map:
		return mapOf(t, tt)
	case reflect.Ptr:
		return pointerOf(t, tt)
	case reflect.Slice:
		return sliceOf(t, tt)
	case reflect.Struct:
		return structOf(t, tt)
	case reflect.Interface:
		return interfaceOf(t, tt)
	case reflect.String:
		// TODO: This method lacks constraints control. Finish validation library
		// TODO: Add its interface to the core package
		return StringOf(prng.Int(1, 50), charsets.AlphaNumeric)
	case reflect.UnsafePointer:
		return UnsafeOf()
	default:
		panic(fmx.Sprintf("AnyValue: Invalid kind '%v' of type %s", t.Kind(), t.String()))
		return InvalidValue()
	}
}

// anyInterfaceValue generates a value for an interface, excluding another interface
func anyInterfaceValue(tt *reflectutils.Tracker) reflect.Value {
	t := interfaceType()
	switch t.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(prng.Bool())
	case reflect.Int:
		return reflect.ValueOf(prng.Int(-1, 1))
	case reflect.Int8:
		return reflect.ValueOf(prng.Int8(-1, 1))
	case reflect.Int16:
		return reflect.ValueOf(prng.Int16(-1, 1))
	case reflect.Int32:
		return reflect.ValueOf(prng.Int32(-1, 1))
	case reflect.Int64:
		return reflect.ValueOf(prng.Int64(-1, 1))
	case reflect.Uint:
		return reflect.ValueOf(prng.Uint(0, 1))
	case reflect.Uintptr:
		return reflect.ValueOf(uintptr(prng.Uint(0, 1)))
	case reflect.Uint8:
		return reflect.ValueOf(prng.Uint8(0, 1))
	case reflect.Uint16:
		return reflect.ValueOf(prng.Uint16(0, 1))
	case reflect.Uint32:
		return reflect.ValueOf(prng.Uint32(0, 1))
	case reflect.Uint64:
		return reflect.ValueOf(prng.Uint64(0, 1))
	case reflect.Float32:
		return reflect.ValueOf(prng.Float32(-1, 1))
	case reflect.Float64:
		return reflect.ValueOf(prng.Float64(-1, 1))
	case reflect.Complex64:
		return reflect.ValueOf(prng.Complex64(-1, 1, -1, 1))
	case reflect.Complex128:
		return reflect.ValueOf(prng.Complex128(-1, 1, -1, 1))
	case reflect.Array:
		return arrayOf(t, tt)
	case reflect.Chan:
		return ChanOf(t)
	case reflect.Func:
		return funcOf(t, tt)
	case reflect.Map:
		return mapOf(t, tt)
	case reflect.Ptr:
		return pointerOf(t, tt)
	case reflect.Slice:
		return sliceOf(t, tt)
	case reflect.Struct:
		return structOf(t, tt)
	case reflect.String:
		// TODO: This method lacks constraints control. Finish validation library
		// TODO: Add its interface to the core package
		return StringOf(prng.Int(1, 50), charsets.AlphaNumeric)
	case reflect.UnsafePointer:
		return UnsafeOf()
	default:
		panic(fmx.Sprintf("anyInterfaceValue: Invalid kind '%v' of type %s",
			t.Kind(), t.String()))
		return InvalidValue()
	}
}

// Array generates a random array type
func Array(size ...int) reflect.Type {
	if len(size) > 0 && size[0] > 0 {
		return reflect.ArrayOf(size[0], Type())
	}
	return reflect.ArrayOf(prng.Int(0, 10), Type())
}

// ArrayOf generates a random array by the given type.
func ArrayOf(t reflect.Type, maxDepth ...int) reflect.Value {
	return arrayOf(t, reflectutils.TrackerOf(maxDepth...))
}

func arrayOf(t reflect.Type, tt *reflectutils.Tracker) reflect.Value {
	length := t.Len()
	elemType := t.Elem()
	arr := reflect.New(t).Elem()
	for i := 0; i < length; i++ {
		arr.Index(i).Set(valueOf(elemType, tt))
	}
	return arr
}

// Chan generates a random channel type
func Chan() reflect.Type {
	return reflect.ChanOf(reflect.BothDir, Type())
}

// ChanOf generates a random channel by the given type.
func ChanOf(t reflect.Type) reflect.Value {
	size := prng.Int(1, 10)

	// always make from bidirectional
	bidir := t
	if t.ChanDir() != reflect.BothDir {
		bidir = reflect.ChanOf(reflect.BothDir, t.Elem())
	}

	// create the bidirectional channel
	ch := reflect.MakeChan(bidir, size)

	// if the requested type was uni-directional, convert back
	if t.ChanDir() != reflect.BothDir {
		ch = ch.Convert(t)
	}

	return ch
}

// CmpArray generates a random comparable array type
func CmpArray() reflect.Type {
	return reflect.ArrayOf(prng.Int(0, 10), Cmp())
}

// CmpKind returns a comparable random reflect.Kind
func CmpKind() reflect.Kind {
	return reflectutils.CmpKinds[prng.Int(0,
		len(reflectutils.CmpKinds)-1)]
}

// Cmp returns a comparable random reflect.Type
func Cmp() reflect.Type {
	kind := CmpKind()
	switch kind {
	case reflect.Bool:
		return reflect.TypeOf(false)
	case reflect.Int:
		return reflect.TypeOf(0)
	case reflect.Int8:
		return reflect.TypeOf(int8(0))
	case reflect.Int16:
		return reflect.TypeOf(int16(0))
	case reflect.Int32:
		return reflect.TypeOf(int32(0))
	case reflect.Int64:
		return reflect.TypeOf(int64(0))
	case reflect.Uint:
		return reflect.TypeOf(uint(0))
	case reflect.Uint8:
		return reflect.TypeOf(uint8(0))
	case reflect.Uint16:
		return reflect.TypeOf(uint16(0))
	case reflect.Uint32:
		return reflect.TypeOf(uint32(0))
	case reflect.Uint64:
		return reflect.TypeOf(uint64(0))
	case reflect.Uintptr:
		return reflect.TypeOf(uintptr(0))
	case reflect.Float32:
		return reflect.TypeOf(float32(0))
	case reflect.Float64:
		return reflect.TypeOf(0.0)
	case reflect.Complex64:
		return reflect.TypeOf(complex64(0))
	case reflect.Complex128:
		return reflect.TypeOf(complex128(0))
	case reflect.String:
		return String()
	case reflect.Array:
		return CmpArray()
	case reflect.Ptr:
		return Pointer()
	case reflect.UnsafePointer:
		return Unsafe()
	default:
		// Unreachable
		panic(fmx.Sprintf("Cmp: Invalid kind '%v'", kind))
		return InvalidType()
	}
}

// Func generates a random function reflect.Type
func Func() reflect.Type {
	in := make([]reflect.Type, prng.Int(0, 4))
	out := make([]reflect.Type, prng.Int(0, 3))
	for i := 0; i < len(in); i++ {
		in[i] = Type()
	}
	for i := 0; i < len(out); i++ {
		out[i] = Type()
	}
	return reflect.FuncOf(in, out, false)
}

// FuncOf generates a random function by the given type.
func FuncOf(t reflect.Type, maxDepth ...int) reflect.Value {
	return funcOf(t, reflectutils.TrackerOf(maxDepth...))
}

func funcOf(t reflect.Type, tt *reflectutils.Tracker) reflect.Value {
	return reflect.MakeFunc(t, func(args []reflect.Value) []reflect.Value {
		results := make([]reflect.Value, t.NumOut())
		for i := 0; i < t.NumOut(); i++ {
			outType := t.Out(i)
			results[i] = valueOf(outType, tt)
		}
		return results
	})
}

// Hashable returns a random hashable type that can be a value of an interface{}
func Hashable() reflect.Type {
	kind := reflectutils.HashableKinds.Rand()
	switch kind {
	case reflect.Bool:
		return reflect.TypeOf(false)
	case reflect.Int:
		return reflect.TypeOf(0)
	case reflect.Int8:
		return reflect.TypeOf(int8(0))
	case reflect.Int16:
		return reflect.TypeOf(int16(0))
	case reflect.Int32:
		return reflect.TypeOf(int32(0))
	case reflect.Int64:
		return reflect.TypeOf(int64(0))
	case reflect.Uint:
		return reflect.TypeOf(uint(0))
	case reflect.Uint8:
		return reflect.TypeOf(uint8(0))
	case reflect.Uint16:
		return reflect.TypeOf(uint16(0))
	case reflect.Uint32:
		return reflect.TypeOf(uint32(0))
	case reflect.Uint64:
		return reflect.TypeOf(uint64(0))
	case reflect.Uintptr:
		return reflect.TypeOf(uintptr(0))
	case reflect.Float32:
		return reflect.TypeOf(float32(0))
	case reflect.Float64:
		return reflect.TypeOf(0.0)
	case reflect.Complex64:
		return reflect.TypeOf(complex64(0))
	case reflect.Complex128:
		return reflect.TypeOf(complex128(0))
	case reflect.String:
		return String()
	case reflect.Array:
		return CmpArray()
	case reflect.Chan:
		return Chan()
	case reflect.Ptr:
		return Pointer()
	case reflect.UnsafePointer:
		return Unsafe()
	default:
		// Unreachable
		panic(fmx.Sprintf("interfaceType: Invalid kind '%v'", kind))
		return InvalidType()
	}
}

// hashableOf generates a random value for a hashable type
// Note> this method still accepts non hashable types, like
// non-comparable structs and arrays.
// this control must be done outside for performance.
func hashableOf(t reflect.Type, tt *reflectutils.Tracker) reflect.Value {
	if reflectutils.IsAnyInterface(t) {
		t = Hashable()
	}
	switch t.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(prng.Bool())
	case reflect.Int:
		return reflect.ValueOf(prng.Int(-1, 1))
	case reflect.Int8:
		return reflect.ValueOf(prng.Int8(-1, 1))
	case reflect.Int16:
		return reflect.ValueOf(prng.Int16(-1, 1))
	case reflect.Int32:
		return reflect.ValueOf(prng.Int32(-1, 1))
	case reflect.Int64:
		return reflect.ValueOf(prng.Int64(-1, 1))
	case reflect.Uint:
		return reflect.ValueOf(prng.Uint(0, 1))
	case reflect.Uintptr:
		return reflect.ValueOf(uintptr(prng.Uint(0, 1)))
	case reflect.Uint8:
		return reflect.ValueOf(prng.Uint8(0, 1))
	case reflect.Uint16:
		return reflect.ValueOf(prng.Uint16(0, 1))
	case reflect.Uint32:
		return reflect.ValueOf(prng.Uint32(0, 1))
	case reflect.Uint64:
		return reflect.ValueOf(prng.Uint64(0, 1))
	case reflect.Float32:
		return reflect.ValueOf(prng.Float32(-1, 1))
	case reflect.Float64:
		return reflect.ValueOf(prng.Float64(-1, 1))
	case reflect.Complex64:
		return reflect.ValueOf(prng.Complex64(-1, 1, -1, 1))
	case reflect.Complex128:
		return reflect.ValueOf(prng.Complex128(-1, 1, -1, 1))
	case reflect.Array:
		return arrayOf(t, tt)
	case reflect.Chan:
		return ChanOf(t)
	case reflect.Ptr:
		return pointerOf(t, tt)
	case reflect.Struct:
		return structOf(t, tt)
	case reflect.String:
		// TODO: This method lacks constraints control. Finish validation library
		// TODO: Add its interface to the core package
		return StringOf(prng.Int(1, 50), charsets.AlphaNumeric)
	case reflect.UnsafePointer:
		return UnsafeOf()
	default:
		panic(fmx.Sprintf("hashableOf: Invalid kind '%v' of type %s",
			t.Kind(), t.String()))
		return InvalidValue()
	}
}

// Interface generates a interface type
// this method only generates interfaces without methods
func Interface() reflect.Type {
	return reflect.TypeOf((*interface{})(nil)).Elem()
}

// InterfaceOf generates a random interface by the given type.
// If the interface has any method, it returns a zero value of the given type.
func InterfaceOf(t reflect.Type, maxDepth ...int) reflect.Value {
	return interfaceOf(t, reflectutils.TrackerOf(maxDepth...))
}

func interfaceOf(t reflect.Type, tt *reflectutils.Tracker) reflect.Value {
	if t.NumMethod() == 0 {
		return anyInterfaceValue(tt).Convert(t)
	}
	return reflect.Zero(t)
}

// interfaceType returns a random type that can be a value of an interface{}
func interfaceType() reflect.Type {
	kind := reflectutils.InterfaceKinds[prng.Int(0,
		len(reflectutils.InterfaceKinds)-1)]
	switch kind {
	case reflect.Bool:
		return reflect.TypeOf(false)
	case reflect.Int:
		return reflect.TypeOf(0)
	case reflect.Int8:
		return reflect.TypeOf(int8(0))
	case reflect.Int16:
		return reflect.TypeOf(int16(0))
	case reflect.Int32:
		return reflect.TypeOf(int32(0))
	case reflect.Int64:
		return reflect.TypeOf(int64(0))
	case reflect.Uint:
		return reflect.TypeOf(uint(0))
	case reflect.Uint8:
		return reflect.TypeOf(uint8(0))
	case reflect.Uint16:
		return reflect.TypeOf(uint16(0))
	case reflect.Uint32:
		return reflect.TypeOf(uint32(0))
	case reflect.Uint64:
		return reflect.TypeOf(uint64(0))
	case reflect.Uintptr:
		return reflect.TypeOf(uintptr(0))
	case reflect.Float32:
		return reflect.TypeOf(float32(0))
	case reflect.Float64:
		return reflect.TypeOf(0.0)
	case reflect.Complex64:
		return reflect.TypeOf(complex64(0))
	case reflect.Complex128:
		return reflect.TypeOf(complex128(0))
	case reflect.String:
		return String()
	case reflect.Array:
		return Array()
	case reflect.Chan:
		return Chan()
	case reflect.Func:
		return Func()
	case reflect.Map:
		return Map()
	case reflect.Ptr:
		return Pointer()
	case reflect.Slice:
		return Slice()
	case reflect.UnsafePointer:
		return Unsafe()
	default:
		// Unreachable
		panic(fmx.Sprintf("interfaceType: Invalid kind '%v'", kind))
		return InvalidType()
	}
}

// InvalidType returns an invalid reflect.Type
func InvalidType() reflect.Type {
	return nil
}

// InvalidValue returns an invalid reflect.Value
func InvalidValue() reflect.Value {
	return reflect.Value{}
}

func Kind() reflect.Kind {
	return reflectutils.RandomizableKinds.Rand()
}

// Map generates a random type of map
func Map() reflect.Type {
	return reflect.MapOf(Hashable(), Type())
}

// MapOf generates a random map by the given type
func MapOf(t reflect.Type, maxDepth ...int) reflect.Value {
	return mapOf(t, reflectutils.TrackerOf(maxDepth...))
}

func mapOf(t reflect.Type, tt *reflectutils.Tracker) reflect.Value {
	mapValue := reflect.MakeMap(t)
	size := prng.Int(1, 10)
	for i := 0; i < size; i++ {
		mapValue.SetMapIndex(hashableOf(t.Key(), tt), valueOf(t.Elem(), tt))
	}
	return mapValue
}

// Pointer generates a random pointer type
func Pointer() reflect.Type {
	return reflect.PointerTo(Type())
}

// PointerOf generates a random pointer by the given type
func PointerOf(t reflect.Type, maxDepth ...int) reflect.Value {
	return pointerOf(t, reflectutils.TrackerOf(maxDepth...))
}

func pointerOf(t reflect.Type, tt *reflectutils.Tracker) reflect.Value {
	elemType := t.Elem()
	ptr := reflect.New(elemType)

	val := valueOf(elemType, tt)

	// Named types won't be assignable
	if val.Type() != elemType {
		val = val.Convert(elemType)
	}
	ptr.Elem().Set(val)
	return ptr
}

// Slice generates a random slice type
func Slice() reflect.Type {
	return reflect.SliceOf(Type())
}

// SliceOf generates a random slice by the given type
func SliceOf(t reflect.Type, maxDepth ...int) reflect.Value {
	return sliceOf(t, reflectutils.TrackerOf(maxDepth...))
}

func sliceOf(t reflect.Type, tt *reflectutils.Tracker) reflect.Value {
	size := prng.Int(1, 10)
	slice := reflect.MakeSlice(t, size, size)
	for i := 0; i < size; i++ {
		slice.Index(i).Set(valueOf(t.Elem(), tt))
	}
	return slice
}

// String generates a string type
func String() reflect.Type {
	return reflect.TypeOf("")
}

// StringOf generates a random string with the given parameters and
// return it as a reflect.Value
func StringOf(length int, charset string) reflect.Value {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return reflect.ValueOf(string(b))
}

// StructOf generates a new instance of the given struct type with random values for its fields.
// It uses reflection to dynamically create an instance and set the fields.
func StructOf(t reflect.Type, maxDepth ...int) reflect.Value {
	return structOf(t, reflectutils.TrackerOf(maxDepth...))
}

// structOf generates a new instance of the given struct type with random values for
// its fields and a max depth for its inner values
func structOf(t reflect.Type, tt *reflectutils.Tracker) reflect.Value {
	v := reflect.New(t).Elem() // v is always addressable
	if !tt.Next(t) {
		return v
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		value := valueOf(field.Type(), tt)
		addr := reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
		addr.Set(value)
	}
	return v
}

// Type generates a random reflect type
func Type() reflect.Type {
	kind := Kind()
	switch kind {
	case reflect.Bool:
		return reflect.TypeOf(false)
	case reflect.Int:
		return reflect.TypeOf(0)
	case reflect.Int8:
		return reflect.TypeOf(int8(0))
	case reflect.Int16:
		return reflect.TypeOf(int16(0))
	case reflect.Int32:
		return reflect.TypeOf(int32(0))
	case reflect.Int64:
		return reflect.TypeOf(int64(0))
	case reflect.Uint:
		return reflect.TypeOf(uint(0))
	case reflect.Uint8:
		return reflect.TypeOf(uint8(0))
	case reflect.Uint16:
		return reflect.TypeOf(uint16(0))
	case reflect.Uint32:
		return reflect.TypeOf(uint32(0))
	case reflect.Uint64:
		return reflect.TypeOf(uint64(0))
	case reflect.Uintptr:
		return reflect.TypeOf(uintptr(0))
	case reflect.Float32:
		return reflect.TypeOf(float32(0))
	case reflect.Float64:
		return reflect.TypeOf(0.0)
	case reflect.Complex64:
		return reflect.TypeOf(complex64(0))
	case reflect.Complex128:
		return reflect.TypeOf(complex128(0))
	case reflect.String:
		return String()
	case reflect.Array:
		return Array()
	case reflect.Chan:
		return Chan()
	case reflect.Func:
		return Func()
	case reflect.Interface:
		return Interface()
	case reflect.Map:
		return Map()
	case reflect.Ptr:
		return Pointer()
	case reflect.Slice:
		return Slice()
	case reflect.UnsafePointer:
		return Unsafe()
	default:
		// Unreachable
		panic(fmx.Sprintf("Type: Invalid kind '%v'", kind))
		return InvalidType()
	}
}

// Unsafe returns an unsafe pointer type
func Unsafe() reflect.Type {
	return reflect.TypeOf(unsafe.Pointer(nil))
}

func UnsafeOf() reflect.Value {
	return reflect.ValueOf(unsafe.Pointer(nil))
}

func ValueOf(t reflect.Type, maxDepth ...int) reflect.Value {
	return valueOf(t, reflectutils.TrackerOf(maxDepth...))
}

// ValueOf generates a single random reflect.Value
func valueOf(t reflect.Type, tt *reflectutils.Tracker) reflect.Value {
	if t == nil {
		return reflect.Value{}
	}
	kind := t.Kind()
	switch kind {
	case reflect.Bool:
		return reflect.ValueOf(prng.Bool())
	case reflect.Int:
		return reflect.ValueOf(prng.Int(-1, 1))
	case reflect.Int8:
		return reflect.ValueOf(prng.Int8(-1, 1))
	case reflect.Int16:
		return reflect.ValueOf(prng.Int16(-1, 1))
	case reflect.Int32:
		return reflect.ValueOf(prng.Int32(-1, 1))
	case reflect.Int64:
		return reflect.ValueOf(prng.Int64(-1, 1))
	case reflect.Uint:
		return reflect.ValueOf(prng.Uint(0, 1))
	case reflect.Uintptr:
		return reflect.ValueOf(uintptr(prng.Uint(0, 1)))
	case reflect.Uint8:
		return reflect.ValueOf(prng.Uint8(0, 1))
	case reflect.Uint16:
		return reflect.ValueOf(prng.Uint16(0, 1))
	case reflect.Uint32:
		return reflect.ValueOf(prng.Uint32(0, 1))
	case reflect.Uint64:
		return reflect.ValueOf(prng.Uint64(0, 1))
	case reflect.Float32:
		return reflect.ValueOf(prng.Float32(-1, 1))
	case reflect.Float64:
		return reflect.ValueOf(prng.Float64(-1, 1))
	case reflect.Complex64:
		return reflect.ValueOf(prng.Complex64(-1, 1, -1, 1))
	case reflect.Complex128:
		return reflect.ValueOf(prng.Complex128(-1, 1, -1, 1))
	case reflect.Array:
		return arrayOf(t, tt)
	case reflect.Chan:
		return ChanOf(t)
	case reflect.Func:
		return funcOf(t, tt)
	case reflect.Map:
		return mapOf(t, tt)
	case reflect.Ptr:
		return pointerOf(t, tt)
	case reflect.Slice:
		return sliceOf(t, tt)
	case reflect.Struct:
		return structOf(t, tt)
	case reflect.Interface:
		return interfaceOf(t, tt)
	case reflect.String:
		// TODO: This method lacks constraints control. Finish validation library
		// TODO: Add its interface to the core package
		return StringOf(prng.Int(1, 50), charsets.AlphaNumeric)
	case reflect.UnsafePointer:
		return UnsafeOf()
	default:
		panic(fmx.Sprintf("ValueOf: Invalid kind '%v' of type %s", t.Kind(), t.String()))
		return InvalidValue()
	}
}

// ValueOf generates a single random reflect.Value of a comparable type
func valueOfCmp(t reflect.Type, tt *reflectutils.Tracker) reflect.Value {
	if t == nil {
		return reflect.Value{}
	}
	kind := t.Kind()
	switch kind {
	case reflect.Bool:
		return reflect.ValueOf(prng.Bool())
	case reflect.Int:
		return reflect.ValueOf(prng.Int(-1, 1))
	case reflect.Int8:
		return reflect.ValueOf(prng.Int8(-1, 1))
	case reflect.Int16:
		return reflect.ValueOf(prng.Int16(-1, 1))
	case reflect.Int32:
		return reflect.ValueOf(prng.Int32(-1, 1))
	case reflect.Int64:
		return reflect.ValueOf(prng.Int64(-1, 1))
	case reflect.Uint:
		return reflect.ValueOf(prng.Uint(0, 1))
	case reflect.Uintptr:
		return reflect.ValueOf(uintptr(prng.Uint(0, 1)))
	case reflect.Uint8:
		return reflect.ValueOf(prng.Uint8(0, 1))
	case reflect.Uint16:
		return reflect.ValueOf(prng.Uint16(0, 1))
	case reflect.Uint32:
		return reflect.ValueOf(prng.Uint32(0, 1))
	case reflect.Uint64:
		return reflect.ValueOf(prng.Uint64(0, 1))
	case reflect.Float32:
		return reflect.ValueOf(prng.Float32(-1, 1))
	case reflect.Float64:
		return reflect.ValueOf(prng.Float64(-1, 1))
	case reflect.Complex64:
		return reflect.ValueOf(prng.Complex64(-1, 1, -1, 1))
	case reflect.Complex128:
		return reflect.ValueOf(prng.Complex128(-1, 1, -1, 1))
	case reflect.Array:
		return arrayOf(t, tt)
	case reflect.Chan:
		return ChanOf(t)
	case reflect.Func:
		return funcOf(t, tt)
	case reflect.Map:
		return mapOf(t, tt)
	case reflect.Ptr:
		return pointerOf(t, tt)
	case reflect.Slice:
		return sliceOf(t, tt)
	case reflect.Struct:
		return structOf(t, tt)
	case reflect.Interface:
		return interfaceOf(t, tt)
	case reflect.String:
		// TODO: This method lacks constraints control. Finish validation library
		// TODO: Add its interface to the core package
		return StringOf(prng.Int(1, 50), charsets.AlphaNumeric)
	case reflect.UnsafePointer:
		return UnsafeOf()
	default:
		panic(fmx.Sprintf("ValueOf: Invalid kind '%v' of type %s", t.Kind(), t.String()))
		return InvalidValue()
	}
}
