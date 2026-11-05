// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package obs

type TransformPipe[I, O any] func(Observable[I]) Observable[O]

func Transform[I, O any](o Observable[I], pipe TransformPipe[I, O]) Observable[O] {
	return pipe(o)
}
