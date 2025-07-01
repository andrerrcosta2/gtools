// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package zero

func referencesSet() []interface{} {
	var p1 bool
	var p2 int
	var p3 int8
	var p4 int16
	var p5 int32
	var p6 int64
	var p7 uint
	var p8 uint8
	var p9 uint16
	var p10 uint32
	var p11 uint64
	var p12 uintptr
	var p13 float32
	var p14 float64
	var p15 complex64
	var p16 complex128
	var p17 string
	return []interface{}{&p1, &p2, &p3, &p4, &p5, &p6, &p7, &p8, &p9, &p10, &p11, &p12, &p13, &p14, &p15, &p16, &p17}
}
