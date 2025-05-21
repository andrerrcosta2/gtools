// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package typeutil

func Or[G any, K any](a any) any {
	if val, ok := a.(G); ok {
		return val
	}
	if val, ok := a.(K); ok {
		return val
	}
	return nil
}

func Ors[G any, K any](a ...any) ([]G, []K) {
	gs := make([]G, 0, len(a))
	ks := make([]K, 0, len(a))

	for _, value := range a {
		if val, ok := value.(G); ok {
			gs = append(gs, val)
		}
		if val, ok := value.(K); ok {
			ks = append(ks, val)
		}
	}
	return gs, ks
}
