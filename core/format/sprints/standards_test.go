// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sprints

import (
	"github.com/andrerrcosta2/gtools/core/format/code/indent"
	"github.com/andrerrcosta2/gtools/core/format/fmx"
	"strconv"
	"testing"
	"unsafe"
)

// TestAddr tests the Addr function
//
// Just a standard 0xaddress sprint
//   - should print the address in hex
//   - should print 0x0 if the address is nil
func TestAddr(t *testing.T) {
	// Valid address
	i := 1
	p := uintptr(unsafe.Pointer(&i))
	s := Addr(&i)

	if s != "0x"+strconv.FormatInt(int64(p), 16) {
		t.Errorf("Expected %s, got %s", "0x"+strconv.FormatInt(int64(p), 16), s)
	}

	// Invalid address
	var intPtr *int
	s = Addr(intPtr)
	if s != "0x0" {
		t.Errorf("Expected %s, got %s", "0x0", s)
	}
}

func TestObjects(t *testing.T) {
	zero := indent.Zero()
	tc := []struct {
		name     string
		object   string
		expected string
	}{
		{
			name: "Test Simple Struct",
			object: BClosedobj(indent.Zero(), "SimpleObject",
				Field(indent.Zero(), "Name", "John"),
				Field(indent.Zero(), "Age", Digit(30)),
			),
			expected: fmx.SBold("SimpleObject") + "{" +
				"\n\tName: John," + "" +
				"\n\tAge: 30," +
				"\n}",
		},
		{
			name: "Test Nested Struct",
			object: BClosedobj(indent.Zero(), "NestedObject",
				Field(zero, "Name", "Alice"),
				Field(zero, "Age", Digit(30)),
				Field(zero, "Child", BClosedobj(zero.Inc(), "Child",
					Field(zero, "Name", "Bob"),
					Field(zero, "Age", Digit(5)),
				)),
			),
			expected: fmx.SBold("NestedObject") + "{" +
				"\n\tName: Alice," +
				"\n\tAge: 30," +
				"\n\tChild: " + fmx.SBold("Child") + "{" +
				"\n\t\tName: Bob," +
				"\n\t\tAge: 5," +
				"\n\t}," +
				"\n}",
		},
		{
			name: "Test Struct With Maps And Slices",
			object: BClosedobj(indent.Zero(), "StructWithMapsAndSlices",
				Field(zero, "User", "Alice"),
				Field(zero, "Email", "alice@me.com"),
				Field(zero, "Roles", ClosedSlice(zero.Inc(), "string", "reader", "writer")),
				Field(zero, "Permissions", ClosedMap(zero.Inc(), "string", "string",
					Field(zero, "read", "true"),
					Field(zero, "write", "false"),
				)),
			),
			expected: fmx.SBold("StructWithMapsAndSlices") + "{" +
				"\n\tUser: Alice," +
				"\n\tEmail: alice@me.com," +
				"\n\tRoles: []string[" +
				"\n\t\treader," +
				"\n\t\twriter," +
				"\n\t]," +
				"\n\tPermissions: map[string]string{" +
				"\n\t\tread: true," +
				"\n\t\twrite: false," +
				"\n\t}," +
				"\n}",
		},
		{
			name: "Test Slicef of Maps",
			object: ClosedSlice(zero, "map[string]int",
				ClosedMap(zero.Inc(), "string", "int",
					Field(zero, "one", "1"),
					Field(zero, "two", "2"),
				),
				ClosedMap(zero.Inc(), "string", "int",
					Field(zero, "three", "3"),
					Field(zero, "four", "4"),
				),
			),
			expected: "[]map[string]int[" +
				"\n\tmap[string]int{" +
				"\n\t\tone: 1," +
				"\n\t\ttwo: 2," +
				"\n\t}," +
				"\n\tmap[string]int{" +
				"\n\t\tthree: 3," +
				"\n\t\tfour: 4," +
				"\n\t}," +
				"\n]",
		},
		{
			name: "Test Slicef of Anonymous Structs",
			object: ClosedArray(zero, Anonymous(zero, "struct",
				KeyValue(zero, "Name", "string"),
				KeyValue(zero, "Age", "int")), 2,
				BClosedobj(zero.Inc(), "",
					Field(zero, "Name", TypedString("Alice")),
					Field(zero, "Age", Typed("int", 30)),
				),
				BClosedobj(zero.Inc(), "",
					Field(zero, "Name", TypedString("Bob")),
					Field(zero, "Age", Typed("int", 5)),
				),
			),
			expected: "[2]struct{ Name string; Age int }[" +
				"\n\t{" +
				"\n\t\tName: <string>Alice," +
				"\n\t\tAge: <int>30," +
				"\n\t}," +
				"\n\t{" +
				"\n\t\tName: <string>Bob," +
				"\n\t\tAge: <int>5," +
				"\n\t}," +
				"\n]",
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			result := test.object
			if result != test.expected {
				t.Errorf("\nReceived: \n%s, \nExpected \n%s", result, test.expected)
				t.Log(fmx.SRedf("result length %d", len(result)))
				t.Log(fmx.SRedf("expected length %d", len(test.expected)))
			}
		})
	}
}

// TestNilType tests the TypedString function
//
// just should print the standard format of a 'type<nil>'
//   - assert correct indentation
func TestNilType(t *testing.T) {
	s := NilType(indent.Zero(), "string")
	if s != "string<nil>" {
		t.Errorf("expected: '%q', got: '%q'", "string<nil>", s)
	}

	s = NilType(indent.Tab(1), "string")
	if s != "\tstring<nil>" {
		t.Errorf("expected: '%q', got: '%q'", "string<nil>", s)
	}
}

// TestType tests the TypedString function
//
// just should print the standard format of a '<string>quoted-value':
//   - should print the quoted value if it is not empty
//   - should print "(empty)" if the value is empty
//   - assert correct indentation
func TestType(t *testing.T) {
	// Empty string
	s := TypedString("")
	if s != "<string>(empty)" {
		t.Errorf("expected: '%q', got: '%q'", "<string>(empty)", s)
	}

	// Non-empty string
	s = TypedString("hello")
	if s != "<string>\"hello\"" {
		t.Errorf("expected: '%q', got: '%q'", "<string>hello", s)
	}
}
