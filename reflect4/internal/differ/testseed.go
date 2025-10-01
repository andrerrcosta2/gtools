// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package differ

import "math"

type MyInt int
type MyInt64 int64

type primTableTest[T any] struct {
	name     string
	a, b     T
	aType    any
	bType    any
	equal    bool
	typeName string
}

var intSeed = []primTableTest[int64]{
	{
		name:     "int: equals values",
		a:        10,
		b:        10,
		aType:    int(10),
		bType:    int(10),
		equal:    true,
		typeName: "int",
	},
	{
		name:     "int: different values",
		a:        20,
		b:        21,
		aType:    int(20),
		bType:    int(21),
		equal:    false,
		typeName: "int",
	},
	{
		name:     "int64: equals values",
		a:        42,
		b:        42,
		aType:    int64(42),
		bType:    int64(42),
		equal:    true,
		typeName: "int64",
	},
	{
		name:     "int64: different values",
		a:        -1,
		b:        1,
		aType:    int64(-1),
		bType:    int64(1),
		equal:    false,
		typeName: "int64",
	},
	{
		name:     "MyInt: different values",
		a:        100,
		b:        101,
		aType:    MyInt(100),
		bType:    MyInt(101),
		equal:    false,
		typeName: "github.com/andrerrcosta2/gtools/reflect4/internal/differ.MyInt",
	},
	{
		name:     "MyInt64: equals values",
		a:        -999,
		b:        -999,
		aType:    MyInt64(-999),
		bType:    MyInt64(-999),
		equal:    true,
		typeName: "github.com/andrerrcosta2/gtools/reflect4/internal/differ.MyInt64",
	},
}

type MyFloat32 float32
type MyFloat64 float64

var floatSeed = []primTableTest[float64]{
	{
		name:     "float32: equals values",
		a:        1.5,
		b:        1.5,
		aType:    float32(1.5),
		bType:    float32(1.5),
		equal:    true,
		typeName: "float32",
	},
	{
		name:     "float32: different values",
		a:        2.5,
		b:        3.5,
		aType:    float32(2.5),
		bType:    float32(3.5),
		equal:    false,
		typeName: "float32",
	},
	{
		name:     "float64: equals values",
		a:        1.2345,
		b:        1.2345,
		aType:    float64(1.2345),
		bType:    float64(1.2345),
		equal:    true,
		typeName: "float64",
	},
	{
		name:     "float64: different values",
		a:        5.6789,
		b:        0.1234,
		aType:    float64(5.6789),
		bType:    float64(0.1234),
		equal:    false,
		typeName: "float64",
	},
	{
		name:     "MyFloat32: equals values",
		a:        7.0,
		b:        7.0,
		aType:    MyFloat32(7.0),
		bType:    MyFloat32(7.0),
		equal:    true,
		typeName: "github.com/andrerrcosta2/gtools/reflect4/internal/differ.MyFloat32",
	},
	{
		name:     "MyFloat64: different values",
		a:        9.0,
		b:        10.0,
		aType:    MyFloat64(9.0),
		bType:    MyFloat64(10.0),
		equal:    false,
		typeName: "github.com/andrerrcosta2/gtools/reflect4/internal/differ.MyFloat64",
	},
	{
		name:     "Inf vs Inf: equals",
		a:        math.Inf(1),
		b:        math.Inf(1),
		aType:    math.Inf(1),
		bType:    math.Inf(1),
		equal:    true,
		typeName: "float64",
	},
	{
		name:     "NaN vs NaN: unequal",
		a:        math.NaN(),
		b:        math.NaN(),
		aType:    math.NaN(),
		bType:    math.NaN(),
		equal:    false, // NaN != NaN
		typeName: "float64",
	},
}

type MyComplex64 complex64
type MyComplex128 complex128

var complexSeed = []primTableTest[complex128]{
	{
		name:     "primitive complex128: equals values",
		a:        complex(1, 2),
		b:        complex(1, 2),
		aType:    complex(1, 2),
		bType:    complex(1, 2),
		equal:    true,
		typeName: "complex128",
	},
	{
		name:     "primitive complex128: different values",
		a:        complex(1, 2),
		b:        complex(2, 1),
		aType:    complex(1, 2),
		bType:    complex(2, 1),
		equal:    false,
		typeName: "complex128",
	},
	{
		name:     "named complex64: equals values",
		a:        complex(3, 4),
		b:        complex(3, 4),
		aType:    MyComplex64(complex(3, 4)),
		bType:    MyComplex64(complex(3, 4)),
		equal:    true,
		typeName: "github.com/andrerrcosta2/gtools/reflect4/internal/differ.MyComplex64",
	},
	{
		name:     "named complex64: different values",
		a:        complex(3, 4),
		b:        complex(4, 3),
		aType:    MyComplex64(complex(3, 4)),
		bType:    MyComplex64(complex(4, 3)),
		equal:    false,
		typeName: "github.com/andrerrcosta2/gtools/reflect4/internal/differ.MyComplex64",
	},
	{
		name:     "named complex128: real differs",
		a:        complex(7, 1),
		b:        complex(8, 1),
		aType:    MyComplex128(complex(7, 1)),
		bType:    MyComplex128(complex(8, 1)),
		equal:    false,
		typeName: "github.com/andrerrcosta2/gtools/reflect4/internal/differ.MyComplex128",
	},
	{
		name:     "named complex128: imaginary differs",
		a:        complex(5, 6),
		b:        complex(5, 7),
		aType:    MyComplex128(complex(5, 6)),
		bType:    MyComplex128(complex(5, 7)),
		equal:    false,
		typeName: "github.com/andrerrcosta2/gtools/reflect4/internal/differ.MyComplex128",
	},
}

type MyBool bool

var boolSeed = []primTableTest[bool]{
	{
		name:     "primitive bool: true vs true",
		a:        true,
		b:        true,
		aType:    true,
		bType:    true,
		equal:    true,
		typeName: "bool",
	},
	{
		name:     "primitive bool: true vs false",
		a:        true,
		b:        false,
		aType:    true,
		bType:    false,
		equal:    false,
		typeName: "bool",
	},
	{
		name:     "named bool: MyBool true vs false",
		a:        true,
		b:        false,
		aType:    MyBool(true),
		bType:    MyBool(false),
		equal:    false,
		typeName: "github.com/andrerrcosta2/gtools/reflect4/internal/differ.MyBool",
	},
	{
		name:     "named bool: equals values",
		a:        false,
		b:        false,
		aType:    MyBool(false),
		bType:    MyBool(false),
		equal:    true,
		typeName: "github.com/andrerrcosta2/gtools/reflect4/internal/differ.MyBool",
	},
}

type MyUint uint
type MyUint64 uint64

var uintSeed = []primTableTest[uint64]{
	{
		name:     "uint: equals values",
		a:        10,
		b:        10,
		aType:    uint(10),
		bType:    uint(10),
		equal:    true,
		typeName: "uint",
	},
	{
		name:     "uint: different values",
		a:        20,
		b:        21,
		aType:    uint(20),
		bType:    uint(21),
		equal:    false,
		typeName: "uint",
	},
	{
		name:     "uint64: equals values",
		a:        42,
		b:        42,
		aType:    uint64(42),
		bType:    uint64(42),
		equal:    true,
		typeName: "uint64",
	},
	{
		name:     "uint64: different values",
		a:        0,
		b:        1,
		aType:    uint64(0),
		bType:    uint64(1),
		equal:    false,
		typeName: "uint64",
	},
	{
		name:     "MyUint: equals values",
		a:        100,
		b:        100,
		aType:    MyUint(100),
		bType:    MyUint(100),
		equal:    true,
		typeName: "github.com/andrerrcosta2/gtools/reflect4/internal/differ.MyUint",
	},
	{
		name:     "MyUint: different values",
		a:        100,
		b:        101,
		aType:    MyUint(100),
		bType:    MyUint(101),
		equal:    false,
		typeName: "github.com/andrerrcosta2/gtools/reflect4/internal/differ.MyUint",
	},
	{
		name:     "MyUint64: equals values",
		a:        999,
		b:        999,
		aType:    MyUint64(999),
		bType:    MyUint64(999),
		equal:    true,
		typeName: "github.com/andrerrcosta2/gtools/reflect4/internal/differ.MyUint64",
	},
	{
		name:     "MyUint64: different values",
		a:        1234,
		b:        5678,
		aType:    MyUint64(1234),
		bType:    MyUint64(5678),
		equal:    false,
		typeName: "github.com/andrerrcosta2/gtools/reflect4/internal/differ.MyUint64",
	},
	// Edge case: max uint64
	{
		name:     "uint64: max value equals",
		a:        math.MaxUint64,
		b:        math.MaxUint64,
		aType:    uint64(math.MaxUint64),
		bType:    uint64(math.MaxUint64),
		equal:    true,
		typeName: "uint64",
	},
	{
		name:     "uint64: max vs max-1",
		a:        math.MaxUint64,
		b:        math.MaxUint64 - 1,
		aType:    uint64(math.MaxUint64),
		bType:    uint64(math.MaxUint64 - 1),
		equal:    false,
		typeName: "uint64",
	},
}
