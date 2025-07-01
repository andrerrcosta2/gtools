// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package meml

import "unsafe"

type emptyInterface struct {
	typx unsafe.Pointer
	word unsafe.Pointer
}

func UnsafePointerOf(v any) unsafe.Pointer {
	return (*emptyInterface)(unsafe.Pointer(&v)).word
}
