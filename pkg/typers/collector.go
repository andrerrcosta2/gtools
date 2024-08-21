// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package typers

func Collector() *TypeCollector {
	return &TypeCollector{}
}

type TypeCollector struct {
	Integer      []int
	Float64      []float64
	Float32      []float32
	Byte         []byte
	Complex128   []complex128
	Complex64    []complex64
	String       []string
	Boolean      []bool
	Unrecognized []any
}

func (tc *TypeCollector) Collect(values ...any) {
	for _, v := range values {
		switch v := v.(type) {
		case int:
			tc.Integer = append(tc.Integer, v)
		case float64:
			tc.Float64 = append(tc.Float64, v)
		case float32:
			tc.Float32 = append(tc.Float32, v)
		case byte:
			tc.Byte = append(tc.Byte, v)

		case complex128:
			tc.Complex128 = append(tc.Complex128, v)
		case complex64:
			tc.Complex64 = append(tc.Complex64, v)
		case string:
			tc.String = append(tc.String, v)
		case bool:
			tc.Boolean = append(tc.Boolean, v)
		default:
			tc.Unrecognized = append(tc.Unrecognized, v)
		}
	}
}
