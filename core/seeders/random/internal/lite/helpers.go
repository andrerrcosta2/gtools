// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package lite

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/prng"
	"github.com/andrerrcosta2/gtools/core/seeders/random/internal/reflectutils/randreflect"
	"github.com/andrerrcosta2/gtools/core/util/typeutil/charsets"
	"github.com/google/uuid"
	"math/rand"
	"reflect"
	"time"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

// RandAny generates a random value of any type except structs.
func RandAny() any {
	switch randreflect.Kind() {
	case reflect.Bool:
		return prng.Bool()
	case reflect.Int:
		return prng.Int()
	case reflect.Int8:
		return prng.Int8()
	case reflect.Int16:
		return prng.Int16()
	case reflect.Int32:
		return prng.Int32()
	case reflect.Int64:
		return prng.Int64()
	case reflect.Uint, reflect.Uintptr:
		return prng.Uint()
	case reflect.Uint8:
		return prng.Uint8()
	case reflect.Uint16:
		return prng.Uint16()
	case reflect.Uint32:
		return prng.Uint32()
	case reflect.Uint64:
		return prng.Uint64()
	case reflect.Float32:
		return prng.Float32()
	case reflect.Float64:
		return prng.Float64()
	case reflect.Complex64:
		return prng.Complex64()
	case reflect.Complex128:
		return prng.Complex128()
	case reflect.Array:
		// Generate a random value of the array element type
		typx := randreflect.Type()
		// Create a new array of the same type
		size := prng.Int(1, 10)
		arr := reflect.New(reflect.ArrayOf(size, typx)).Elem()
		for i := 0; i < size; i++ {
			val := randreflect.ValueOf(typx)
			arr.Index(i).Set(val)
		}
		return arr.Interface()
	case reflect.Chan:
		et := randreflect.Type()
		size := prng.Int(1, 10)

		// We can randomize directions, but we can only instantiate
		// "Bothdir" channels
		// The types will make a switch return a mess
		ch := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, et), size)

		go func() {
			for i := 0; i < size; i++ {
				val := randreflect.ValueOf(et)
				ch.Send(val)
			}
			ch.Close()
		}()

		return ch.Interface()

	case reflect.Map:
		kt := randreflect.CmpType()
		vt := randreflect.Type()
		size := prng.Int(1, 10)
		// Create an empty map
		mapValue := reflect.MakeMap(reflect.MapOf(kt, vt))
		// Generate a random key
		// Generate a random key
		for i := 0; i < size; i++ {
			key := randreflect.ValueOf(kt)
			// Generate a random value
			value := randreflect.ValueOf(vt)
			// Add the key-value pair to the map
			mapValue.SetMapIndex(key, value)
		}
		return mapValue.Interface()

	case reflect.Ptr:
		value := randreflect.AnyValue()
		ptr := reflect.New(value.Type())
		ptr.Elem().Set(value)
		return ptr.Interface()
	case reflect.Slice:
		et := randreflect.Type()
		size := prng.Int(1, 10)
		arr := reflect.MakeSlice(reflect.SliceOf(et), size, size)
		for i := 0; i < size; i++ {
			val := randreflect.ValueOf(et)
			arr.Index(i).Set(val)
		}
		return arr.Interface()

	case reflect.String:
		// TODO: This method lacks constraints control. Finish validation library
		// TODO: Add its interface to the core package
		return RandString(prng.Int(1, 50), charsets.AlphaNumeric)
	default:
		return nil
	}
}

// RandOf generates a single random value
func RandOf(t reflect.Type) any {
	if t == nil {
		return nil
	}
	kind := t.Kind()
	switch kind {
	case reflect.Bool:
		return prng.Bool()
	case reflect.Int:
		return prng.Int()
	case reflect.Int8:
		return prng.Int8()
	case reflect.Int16:
		return prng.Int16()
	case reflect.Int32:
		return prng.Int32()
	case reflect.Int64:
		return prng.Int64()
	case reflect.Uint, reflect.Uintptr:
		return prng.Uint()
	case reflect.Uint8:
		return prng.Uint8()
	case reflect.Uint16:
		return prng.Uint16()
	case reflect.Uint32:
		return prng.Uint32()
	case reflect.Uint64:
		return prng.Uint64()
	case reflect.Float32:
		return prng.Float32()
	case reflect.Float64:
		return prng.Float64()
	case reflect.Complex64:
		return prng.Complex64()
	case reflect.Complex128:
		return prng.Complex128()
	case reflect.Array:
		length := t.Len()
		elemType := t.Elem()
		arr := reflect.New(reflect.ArrayOf(length, elemType)).Elem()
		for i := 0; i < length; i++ {
			val := reflect.ValueOf(RandOf(elemType))
			arr.Index(i).Set(val)
		}
		return arr.Interface()
	case reflect.Chan:
		//fmt.Printf(">> Creating channel of type: %v\n", t)
		size := prng.Int(1, 10)
		ch := reflect.MakeChan(t, size)
		elemType := t.Elem()
		// Generate random values and send them to the channel
		go func() {
			for i := 0; i < size; i++ {
				val := reflect.ValueOf(RandOf(elemType))
				ch.Send(val)
			}
			ch.Close()
		}()
		return ch.Interface()
	case reflect.Map:
		// Create an empty map
		mapValue := reflect.MakeMap(t)
		size := prng.Int(1, 10)
		for i := 0; i < size; i++ {
			key := RandOf(t.Key())
			// Generate a random value
			value := RandOf(t.Elem())
			// Add the key-value pair to the map
			mapValue.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
		}
		return mapValue.Interface()

	case reflect.Ptr:
		elemType := t.Elem()
		value := RandOf(elemType)
		ptr := reflect.New(elemType)
		ptr.Elem().Set(reflect.ValueOf(value))
		return ptr.Interface()
	case reflect.Slice:
		size := prng.Int(1, 10)
		slice := reflect.MakeSlice(t, size, size)
		for i := 0; i < size; i++ {
			elem := reflect.ValueOf(RandOf(t.Elem()))
			slice.Index(i).Set(elem)
		}
		return slice.Interface()
	case reflect.Struct:
		return RandStructOf(t)
	case reflect.Interface:
		if t.NumMethod() == 0 {
			return RandAny()
		}
		return nil

	case reflect.String:
		// TODO: This method lacks constraints control. Finish validation library
		// TODO: Add its interface to the core package
		return RandString(prng.Int(1, 50), charsets.AlphaNumeric)
	default:
		panic(fmt.Sprintf("gtools:random: nil or invalid type: '%T' passed to be randomized", t))
	}
}

func RandCmp() any {
	switch randreflect.CmpKind() {
	case reflect.Bool:
		return prng.Bool()
	case reflect.Int:
		return prng.Int()
	case reflect.Int8:
		return prng.Int8()
	case reflect.Int16:
		return prng.Int16()
	case reflect.Int32:
		return prng.Int32()
	case reflect.Int64:
		return prng.Int64()
	case reflect.Uint, reflect.Uintptr:
		return prng.Uint()
	case reflect.Uint8:
		return prng.Uint8()
	case reflect.Uint16:
		return prng.Uint16()
	case reflect.Uint32:
		return prng.Uint32()
	case reflect.Uint64:
		return prng.Uint64()
	case reflect.String:
		return RandString(prng.Int(1, 50), charsets.AlphaNumeric)
	case reflect.Float32:
		return prng.Float32()
	case reflect.Float64:
		return prng.Float64()
	case reflect.Chan:
		return reflect.MakeChan(reflect.ChanOf(reflect.BothDir, reflect.TypeOf(prng.Int())), 0).Interface()
	case reflect.Ptr:
		return reflect.New(reflect.TypeOf(RandAny())).Interface()
	default:
		return nil
	}
}

// RandRune generates a random Unicode rune that is valid and graphic.
// It continues to generate runes until one satisfies these conditions.
func RandRune() rune {
	for {
		// Generate a random rune within the full Unicode range (up to 0x10FFFF)
		r := rune(prng.Uint(0, 0x10FFFF))
		// Check if the rune is valid and a graphic character (i.e. not a control character like '\n', '\t', etc.)
		if utf8.ValidRune(r) && unicode.IsGraphic(r) {
			return r
		}
	}
}

func RandString(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// RandStructOf generates a new instance of the given struct type with random values for its fields.
// It uses reflection to dynamically create an instance and set the fields.
func RandStructOf(t reflect.Type) any {
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
		fieldType := field.Type()

		if !field.CanSet() {
			addr := reflect.NewAt(fieldType, unsafe.Pointer(field.UnsafeAddr())).Elem()
			addr.Set(randreflect.ValueOf(fieldType))
			continue
		}

		// Generate a random value for the field using RandOf
		field.Set(randreflect.ValueOf(fieldType))
	}

	// Return the newly created struct instance with populated fields
	return v.Interface()
}

func RandTimestamp(from time.Time, diff time.Duration) time.Time {
	randomDuration := time.Duration(rand.Int63n(int64(diff)))
	return from.Add(randomDuration)
}

func RandUuid() string {
	return uuid.New().String()
}
