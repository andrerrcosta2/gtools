// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package sorters

type Sorter[S ~[]E, E any] interface {
	Sort(arr *S)
}
