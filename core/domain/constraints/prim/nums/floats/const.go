// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package floats

const (
	// Exact integer range: [-MaxPrecision32, MaxPrecision32]
	MaxPrecision32 = 1 << 24 // 16_777_216
	MinPrecision32 = -MaxPrecision32

	// First integers beyond exact representation (rounding begins)
	FirstRounded32 = MaxPrecision32 + 1 // 16_777_217

	// Spacing changes at powers of two
	Spacing2Start32  = 1 << 24 // 16_777_216
	Spacing4Start32  = 1 << 25 // 33_554_432
	Spacing8Start32  = 1 << 26 // 67_108_864
	Spacing16Start32 = 1 << 27 // 134_217_728

	// Max32 highest finite value representable by float32
	Max32 = 3.40282346638528859811704183484516925440e+38
	// Min32 smallest finite value representable by float32
	Min32 = -Max32
	// MinPositiveNormal32 smallest possible positive number that can be represented by the float64
	// Normal numbers have:
	//	- full precision - full exponent + mantissa
	//  - normalized representation - leading 1 in the mantissa is implicit
	//  - anything smaller than this becomes subnormal and loses precision.
	MinPositiveNormal32 = 1.17549435082228750796873653722224567781e-38
	// MinPositiveSubnormal32  smallest positive number of any kind that float64 can represent
	// Subnormals:
	//	- have no leading 1 bitwise in the mantissa.
	//  - allow floating-point numbers to gradually approach zero, avoiding a sudden underflow to 0
	//  - have less precision, especially in arithmetic.
	MinPositiveSubnormal32 = 1.40129846432481707092372958328991613128e-45

	// Epsilon32 is the difference between 1 and the next larger representable float
	Epsilon32 = 1.1920928955078125e-7
)

const (
	// 	MaxPrecision64 largest integer of exact representation
	MaxPrecision64 = 1 << 53 // 9_007_199_254_740_992
	MinPrecision64 = -MaxPrecision64

	// FirstRounded64 first integer beyond exact representation
	FirstRounded64 = MaxPrecision64 + 1 // 9_007_199_254_740_993

	// Spacing2Start64 smallest float64 round by groups of 2
	Spacing2Start64 = 1 << 53 // 9_007_199_254_740_992
	// 	Spacing4Start64 smallest float64 round by groups of 4
	Spacing4Start64 = 1 << 54 // 18_014_398_509_481_984
	// Spacing8Start64 smallest float64 round by groups of 8
	Spacing8Start64 = 1 << 55 // 36_028_797_018_963_968
	// Spacing16Start64 smallest float64 round by groups of 16
	Spacing16Start64 = 1 << 56 // 72_057_594_037_927_936

	// Max64 largest finite value representable by float64
	Max64 = 1.7976931348623157081452742373170435679807e+308
	// Min64 smallest finite value representable by float64
	Min64 = -Max64
	// MinPositiveNormal64 smallest possible positive number that can be represented by the float64
	// Normal numbers have:
	//	- full precision - full exponent + mantissa
	//  - normalized representation - leading 1 in the mantissa is implicit
	//  - anything smaller than this becomes subnormal and loses precision.
	MinPositiveNormal64 = 2.225073858507201383090232717332404064219e-308
	// MinPositiveSubnormal64  smallest positive number of any kind that float64 can represent
	// Subnormals:
	//	- have no leading 1 bitwise in the mantissa.
	//  - allow floating-point numbers to gradually approach zero, avoiding a sudden underflow to 0
	//  - have less precision, especially in arithmetic.
	MinPositiveSubnormal64 = 4.940656458412465441765687928682213723651e-324

	// Epsilon64 is the difference between 1 and the next larger representable float
	Epsilon64 = 2.220446049250313080847263336181640625e-16
)

// Relationships between different number type4
const (
	// Div32Uint64 division of floats.Max32 by uints.Max64
	Div32Uint64 = 18446742974197923840
	// Div32Uint32 division of floats.Max32 by uints.Max32
	Div32Uint32 = 79228157810344598797608288256
	// Div32Int64 division of floats.Max32 by ints.Max64
	Div32Int64 = 36893485948395847680
	// Div32Int32 division of floats.Max32 by ints.Max32
	Div32Int32 = 158456315657582685742635679744
	// Div64Uint32 division of floats.Max64 by uints.Max32
	Div64Uint32 = 4.185580497795888e+298
	// Div64Uint64 division of floats.Max64 by uints.Max64
	Div64Uint64 = 9.745314011399998e+288
	// Div64Int32 division of floats.Max64 by ints.Max32
	Div64Int32 = 8.371160997540839e+298
	// Div64Int64 division of floats.Max64 by ints.Max64
	Div64Int64 = 1.9490628022799996e+289
)
