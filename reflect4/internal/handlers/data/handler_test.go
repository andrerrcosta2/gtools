// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package data

import (
	"fmt"
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/testlite/assertlite"
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/reflect4/internal/differ"
	"github.com/andrerrcosta2/gtools/reflect4/internal/values"
	"reflect"
	"testing"
	"unsafe"
)

func TestDeepCopy(t *testing.T) {

}

func TestDeepCopyStruct(t *testing.T) {
	values := gtests.Structs.Fuzz().Categories().Values().All()
	pointers := gtests.Structs.Fuzz().Categories().Refs().All()

	// Test 1 - Deep copy values
	t.Run("Deep copy values", func(t *testing.T) {
		values.Each(func(value any) {
			target := reflect.ValueOf(value)
			cp, err := DeepCopy(target)
			assertlite.NoError(t, err, "error while deep copying:\n%v\n", err)

			diff, eq, err := differ.Between(indent.Zero(), cp, reflect.ValueOf(value))
			assertlite.NoError(t, err, "error on differ between: %v\n", err)
			assertlite.True(t, eq, "deep copy failed\n%s", diff)
		})
	})

	// Test 2 - Deep copy pointers
	t.Run("Deep copy pointers", func(t *testing.T) {
		pointers.Each(func(pointer any) {
			if pointer == nil {
				panic("nil pointer")
			}
			cp, err := DeepCopy(reflect.ValueOf(pointer))
			assertlite.NoError(t, err, "error while deep copying:\n%v\n", err)

			diff, eq, err := differ.Between(indent.Zero(), cp, reflect.ValueOf(pointer))
			assertlite.NoError(t, err, "error on differ between: %v\n", err)
			assertlite.True(t, eq, "deep copy failed\n%s", diff)
		})
	})
}

func TestPlayground(t *testing.T) {
	// Get a fuzzed Account with references
	account := gtests.Structs.Fuzz().Values().Account()

	// Extract account.Profile.DateOfBirth.Location()
	profile := reflect.ValueOf(account).FieldByName("Profile")
	dateOfBirth := profile.FieldByName("DateOfBirth")
	loc := dateOfBirth.MethodByName("Location").Call(nil)[0] // returns *time.Location

	// Dereference the pointer to get the struct
	locStruct := loc.Elem()

	fmt.Printf("Struct Type: %s\n", locStruct.Type())
	fmt.Println("=== Field Addresses ===")
	idx := 0
	// Use your UnsafeFieldsAccess on time.Location
	err := values.UnsafeFieldsAccess(locStruct, func(field reflect.Value, ptr unsafe.Pointer) {
		fieldName := locStruct.Type().Field(idx).Name
		fmt.Printf("Field: %-15s | Type: %-15s | Address: %v\n",
			fieldName, field.Type(), ptr)
		idx++
	})
	if err != nil {
		return
	}
}
