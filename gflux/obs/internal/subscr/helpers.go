// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package subscr

func CompactAliveSlots[T any](slots []*Slot[T]) []*Slot[T] {
	dst := slots[:0]
	for _, s := range slots {
		if !s.Closed() {
			dst = append(dst, s)
		}
	}
	return dst
}
