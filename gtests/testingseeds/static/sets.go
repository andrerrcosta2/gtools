// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package static

type SeedSets[V any, R any] interface {
	Values() V
	Refs() R
}
