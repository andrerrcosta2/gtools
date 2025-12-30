// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprint

import (
	"reflect"
	"regexp"
	"testing"
	"unsafe"

	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/differs"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"github.com/andrerrcosta2/gtools/core/format/sprints"
	"github.com/andrerrcosta2/gtools/gtests"
	"github.com/andrerrcosta2/gtools/gtests/testingseeds/static/structs/models"
	"github.com/andrerrcosta2/gtools/gtests/testingtools"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/themes"
	"github.com/andrerrcosta2/gtools/reflect4/internal"
	"github.com/andrerrcosta2/gtools/reflect4/internal/_testdata"
)

var (
	cyclicRefRE = regexp.MustCompile(`<cyclic-ref\|0x[0-9a-f]+>`)
	addrRE      = regexp.MustCompile(`<0x[0-9a-f]+>`)
)

// TestValue_BasicTypes tests sprinting of basic type
func TestValue_BasicTypes(t *testing.T) {
	zero := indent.Zero()
	tests := []struct {
		name     string
		input    reflect.Value
		expected string
	}{
		// Primitive Types
		{
			name:     "Test int",
			input:    reflect.ValueOf(42),
			expected: sprints.TypedDigit("int", 42),
		},
		{
			name:     "Test float64",
			input:    reflect.ValueOf(3.14),
			expected: sprints.TypedRoundFloat("float64", 3.14),
		},
		{
			name:     "Test string",
			input:    reflect.ValueOf("hello"), // reflect.ValueOf("hello",
			expected: sprints.TypedString("hello"),
		},

		// StructsPtr
		{
			name: "Test struct{ Name string, Age  }",
			input: reflect.ValueOf(struct {
				Name string
				Age  int
			}{
				Name: "Alice",
				Age:  30,
			}),
			expected: sprints.BClosedobj(zero, sprints.Anonymous(zero, "struct",
				sprints.KeyValue(zero, "Name", "string"),
				sprints.KeyValue(zero, "Age", "int")),
				sprints.Field(zero, "Name", sprints.TypedString("Alice")),
				sprints.Field(zero, "Age", sprints.Typed("int", 30)),
			),
		},

		// Slices
		{
			name:  "Test slice of ints",
			input: reflect.ValueOf([]int{1, 2, 3}),
			expected: sprints.ClosedSlice(zero, "int",
				sprints.TypedDigit("int", 1),
				sprints.TypedDigit("int", 2),
				sprints.TypedDigit("int", 3),
			),
		},
		{
			name:     "Test nil slice",
			input:    reflect.ValueOf([]int(nil)),
			expected: sprints.NilType(zero, "[]int"),
		},
		{
			name:  "Test array of strings",
			input: reflect.ValueOf([3]string{"a", "b", "c"}),
			expected: sprints.ClosedArray(zero, "string", 3,
				zero.Inc().Sprint(sprints.TypedString("a")),
				zero.Inc().Sprint(sprints.TypedString("b")),
				zero.Inc().Sprint(sprints.TypedString("c")),
			),
		},

		// maps
		{
			name: "Test map[string]int",
			input: reflect.ValueOf(map[string]int{
				"one": 1,
				"two": 2,
			}),
			expected: sprints.ClosedMap(zero, "string", "int",
				sprints.Field(zero, sprints.TypedString("one"), sprints.TypedDigit("int", 1)),
				sprints.Field(zero, sprints.TypedString("two"), sprints.TypedDigit("int", 2)),
			),
		},
		{
			name:     "Test nil map",
			input:    reflect.ValueOf((map[string]int)(nil)),
			expected: sprints.NilType(zero, "map[string]int"),
		},

		// channels
		{
			name:     "Test bidirectional channel",
			input:    reflect.ValueOf(make(chan int)),
			expected: "chan int",
		},
		{
			name:     "Test receive-only channel",
			input:    reflect.ValueOf(make(<-chan int)),
			expected: "<-chan int",
		},
		{
			name:     "Test send-only channel",
			input:    reflect.ValueOf(make(chan<- int)),
			expected: "chan<- int",
		},
		{
			name:     "Test nil channel",
			input:    reflect.ValueOf((chan int)(nil)),
			expected: sprints.NilType(zero, "chan int"),
		},

		// functions
		{
			name:     "Test func(int) int",
			input:    reflect.ValueOf(func(x int) int { return x + 1 }),
			expected: "func(int) int",
		},

		// PointerValues and Interfaces
		{
			name: "Test pointer to int",
			input: reflect.ValueOf(func() *int {
				x := 42
				return &x
			}()),
			expected: "*<int>42",
		},
		{
			name:     "Test nil pointer *int",
			input:    reflect.ValueOf((*int)(nil)),
			expected: sprints.NilType(zero, "*int"),
		},

		// Invalid Values
		{
			name:     "Test invalid value",
			input:    reflect.Value{},
			expected: "invalid type: <invalid>",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
			tt.StackTitle(tc.name, tc.expected)
			result := Of[internal.Option](zero, tc.input)
			compareMessages(tt, result, tc.expected)
			tt.StackLn()
		})
	}
}

func TestValue_ComplexTypes(t *testing.T) {
	zero := indent.Zero()
	tests := []struct {
		name     string
		input    reflect.Value
		expected string
	}{
		// Nested StructsPtr
		{
			name: "Test nested struct",
			input: reflect.ValueOf(struct {
				Name  string
				Age   int
				Child struct {
					Name string
					Age  int
				}
			}{
				Name: "Alice",
				Age:  30,
				Child: struct {
					Name string
					Age  int
				}{
					Name: "Bob",
					Age:  5,
				},
			}),
			expected: sprints.BClosedobj(zero, sprints.Anonymous(zero, "struct",
				sprints.KeyValue(zero, "Name", "string"),
				sprints.KeyValue(zero, "Age", "int"),
				sprints.KeyValue(zero, "Child", sprints.Anonymous(zero, "struct",
					sprints.KeyValue(zero, "Name", "string"), // tab doesn't matter here (no close)
					sprints.KeyValue(zero, "Age", "int")))),  // tab doesn't matter here (no close)
				sprints.Field(zero, "Name", sprints.TypedString("Alice")), // tab doesn't matter here (no close)
				sprints.Field(zero, "Age", sprints.Typed("int", 30)),      // tab doesn't matter here
				sprints.Field(zero, "Child", sprints.BClosedobj(zero, sprints.Anonymous(zero, "struct",
					sprints.KeyValue(zero, "Name", "string"),
					sprints.KeyValue(zero, "Age", "int")),
					sprints.Field(zero, "Name", sprints.TypedString("Bob")),
					sprints.Field(zero, "Age", sprints.Typed("int", 5)),
				)),
			),
		},

		// Slices of maps
		{
			name: "Test slice of maps",
			input: reflect.ValueOf([]map[string]int{
				{"one": 1, "two": 2},
				{"three": 3, "four": 4},
			}),
			expected: sprints.ClosedSlice(zero, "map[string]int",
				sprints.ClosedMap(zero.Inc(), "string", "int",
					sprints.Field(zero.Inc(), sprints.TypedString("one"), sprints.TypedDigit("int", 1)),
					sprints.Field(zero.Inc(), sprints.TypedString("two"), sprints.TypedDigit("int", 2)),
				),
				sprints.ClosedMap(zero.Inc(), "string", "int",
					sprints.Field(zero.Inc(), sprints.TypedString("four"), sprints.TypedDigit("int", 4)),
					sprints.Field(zero.Inc(), sprints.TypedString("three"), sprints.TypedDigit("int", 3)),
				),
			),
		},

		// maps with Slicef Values
		{
			name: "Test map with slice values",
			input: reflect.ValueOf(map[string][]int{
				"odd":  {1, 3, 5},
				"even": {2, 4, 6},
			}),
			expected: sprints.ClosedMap(zero, "string", "[]int",
				sprints.Field(zero, sprints.TypedString("even"),
					sprints.ClosedSlice(zero.Inc(), "int",
						sprints.TypedDigit("int", 2),
						sprints.TypedDigit("int", 4),
						sprints.TypedDigit("int", 6),
					),
				),
				sprints.Field(zero, sprints.TypedString("odd"),
					sprints.ClosedSlice(zero.Inc(), "int",
						sprints.TypedDigit("int", 1),
						sprints.TypedDigit("int", 3),
						sprints.TypedDigit("int", 5),
					),
				),
			),
		},

		// Arrays of StructsPtr
		{
			name: "Test array of structs",
			input: reflect.ValueOf([2]struct {
				Name string
				Age  int
			}{
				{"Alice", 30},
				{"Bob", 5},
			}),
			expected: sprints.ClosedArray(zero, sprints.Anonymous(zero, "struct",
				sprints.KeyValue(zero, "Name", "string"),
				sprints.KeyValue(zero, "Age", "int")), 2,
				sprints.BClosedobj(zero.Inc(), sprints.Anonymous(zero, "struct",
					sprints.KeyValue(zero, "Name", "string"),
					sprints.KeyValue(zero, "Age", "int")),
					sprints.Field(zero, "Name", sprints.TypedString("Alice")),
					sprints.Field(zero, "Age", sprints.Typed("int", 30)),
				),
				sprints.BClosedobj(zero.Inc(), sprints.Anonymous(zero, "struct",
					sprints.KeyValue(zero, "Name", "string"),
					sprints.KeyValue(zero, "Age", "int")),
					sprints.Field(zero, "Name", sprints.TypedString("Bob")),
					sprints.Field(zero, "Age", sprints.Typed("int", 5)),
				),
			),
		},

		// PointerValues to StructsPtr
		{
			name: "Test pointer to struct",
			input: reflect.ValueOf(func() *struct {
				Name string
				Age  int
			} {
				x := struct {
					Name string
					Age  int
				}{
					Name: "Alice",
					Age:  30,
				}
				return &x
			}()),
			expected: sprints.Ptr(zero, sprints.BClosedobj(zero, sprints.Anonymous(zero, "struct",
				sprints.KeyValue(zero, "Name", "string"),
				sprints.KeyValue(zero, "Age", "int")),
				sprints.Field(zero, "Name", sprints.TypedString("Alice")),
				sprints.Field(zero, "Age", sprints.Typed("int", 30)),
			)),
		},

		// Nil maps defaultInterface a ReadStruct
		{
			name: "Test struct with nil map",
			input: reflect.ValueOf(struct {
				Name string
				Data map[string]int
			}{
				Name: "Alice",
				Data: nil,
			}),
			expected: sprints.BClosedobj(zero, sprints.Anonymous(zero, "struct",
				sprints.KeyValue(zero, "Name", "string"),
				sprints.KeyValue(zero, "Data", "map[string]int")),
				sprints.Field(zero, "Name", sprints.TypedString("Alice")),
				sprints.Field(zero, "Data", sprints.NilType(zero, "map[string]int")),
			),
		},

		// Empty maps defaultInterface a Slicef
		{
			name: "Test slice with empty map",
			input: reflect.ValueOf([]map[string]int{
				{},
			}),
			expected: sprints.ClosedSlice(zero, "map[string]int",
				sprints.ClosedMap(zero, "string", "int"),
			),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
			tt.StackTitle(tc.name, tc.expected)
			result := Of[internal.Option](zero, tc.input)
			compareMessages(tt, result, tc.expected)
			tt.StackLn()
		})
	}
}

func TestValue_EdgeTypes(t *testing.T) {
	zero := indent.Zero()
	tests := []struct {
		name     string
		input    reflect.Value
		expected string
	}{
		{
			name:  "Test AnonymousInterface",
			input: reflect.ValueOf(_testdata.AnonymousInterface),
			expected: sprints.BClosedobj(zero, sprints.Anonymous(zero, "struct",
				sprints.KeyValue(zero, "Name", "string"),
				sprints.KeyValue(zero, "Age", "int")),
				sprints.Field(zero, "Name", sprints.TypedString("Alice")),
				sprints.Field(zero, "Age", sprints.Typed("int", 30)),
			),
		},
		{
			name:  "Test Self Referenced Next",
			input: reflect.ValueOf(_testdata.SelfReferencedNode),
			expected: "*" + sprints.BClosedobj(zero, "github.com/andrerrcosta2/gtools/reflect4/internal/_testdata.Node",
				sprints.Field(zero, "Value", sprints.Typed("int", 1)),
				sprints.Field(zero, "Next", sprints.Ptr(zero, sprints.CyclicRef(zero, "github.com/andrerrcosta2/gtools/reflect4/internal/_testdata.Node",
					sprints.Addr(_testdata.SelfReferencedNode)))),
			),
		},
		{
			name:     "Test named int",
			input:    reflect.ValueOf(_testdata.NamedInt),
			expected: sprints.Typed("github.com/andrerrcosta2/gtools/reflect4/internal/_testdata.MyInt", 42),
		},
		{
			name:  "Test unexported field",
			input: reflect.ValueOf(_testdata.UnexportedField),
			expected: sprints.BClosedobj(zero, "github.com/andrerrcosta2/gtools/reflect4/internal/_testdata.Person",
				sprints.Field(zero, "Name", sprints.TypedString("Alice")),
				sprints.Field(zero, "age", sprints.Typed("int", 30)),
			),
		},
		{
			name:     "Test Cyclic Tree",
			input:    reflect.ValueOf(_testdata.CyclicTree),
			expected: sprints.Ptr(indent.Zero(), buildExpectedTree(indent.Zero(), 4)),
		},
		{
			name:     "Test Function Add",
			input:    reflect.ValueOf(_testdata.FunctionAdd),
			expected: zero.Sprint("func(int, int) int"),
		},
		{
			name:     "Test Simple Channel",
			input:    reflect.ValueOf(_testdata.SimpleChannel),
			expected: zero.Sprint("chan int"),
		},
		{
			name:     "Test unsafe pointer",
			input:    reflect.ValueOf(unsafe.Pointer(models.SimpleZeroInst)),
			expected: sprints.UnsafePointer(zero, unsafe.Pointer(models.SimpleZeroInst)),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tt := testingtools.LoggersLite(t, testlogs.OnFailure, themes.Color)
			tt.StackTitle(tc.name, tc.expected)
			result := Of[internal.Option](zero, tc.input)
			compareMessages(tt, result, tc.expected)
			tt.StackLn()
		})
	}
}

func buildExpectedTree(base indent.Indentor, depth int) string {
	if depth == 0 {
		return sprints.NestedBClosedobj(base, "github.com/andrerrcosta2/gtools/reflect4/internal/_testdata.Tree",
			sprints.Field(base.Inc(), "Value", sprints.Typed("int", 0)),
			// reflect can't find the type package of nil elements
			// this is golang
			sprints.Field(base.Inc(), "Left", sprints.Ptr(base, sprints.NilType(base,
				"_testdata.Tree"))),
			sprints.Field(base.Inc(), "Right", sprints.Ptr(base, sprints.NilType(base,
				"_testdata.Tree"))),
		)
	}

	return sprints.NestedBClosedobj(base, "github.com/andrerrcosta2/gtools/reflect4/internal/_testdata.Tree",
		sprints.Field(base.Inc(), "Value", sprints.Typed("int", depth)),
		sprints.Field(base.Inc(), "Left", sprints.Ptr(base, buildExpectedTree(base.Inc(), depth-1))),
		sprints.Field(base.Inc(), "Right", sprints.Ptr(base,
			sprints.CyclicRef(indent.Zero(), "github.com/andrerrcosta2/gtools/reflect4/internal/_testdata.Tree",
				"0x00000000"))),
	)
}

// Custom comparison function to skip address comparison
func compareMessages(tt gtests.Loggable, result, expected string) bool {
	tt.Helper()
	result = cyclicRefRE.ReplaceAllString(result, "<cyclic-ref|"+fmx.SGreen("0xt3st4ddr")+">")
	expected = cyclicRefRE.ReplaceAllString(expected, "<cyclic-ref|"+fmx.SGreen("0xt3st4ddr")+">")
	result = addrRE.ReplaceAllString(result, "<"+fmx.SGreen("0xt3st4ddr")+">")
	expected = addrRE.ReplaceAllString(expected, "<"+fmx.SGreen("0xt3st4ddr")+">")

	// Compare the modified result and expected strings
	if result != expected {
		tt.Error(gtests.ErrorDiff("Different messages", result, expected))
		tt.StackError(gtests.ErrorDiff("Different messages", result, expected))
		tt.StackLog(differs.Quick(result, expected))
		return false
	} else {
		tt.StackSuccessf(gtests.ErrorDiff("Equal messages", result, expected))
	}
	return true
}
