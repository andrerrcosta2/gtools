// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package slices

func Of[T ~[]E, E any](cap int, values ...E) T {
	if cap < len(values) {
		panic("slice capacity cannot be less than slice length")
	}
	s := make(T, len(values), cap)
	for v := range values {
		s[v] = values[v]
	}
	return s
}
