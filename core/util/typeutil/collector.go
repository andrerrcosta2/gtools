// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package typeutil

func Collector() *TypeCollector {
	return &TypeCollector{}
}

type TypeCollector struct {
	Int8       []int8
	Int16      []int16
	Int32      []int32
	Int64      []int64
	Int        []int
	Uint8      []uint8
	Uint16     []uint16
	Uint32     []uint32
	Uint64     []uint64
	Uintptr    []uintptr
	Float64    []float64
	Float32    []float32
	Complex128 []complex128
	Complex64  []complex64
	String     []string
	Boolean    []bool
	Others     []any
}

func (tc *TypeCollector) Collect(values ...any) {
	for _, v := range values {
		switch v := v.(type) {
		case int8:
			tc.Int8 = append(tc.Int8, v)
		case int16:
			tc.Int16 = append(tc.Int16, v)
		case int32:
			tc.Int32 = append(tc.Int32, v)
		case int64:
			tc.Int64 = append(tc.Int64, v)
		case int:
			tc.Int = append(tc.Int, v)
		case uint8:
			tc.Uint8 = append(tc.Uint8, v)
		case uint16:
			tc.Uint16 = append(tc.Uint16, v)
		case uint32:
			tc.Uint32 = append(tc.Uint32, v)
		case uint64:
			tc.Uint64 = append(tc.Uint64, v)
		case uintptr:
			tc.Uintptr = append(tc.Uintptr, v)
		case float64:
			tc.Float64 = append(tc.Float64, v)
		case float32:
			tc.Float32 = append(tc.Float32, v)
		case complex128:
			tc.Complex128 = append(tc.Complex128, v)
		case complex64:
			tc.Complex64 = append(tc.Complex64, v)
		case string:
			tc.String = append(tc.String, v)
		case bool:
			tc.Boolean = append(tc.Boolean, v)
		default:
			tc.Others = append(tc.Others, v)
		}
	}
}
