// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package randreflect

import (
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/prng"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/reflectutils"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/charsets"
	"math/rand"
	"reflect"
	"unsafe"
)

// AnyValue generates a random reflect.Value of random type
func AnyValue() reflect.Value {
	t := Type()
	switch t.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(prng.Bool())
	case reflect.Int:
		return reflect.ValueOf(prng.Int())
	case reflect.Int8:
		return reflect.ValueOf(prng.Int8())
	case reflect.Int16:
		return reflect.ValueOf(prng.Int16())
	case reflect.Int32:
		return reflect.ValueOf(prng.Int32())
	case reflect.Int64:
		return reflect.ValueOf(prng.Int64())
	case reflect.Uint:
		return reflect.ValueOf(prng.Uint())
	case reflect.Uintptr:
		return reflect.ValueOf(uintptr(prng.Uint()))
	case reflect.Uint8:
		return reflect.ValueOf(prng.Uint8())
	case reflect.Uint16:
		return reflect.ValueOf(prng.Uint16())
	case reflect.Uint32:
		return reflect.ValueOf(prng.Uint32())
	case reflect.Uint64:
		return reflect.ValueOf(prng.Uint64())
	case reflect.Float32:
		return reflect.ValueOf(prng.Float32())
	case reflect.Float64:
		return reflect.ValueOf(prng.Float64())
	case reflect.Complex64:
		return reflect.ValueOf(prng.Complex64())
	case reflect.Complex128:
		return reflect.ValueOf(prng.Complex128())
	case reflect.Array:
		return ArrayOf(t)
	case reflect.Chan:
		return ChanOf(t)
	case reflect.Func:
		return FuncOf(t)
	case reflect.Map:
		return MapOf(t)
	case reflect.Ptr:
		return PointerOf(t)
	case reflect.Slice:
		return SliceOf(t)
	case reflect.Struct:
		return StructOf(t)
	case reflect.Interface:
		return InterfaceOf(t)
	case reflect.String:
		// TODO: This method lacks constraints control. Finish validation library
		// TODO: Add its interface to the core package
		return StringOf(prng.Int(1, 50), charsets.AlphaNumeric)
	default:
		panic(fmx.Sprintf("AnyValue: Invalid kind '%v' of type %s", t.Kind(), t.String()))
		return InvalidValue()
	}
}

// anyInterfaceValue generates a value for an interface, excluding another interface
func anyInterfaceValue() reflect.Value {
	t := interfaceType()
	switch t.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(prng.Bool())
	case reflect.Int:
		return reflect.ValueOf(prng.Int())
	case reflect.Int8:
		return reflect.ValueOf(prng.Int8())
	case reflect.Int16:
		return reflect.ValueOf(prng.Int16())
	case reflect.Int32:
		return reflect.ValueOf(prng.Int32())
	case reflect.Int64:
		return reflect.ValueOf(prng.Int64())
	case reflect.Uint:
		return reflect.ValueOf(prng.Uint())
	case reflect.Uintptr:
		return reflect.ValueOf(uintptr(prng.Uint()))
	case reflect.Uint8:
		return reflect.ValueOf(prng.Uint8())
	case reflect.Uint16:
		return reflect.ValueOf(prng.Uint16())
	case reflect.Uint32:
		return reflect.ValueOf(prng.Uint32())
	case reflect.Uint64:
		return reflect.ValueOf(prng.Uint64())
	case reflect.Float32:
		return reflect.ValueOf(prng.Float32())
	case reflect.Float64:
		return reflect.ValueOf(prng.Float64())
	case reflect.Complex64:
		return reflect.ValueOf(prng.Complex64())
	case reflect.Complex128:
		return reflect.ValueOf(prng.Complex128())
	case reflect.Array:
		return ArrayOf(t)
	case reflect.Chan:
		return ChanOf(t)
	case reflect.Func:
		return FuncOf(t)
	case reflect.Map:
		return MapOf(t)
	case reflect.Ptr:
		return PointerOf(t)
	case reflect.Slice:
		return SliceOf(t)
	case reflect.Struct:
		return StructOf(t)
	case reflect.String:
		// TODO: This method lacks constraints control. Finish validation library
		// TODO: Add its interface to the core package
		return StringOf(prng.Int(1, 50), charsets.AlphaNumeric)
	default:
		panic(fmx.Sprintf("anyInterfaceValue: Invalid kind '%v' of type %s",
			t.Kind(), t.String()))
		return InvalidValue()
	}
}

// Array generates a random array type
func Array() reflect.Type {
	return reflect.ArrayOf(prng.Int(0, 10), Type())
}

// ArrayOf generates a random array by the given type.
func ArrayOf(t reflect.Type) reflect.Value {
	length := t.Len()
	elemType := t.Elem()
	arr := reflect.New(t).Elem()
	for i := 0; i < length; i++ {
		arr.Index(i).Set(ValueOf(elemType))
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
	ch := reflect.MakeChan(t, size)
	elemType := t.Elem()
	// Generate random values and send them to the channel
	go func() {
		for i := 0; i < size; i++ {
			ch.Send(ValueOf(elemType))
		}
		ch.Close()
	}()
	return ch
}

// CmpArray generates a random comparable array type
func CmpArray() reflect.Type {
	return reflect.ArrayOf(prng.Int(0, 10), CmpType())
}

// CmpKind returns a comparable random reflect.Kind
func CmpKind() reflect.Kind {
	return reflectutils.CmpKinds[prng.Int(0,
		len(reflectutils.CmpKinds)-1)]
}

// CmpType returns a comparable random reflect.Type
func CmpType() reflect.Type {
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
	default:
		// Unreachable
		panic(fmx.Sprintf("CmpType: Invalid kind '%v'", kind))
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
func FuncOf(t reflect.Type) reflect.Value {
	return reflect.MakeFunc(t, func(args []reflect.Value) []reflect.Value {
		results := make([]reflect.Value, t.NumOut())
		for i := 0; i < t.NumOut(); i++ {
			outType := t.Out(i)
			results[i] = ValueOf(outType)
		}
		return results
	})
}

// Interface generates a interface type
// this method only generates interfaces without methods
func Interface() reflect.Type {
	return reflect.TypeOf((*interface{})(nil)).Elem()
}

// InterfaceOf generates a random interface by the given type.
// If the interface has any method, it returns a zero value of the given type.
func InterfaceOf(t reflect.Type) reflect.Value {
	if t.NumMethod() == 0 {
		return anyInterfaceValue().Convert(t)
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
	return reflectutils.RandomizableKinds[prng.Int(0,
		len(reflectutils.RandomizableKinds)-1)]
}

// Map generates a random type of map
func Map() reflect.Type {
	return reflect.MapOf(CmpType(), Type())
}

// MapOf generates a random map by the given type
func MapOf(t reflect.Type) reflect.Value {
	// Create an empty map
	mapValue := reflect.MakeMap(t)
	size := prng.Int(1, 10)
	for i := 0; i < size; i++ {
		// Add the key-value pair to the map
		mapValue.SetMapIndex(ValueOf(t.Key()), ValueOf(t.Elem()))
	}
	return mapValue
}

// Pointer generates a random pointer type
func Pointer() reflect.Type {
	return reflect.PointerTo(Type())
}

// PointerOf generates a random pointer by the given type
func PointerOf(t reflect.Type) reflect.Value {
	elemType := t.Elem()
	ptr := reflect.New(elemType)
	ptr.Elem().Set(ValueOf(elemType))
	return ptr
}

// Slice generates a random slice type
func Slice() reflect.Type {
	return reflect.SliceOf(Type())
}

// SliceOf generates a random slice by the given type
func SliceOf(t reflect.Type) reflect.Value {
	size := prng.Int(1, 10)
	slice := reflect.MakeSlice(t, size, size)
	for i := 0; i < size; i++ {
		slice.Index(i).Set(ValueOf(t.Elem()))
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
func StructOf(t reflect.Type) reflect.Value {
	// Create a new instance of the struct
	v := reflect.New(t).Elem()

	if !v.CanAddr() {
		ptr := reflect.New(v.Type())
		ptr.Elem().Set(v)
		v = ptr.Elem()
	}

	// Iterate through each field of the struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		value := ValueOf(field.Type())

		if value.Kind() != field.Type().Kind() {
			panic(fmx.Sprintf("%s: wrong kind generated '%v' for field of '%v'",
				t.String(), value.String(), field.Type().String()))
		}

		if !field.CanSet() {
			addr := reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
			addr.Set(value)
			continue
		}

		// Generate a random value for the field using RandOf
		field.Set(value)
	}

	// Return the newly created struct instance with populated fields
	return v
}

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
	default:
		// Unreachable
		panic(fmx.Sprintf("Type: Invalid kind '%v'", kind))
		return InvalidType()
	}
}

// ValueOf generates a single random reflect.Value
func ValueOf(t reflect.Type) reflect.Value {
	if t == nil {
		return reflect.Value{}
	}
	kind := t.Kind()
	switch kind {
	case reflect.Bool:
		return reflect.ValueOf(prng.Bool())
	case reflect.Int:
		return reflect.ValueOf(prng.Int())
	case reflect.Int8:
		return reflect.ValueOf(prng.Int8())
	case reflect.Int16:
		return reflect.ValueOf(prng.Int16())
	case reflect.Int32:
		return reflect.ValueOf(prng.Int32())
	case reflect.Int64:
		return reflect.ValueOf(prng.Int64())
	case reflect.Uint:
		return reflect.ValueOf(prng.Uint())
	case reflect.Uintptr:
		return reflect.ValueOf(uintptr(prng.Uint()))
	case reflect.Uint8:
		return reflect.ValueOf(prng.Uint8())
	case reflect.Uint16:
		return reflect.ValueOf(prng.Uint16())
	case reflect.Uint32:
		return reflect.ValueOf(prng.Uint32())
	case reflect.Uint64:
		return reflect.ValueOf(prng.Uint64())
	case reflect.Float32:
		return reflect.ValueOf(prng.Float32())
	case reflect.Float64:
		return reflect.ValueOf(prng.Float64())
	case reflect.Complex64:
		return reflect.ValueOf(prng.Complex64())
	case reflect.Complex128:
		return reflect.ValueOf(prng.Complex128())
	case reflect.Array:
		return ArrayOf(t)
	case reflect.Chan:
		return ChanOf(t)
	case reflect.Func:
		return FuncOf(t)
	case reflect.Map:
		return MapOf(t)
	case reflect.Ptr:
		return PointerOf(t)
	case reflect.Slice:
		return SliceOf(t)
	case reflect.Struct:
		return StructOf(t)
	case reflect.Interface:
		return InterfaceOf(t)
	case reflect.String:
		// TODO: This method lacks constraints control. Finish validation library
		// TODO: Add its interface to the core package
		return StringOf(prng.Int(1, 50), charsets.AlphaNumeric)
	default:
		panic(fmx.Sprintf("ValueOf: Invalid kind '%v' of type %s", t.Kind(), t.String()))
		return InvalidValue()
	}
}
